package Database

import (
	"time"
	"log"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	_ "github.com/lib/pq"
	"errors"
)

func InitializeCustomerDB() {
	err := InitializeTables()
	panicErr(err)
}

func SelectAll() map[int]Customer {

	db, err := openDatabase()
	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	customers := make(map[int]Customer)

	for rows.Next() {
		var (
			id        int
			name      string
			surname   string
			street    string
			number    int
			zipcode   int
			city      string
			skill     int
			email     string
			telephone string
			birthday  time.Time
		)
		if err := rows.Scan(&id, &name, &surname, &street, &number, &zipcode, &city, &skill, &email, &telephone, &birthday); err != nil {
			log.Fatal(err)
		}

		customers[id] = Customer{id, name, surname, Address{street, number, zipcode, city}, Skill(skill), email, telephone, birthday}
	}
	rows.Close()
	closeDatabase(db, nil)
	return customers
}

func InitializeTables() (error) {

	db, err := openDatabase()

	//generate customer id
	row, err := db.Query("CREATE TABLE IF NOT EXISTS idNumbers (id serial primary key)")
	panicErr(err)
	row.Close()
	//create customer database if not existent
	row, err = db.Query("CREATE TABLE IF NOT EXISTS customers (id serial PRIMARY KEY, name text NOT NULL, surname text NOT NULL,street text NOT NULL, number integer NOT NULL,zipcode integer NOT NULL,city text NOT NULL,skill integer NOT NULL,email text NOT NULL,telephone text NOT NULL,birthday text NOT NULL)")
	panicErr(err)
	row.Close()
	//local entry, just something is in de database
	closeDatabase(db, nil)

	return err
}

//TODO method stub for lukas to fill up
func Select(ID int) (Customer, error) {

	db, err := openDatabase()
	defer closeDatabase(db, nil)
	row, err := db.Query("SELECT * FROM customers WHERE id = $1", ID)
	panicErr(err)
	var customer Customer

	var (
		id        int
		name      string
		surname   string
		street    string
		number    int
		zipcode   int
		city      string
		skill     int
		email     string
		telephone string
		birthday  time.Time
	)

	if row.Next() {
		row.Scan(&id, &name, &surname, &street, &number, &zipcode, &city, &skill, &email, &telephone, &birthday);
	} else {
		return customer, errors.New("Id not found!")
	}

	customer = Customer{id, name, surname, Address{street, number, zipcode, city}, Skill(skill), email, telephone, birthday}

	row.Close()
	return customer, nil

}

func InsertCustomer(customer Customer) int {

	return Insert(customer.name, customer.surname, customer.address.street, customer.address.number, customer.address.zipcode, customer.address.city, customer.skill, customer.email, customer.telephone, customer.birthday)
}

func Insert(name string, surname string, street string, number int, zipcode int, city string, skill Skill, email string, telephone string, birthday time.Time) int {
	db, err := openDatabase()

	//moved to InitializeTables
	//generate customer id
	//_, err = db.Query("CREATE TABLE IF NOT EXISTS idNumbers (id serial primary key)")
	//panicErr(err)

	rows, err := db.Query("SELECT id FROM idNumbers")
	panicErr(err)

	var iteratedIdPart int

	rows.Next()
	rows.Scan(&iteratedIdPart)

	rows.Close()
	//define seed for random numbers
	seedSource := rand.NewSource(time.Now().UnixNano())
	randSeed := rand.New(seedSource)

	id := 1000000 + (iteratedIdPart * 100) + randSeed.Intn(99)

	rows, err = db.Query("UPDATE idNumbers SET id=$1", iteratedIdPart+1)

	rows.Close()
	// Add a customer to it
	rows, err = db.Query("INSERT INTO customers(id,name,surname,street, number,zipcode,city,skill,email,telephone,birthday) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
		id, name, surname, street, number, zipcode, city, skill, email, telephone, birthday);
	panicErr(err)
	rows.Close()

	rows, err = db.Query("SELECT * FROM customers")
	panicErr(err)

	rows.Close()
	//additional
	closeDatabase(db, rows)

	return id
}

func Remove(customer *Customer) {

	db, err := openDatabase()

	//remove users
	row, err := db.Query("DELETE FROM customers WHERE id=$1;", customer.id);
	panicErr(err)
	row.Close()

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	closeDatabase(db, rows)
}

func Update(customer *Customer) {

	db, err := openDatabase()

	// update user in database
	row, err := db.Query("UPDATE customers SET name=$2, surname=$3,street=$4 ,number=$5,zipcode=$6,city=$7,skill=$8,email=$9,telephone=$10,birthday=$11 WHERE id = $1;",
		customer.id, customer.name, customer.surname, customer.address.street, customer.address.number, customer.address.zipcode,
		customer.address.city, customer.skill, customer.email, customer.telephone, customer.birthday)
	panicErr(err)
	row.Close()

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	closeDatabase(db, rows);
}

func openDatabase() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection

	return db, err
}

func closeDatabase(db *sql.DB, rows *sql.Rows) {

	if rows != nil {
		for rows.Next() {
			var (
				id        string
				name      string
				surname   string
				street    string
				number    int
				zipcode   int
				city      string
				skill     int
				email     string
				telephone string
				birthday  string
			)
			if err := rows.Scan(&id, &name, &surname, &street, &number, &zipcode, &city, &skill, &email, &telephone, &birthday); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("id: %s, name: %s, surname: %s,street: %s,number: %d,zipcode: %d,city: %s,skill: %d,email: %s,telephone: %s,birthday: %s\n", id, name, surname, street, number, zipcode, city, skill, email, telephone, birthday)
		}
	}

	err := db.Close()
	panicErr(err)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
