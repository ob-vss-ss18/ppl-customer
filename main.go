package main

import (
	"net/http"
	"os"
	"github.com/graphql-go/handler"
	"fmt"
)

func main() {
	// Initialize Data
	InitializeUserDB()

	// Start HTTP Server
	handle := handler.New(&handler.Config{
		Schema: &UserSchema,
		Pretty: true,
		GraphiQL: true,
	})
	http.Handle("/query", handle)

	http.HandleFunc("/", hello)

	//for local debugging
	//err := http.ListenAndServe(":5000", nil)

	//for heroku  usage
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)


	if err != nil {
		panic(err)
	}

}


func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}