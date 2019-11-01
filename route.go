package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// API
	router := httprouter.New()
	router.GET("/api/v1/authors", getAuthors)
	router.GET("/api/v1/authors/:author/albums", getAlbums)
	router.GET("/api/v1/authors/:author/albums/:album/songs", getSongs)

	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/api/", router)

	return mux
}

func getAuthors(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var authors = []string{
		"ABBA",
		"Moby",
		"Scorpions",
		"Группа Кино",
	}

	json.NewEncoder(w).Encode(authors)
}

func getAlbums(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var albums = []string{
		"album 1",
		"album 2",
		"album 3",
		"альбом 4",
	}

	json.NewEncoder(w).Encode(albums)
}

func getSongs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var songs = []string{
		"Song 1",
		"Song 2",
		"Song 3",
		"Песня 4",
	}

	json.NewEncoder(w).Encode(songs)
}
