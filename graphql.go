package main

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/relay/examples/starwars"
)

func init() {
	http.HandleFunc("/", configServer)

	http.Handle("/", handler.New(&handler.Config{
		Schema: &starwars.Schema,
		Pretty: true,
	}))
}

func configServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
