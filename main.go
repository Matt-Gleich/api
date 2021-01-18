package main

import (
	"net/http"

	"github.com/Matt-Gleich/api/pkg/schema"
	"github.com/Matt-Gleich/lumber"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: schema.Fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	lumber.Fatal(err, "Failed to create new schema")

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Playground: true,
	})

	http.Handle("/", h)
	err = http.ListenAndServe(":8080", nil)
	lumber.Fatal(err, "Failed to listen and serve for requests")
}
