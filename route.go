package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func newRouter(s *service) *http.ServeMux {
	mux := http.NewServeMux()

	// API
	router := httprouter.New()
	router.GET("/api/v1/authors", logMiddleware(s.getAuthors))
	router.GET("/api/v1/authors/:author/albums", logMiddleware(s.getAlbums))
	router.GET("/api/v1/authors/:author/albums/:album/songs", logMiddleware(s.getSongs))

	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/api/", router)

	return mux
}

func (s *service) getAuthors(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var authors = make([]string, 0)

	// TODO: Избавиться от повторений исполнителей,
	// если те содержаться в двух и более различных форматах.
	for _, format := range s.root.Children {
		for _, author := range format.Children {

			authors = append(authors, author.Name)

		}
	}

	json.NewEncoder(w).Encode(authors)
}

func (s *service) getAlbums(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var albums = make([]string, 0)

	// В результирующем списке альбомов исполнителя
	// находятся все, что есть во всех форматах.
	for _, format := range s.root.Children {
		for _, author := range format.Children {

			if p.ByName("author") == author.Name {
				for _, album := range author.Children {

					albums = append(albums, album.Name)

				}
			}
		}
	}

	json.NewEncoder(w).Encode(albums)
}

func (s *service) getSongs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var songs = make([]string, 0)

	for _, format := range s.root.Children {
		for _, author := range format.Children {

			if p.ByName("author") == author.Name {
				for _, album := range author.Children {

					if p.ByName("album") == album.Name {
						for _, song := range album.Children {

							songs = append(songs, song.Name)

						}
					}
				}
			}
		}
	}

	json.NewEncoder(w).Encode(songs)
}

func logMiddleware(nextHandler httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Println(r.URL.RequestURI())
		defer nextHandler(w, r, p)
	}

}
