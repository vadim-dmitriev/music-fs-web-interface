package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func newRouter() fasthttp.RequestHandler {
	router := fasthttprouter.New()

	router.ServeFiles("/*filepath", "static")

	return router.Handler
}
