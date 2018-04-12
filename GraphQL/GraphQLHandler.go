package GraphQL

import (
	"net/http"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"log"
	"encoding/json"

)

var	schema graphql.Schema



func Test(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Heroku!")

	if req.Method == "POST" && req.Header.Get("Content-Type") == "application/graphql"{
		fields := graphql.Fields{
			"Hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "WORLD", nil
				},
			},
		}
		rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
		schema.AppendType(graphql.NewObject(rootQuery))

	}

}

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

func Incoming(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Es ist nicht so wie du denkst, wenn du das denkst, was ich denke, was du denkst, denn das denken der Gedanken ist ein denkenloses Denken darum denke nicht gedacht zu haben")

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
	fmt.Fprintf(res, "%s \n", rJSON)
	//fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}