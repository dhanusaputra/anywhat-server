package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (*query) Hello() string {
	return "Hello, world!"
}

func main() {
	s := `
                type Query {
                        hello: String!
                }
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":9090", nil))
}
