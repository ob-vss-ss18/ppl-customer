package Database

import (
	"time"
	_"github.com/lib/pq"
	"database/sql"
	"os"
	"log"
	"fmt"
	//_"math/rand"
	"math/rand"
	"github.com/graphql-go/graphql"
)

type Skill int

var(
	customerSchema graphql.Schema
	customerType *graphql.Object
)

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

func SelectAll() map[int]Customer{

	db, err := openDatabase()
	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	customers := make(map[int]Customer)

	for  rows.Next(){
		var (
			id int
			name string
			surname string
			street string
			number int
			zipcode int
			city string
			skill int
			email string
			telephone string
			birthday time.Time
		)
		if err := rows.Scan(&id,&name,&surname,&street,&number,&zipcode,&city,&skill,&email,&telephone,&birthday); err != nil {
			log.Fatal(err)
		}

		customers[id] = Customer{id,name,surname,Address{street,number,zipcode,city},Skill(skill),email,telephone,birthday}
	}

	return customers
}

func InitializeCustomerDB() {

	err := InitializeTables()
	panicErr(err)




}

func defineCustomerObject() {
	customerType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Customer",
		Description: "A customer of the company.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// ToDo: How does this Source work exactly? O_o
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.id, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},
			"surname": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},"address": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},"skill": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},"email": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},"telephone": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},"birthday": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.name, nil
					}
					return nil, nil
				},
			},
		},
	})
}

func defineCustomerSchema() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"customer": &graphql.Field{
				Type: customerType,
				Args: graphql.FieldConfigArgument{
					//Do we need only id?? its to get the whole user, but maybe some selections would be good at customers
					"id": &graphql.ArgumentConfig{
						Description: "id of the user",
						Type:        graphql.NewNonNull(graphql.Int),
					},

				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					//return GetUser(id), nil
					return id, nil
				},
			},
			"customers": &graphql.Field{
				Type: graphql.NewList(customerType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Users are a map, but we need a list!
					userSlice := make([]User, len(Users))
					idx := 0
					for  _, user := range Users {
						userSlice[idx] = user
						idx++
					}
					return userSlice, nil
				},
			},
		},
	})
	customerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

func InitializeTables() (error) {

	db, err := openDatabase()

	//generate customer id
	_, err = db.Query("CREATE TABLE IF NOT EXISTS idNumbers (id serial primary key)")
	panicErr(err)

	closeDatabase(db,nil);

	return err;
}


func Insert(name string, surname string,  street string, number int, zipcode int, city string, skill int, email string, telephone string, birthday time.Time) int{
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

	//define seed for random numbers
	seedSource := rand.NewSource(time.Now().UnixNano())
	randSeed := rand.New(seedSource)

	id := 1000000 + (iteratedIdPart * 100) + randSeed.Intn(99)

	_,err = db.Query("UPDATE idNumbers SET id=$1",iteratedIdPart + 1)


	//create customer database if not existent
	_,err = db.Query("CREATE TABLE IF NOT EXISTS customers (id serial PRIMARY KEY, name text NOT NULL, surname text NOT NULL,street text NOT NULL, number integer NOT NULL,zipcode integer NOT NULL,city text NOT NULL,skill integer NOT NULL,email text NOT NULL,telephone text NOT NULL,birthday text NOT NULL)")
	panicErr(err)

	// Add a customer to it
	_,err = db.Query("INSERT INTO customers(id,name,surname,street, number,zipcode,city,skill,email,telephone,birthday) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
		id, name, surname, street, number, zipcode, city, skill, email, telephone, birthday);
	panicErr(err)

	rows, err = db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	closeDatabase(db,rows);

	return id;
}

func Remove(customer *Customer) {

	db, err := openDatabase()

	//remove users
	_,err = db.Query("DELETE FROM customers WHERE id=$1;", customer.id);
	panicErr(err)

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	closeDatabase(db,rows);
}

func Update(customer *Customer) {

	db, err := openDatabase()

	// update user in database
	_,err = db.Query("UPDATE customers SET name=$2, surname=$3,street=$4 ,number=$5,zipcode=$6,city=$7,skill=$8,email=$9,telephone=$10,birthday=$11 WHERE id = $1;",
		customer.id, customer.name, customer.surname, customer.address.street, customer.address.number, customer.address.zipcode,
		customer.address.city, customer.skill, customer.email, customer.telephone, customer.birthday);
	panicErr(err)

	rows, err := db.Query("SELECT * FROM customers")
	panicErr(err)

	defer rows.Close()

	//additional
	closeDatabase(db,rows);
}

func openDatabase() (*sql.DB,error){
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	panicErr(err)
	panicErr(db.Ping()) //Open does not check the connection

	return db,err
}

func closeDatabase(db *sql.DB, rows *sql.Rows){

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
		fmt.Printf("id: %s, name: %s, surname: %s,street: %s,number: %d,zipcode: %d,city: %s,skill: %d,email: %s,telephone: %s,birthday: %s\n", id, name,surname,street,number,zipcode,city,skill,email,telephone,birthday)
	}

	defer db.Close()
}

func panicErr(err error){
	if err != nil{
		panic(err)
	}
}