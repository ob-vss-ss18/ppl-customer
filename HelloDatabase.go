package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	//URI to connect to the postgress database, requires a user and a password input
	DB_URI_FORMAT string = "postgres://%s:%s@ec2-46-137-109-220.eu-west-1.compute.amazonaws.com:5432/d4ppp9g29cefd9"
)


func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	connStr := fmt.Sprintf(DB_URI_FORMAT, dbUser, dbPass)
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection
	defer db.Close()

	// Create a test table if it doesnt exist already
	_,err = db.Query("CREATE TABLE IF NOT EXISTS example (id serial PRIMARY  KEY, name text NOT NULL)")
	panicErr(err)

	// Add a user to it
	_,err = db.Query("INSERT INTO example(id,name) VALUES($1,$2)", 1,"Hans")
	panicErr(err)

	rows, err := db.Query("SELECT * FROM example")
	panicErr(err)

	defer rows.Close()
	for  rows.Next(){
		var (
			id string
			name string
		)
		if err := rows.Scan(&id,&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %s, name: %s\n", id, name)
	}
}

func panicErr(err error){
	if err != nil{
		panic(err)
	}
}
