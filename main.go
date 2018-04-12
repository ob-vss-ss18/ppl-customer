package main

import (
	"net/http"
	//"os"
	"github.com/ob-vss-ss18/ppl-customer/GraphQL"

)

func main(){

	http.Handle("/query", GraphQL.InitHandler())

	//for local debugging
	//err := http.ListenAndServe(":5000", GraphQL.InitHandler())

	//for heroku  usage
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}


}





