package main

import (
	"net/http"
	"github.com/graphql-go/handler"
	"fmt"
	"github.com/ob-vss-ss18/ppl-customer/Database"
)

func main() {
	// Initialize Data
	Database.InitGraphQL()
	Database.InitializeCustomerDB()

	// Start HTTP Server
	handle := handler.New(&handler.Config{
		Schema: &Database.CustomerSchema,
		Pretty: true,
		GraphiQL: true,
	})
	http.Handle("/query", handle)

	http.HandleFunc("/", hello)

	//for local debugging
	err := http.ListenAndServe(":5000", nil)


	//for heroku  usage
	//err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)


	if err != nil {
		panic(err)
	}
}


func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}