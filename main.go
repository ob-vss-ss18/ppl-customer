package main

import (
	"net/http"
	"os"
	"github.com/graphql-go/handler"
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

	//for local debugging
	//err := http.ListenAndServe(":5000", nil)

	//for heroku  usage
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)


	if err != nil {
		panic(err)
	}

}
