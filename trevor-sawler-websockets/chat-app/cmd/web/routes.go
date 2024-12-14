package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"chat-app/internal/handlers"
)

func routes() http.Handler{
	mux := pat.New()
	mux.Get("/",http.HandlerFunc(handlers.Home))

	return mux
}