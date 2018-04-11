package main

import ("fmt"
	"net/http"
	"os"
	"log"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func main(){
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Heroku!")

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
		fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
	}

