package main

import (
	"time"
	"os"
	"database/sql"
	"log"
	"fmt"
)

type Skill int

type Customer struct{
	id int
	name  string
	surname string
	address Address
	skill Skill
	email string
	telephone string
	birthday time.Time

}

type Address struct{
	street string
	number int
	zipcode int
	city string
}

const (
	BEGINNER Skill = 0
	ADVANCED Skill = 1
	PRO Skill = 2
)

func main(){
	chingling := Customer{5,"ching","ling",Address{"xia lu",94134,1345,"peking"},Skill(0),"Cingling@chingchongchang.co.cn","+12349153",time.Date(1990,01,15,00,00,00,00,nil)}
	Insert(&chingling)
}

func Insert(customer *Customer) {

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection
	defer db.Close()

	//create table if it doesn't exist
	_,err = db.Query("CREATE TABLE IF NOT EXISTS customers (id PRIMARY  KEY, name text NOT NULL, surname text NOT NULL,street text NOT NULL, number integer NOT NULL,zipcode integer NOT NULL,city text NOT NULL,skill integer NOT NULL,email text NOT NULL,telephone text NOT NULL,birthday date NOT NULL)")
	panicErr(err)

	// Add a user to it
	_,err = db.Query("INSERT INTO customers(id,name,surname,street, number,zipcode,city,skill,email,telephone,birthday) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
		customer.id, customer.name, customer.surname, customer.address.street, customer.address.number, customer.address.zipcode,
		customer.address.city, customer.skill, customer.email, customer.telephone, customer.birthday);
	panicErr(err)

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	printDatabaseContent(rows);

}

func Remove(customer *Customer) {

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection
	defer db.Close()

	//remove users
	_,err = db.Query("DELETE FROM customers WHERE id=customer.id");
	panicErr(err)

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	printDatabaseContent(rows);
}

func Update(customer *Customer) {

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection
	defer db.Close()

	// update user in database
	_,err = db.Query("UPDATE customers SET name=customer.name, surname=customer.surname,street=customer.address.street,number=customer.address.number,zipcode=customer.address.zipcode,city=customer.address.city,skill=customer.skill,email=customer.email,telephone=customer.telephone,birthday=customer.birthday WHERE id = customer.id;");
	panicErr(err)

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	printDatabaseContent(rows);
}

func printDatabaseContent(rows *sql.Rows){

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