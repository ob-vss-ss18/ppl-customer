package GraphQL

import (
	"net/http"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"encoding/json"
)

func Test(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Heroku!")
}

func Incoming(res http.ResponseWriter, req *http.Request) {

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
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Fprintln(res, "%s \n", rJSON)
	//fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}