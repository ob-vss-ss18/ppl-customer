package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"os"
)




func main() {
	connStr := os.Getenv("DATABASE_URL")
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
