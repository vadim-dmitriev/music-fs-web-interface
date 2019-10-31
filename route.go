package main

import (
	"encoding/json"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func newRouter() fasthttp.RequestHandler {
	router := fasthttprouter.New()

	// router.ServeFiles("/*filepath", "static")
	router.ServeFiles("/*filepath", "static")

	// API
	// router.GET("/api/v1/authors", getAuthors)
	// router.GET("/api/v1/authors/:author/albums", getAlbums)
	// router.GET("/api/v1/authors/:author/albums/:album/songs", getSongs)

	return router.Handler
}

func getAuthors(ctx *fasthttp.RequestCtx) {
	var authors = []string{
		"ABBA",
		"Moby",
		"Scorpions",
		"Группа Кино",
	}

	resp, err := json.Marshal(authors)
	if err != nil {
		panic(err)
	}

	ctx.Write(resp)
}

func getAlbums(ctx *fasthttp.RequestCtx) {
	var albums = []string{
		"album 1",
		"album 2",
		"album 3",
		"альбом 4",
	}

	resp, err := json.Marshal(albums)
	if err != nil {
		panic(err)
	}

	ctx.Write(resp)
}

func getSongs(ctx *fasthttp.RequestCtx) {
	var songs = []string{
		"Song 1",
		"Song 2",
		"Song 3",
		"Песня 4",
	}

	resp, err := json.Marshal(songs)
	if err != nil {
		panic(err)
	}

	ctx.Write(resp)
}
