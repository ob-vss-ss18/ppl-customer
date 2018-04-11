package main

import ("fmt"
	"net/http"
	"os"
	"github.com/ob-vss-ss18/ppl-customer/GraphQL"
)

func main(){
	http.HandleFunc("/", GraphQL.Incoming)
	http.HandleFunc("/test", GraphQL.Test)
	http.Handle("/query", GraphQL.InitHandler())
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}



