package main

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/relay/examples/starwars"
)

func init() {
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema: &starwars.Schema,
		Pretty: true,
	}))
}
