package GraphQL

import (
	"net/http"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"log"


)

var	schema graphql.Schema

func initializeScheme() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	var err error
	schema, err = graphql.NewSchema(schemaConfig)


	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
}

func InitHandler() http.Handler{

	initializeScheme()


	handle := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	return handle;
}

