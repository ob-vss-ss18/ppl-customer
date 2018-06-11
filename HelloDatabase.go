package main

import (
	_"github.com/lib/pq"
	"database/sql"
	"os"
	"log"
	"fmt"
)




func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection
	defer db.Close()

	// Create a test table if it doesnt exist already
	//_,err = db.Query("CREATE TABLE IF NOT EXISTS example (id serial PRIMARY  KEY, name text NOT NULL)")
	//panicErr(err)

	// Add a user to it
	//_,err = db.Query("INSERT INTO example(id,name) VALUES($1,$2)", 6,"bast");
	//panicErr(err)

	_,err = db.Query("CREATE TABLE IF NOT EXISTS customers (id serial PRIMARY KEY, name text NOT NULL, surname text NOT NULL,street text NOT NULL, number integer NOT NULL,zipcode integer NOT NULL,city text NOT NULL,skill integer NOT NULL,email text NOT NULL,telephone text NOT NULL,birthday date NOT NULL)")
	panicErr(err)

	// Add a user to it
	//_,err = db.Query("INSERT INTO customers(id,name,surname,street, number,zipcode,city,skill,email,telephone,birthday) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", 4,"seppel","mueller","schulestr",13,81325,"peking",1,"test@test.co.cn","+23413241234","2018-03-15");
	//panicErr(err)

	//remove users
	//_,err = db.Query("DELETE FROM example WHERE id='5';");
	//panicErr(err)

	panicErr(err)
	// Add a user to it
	_,err = db.Query("UPDATE customers SET name='ching', surname='ling',street='dung'WHERE id = '4';");

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()
	for  rows.Next(){
		var (
			id string
			name string
			surname string
			street string
			number int
			zipcode int
			city string
			skill int
			email string
			telephone string
			birthday string
		)
		if err := rows.Scan(&id,&name,&surname,&street,&number,&zipcode,&city,&skill,&email,&telephone,&birthday); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %s, name: %s, surname: %s,street: %s,number: %s,zipcode: %s,city: %s,skill: %s,email: %s,telephone: %s,birthday: %s\n", id, name,surname,street,number,zipcode,city,skill,email,telephone,birthday)
	}
}

func panicErr(err error){
	if err != nil{
		panic(err)
	}
}
