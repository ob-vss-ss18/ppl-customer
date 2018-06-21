package Database

import (
	"github.com/graphql-go/graphql"

	"time"
)

var(
	CustomerSchema graphql.Schema
	CustomerType *graphql.Object
	AddressType *graphql.Object
	Customers      map[int]Customer
)



func InitGraphQL(){
	defineCustomerObject()
	defineCustomerSchema()
}

func defineCustomerObject() {
	AddressType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Address",
		Description: "An address of a customer",
		Fields: graphql.Fields{
			//Maybe we don't need this, please check somebody
			"address":&graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The address.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					address, ok := p.Source.(Address)
					if ok {
						return address, nil
					}
					return nil, nil
				},
			},
			"street": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The street of the address.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					address, ok := p.Source.(Address)
					if ok {
						return address.street, nil
					}
					return nil, nil
				},
			},"number": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The number at the street of the address.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					address, ok := p.Source.(Address)
					if ok {
						return address.number, nil
					}
					return nil, nil
				},
			},"zip": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The zip code of the address.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					address, ok := p.Source.(Address)
					if ok {
						return address.zipcode, nil
					}
					return nil, nil
				},
			},"city": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The city of the address.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					address, ok := p.Source.(Address)
					if ok {
						return address.city, nil
					}
					return nil, nil
				},
			},
		},
	})


	CustomerType = graphql.NewObject(graphql.ObjectConfig{
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
				Description: "The surname of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.surname, nil
					}
					return nil, nil
				},
			},"address": &graphql.Field{
				Type:   AddressType,
				Description: "The address of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer)
					if ok {
						return customer.address, nil
					}
					return nil, nil
				},
			},"skill": &graphql.Field{
				Type:        graphql.Int,
				Description: "The name of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.skill, nil
					}
					return nil, nil
				},
			},"email": &graphql.Field{
				Type:        graphql.String,
				Description: "The email of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.email, nil
					}
					return nil, nil
				},
			},"telephone": &graphql.Field{
				Type:        graphql.String,
				Description: "The telephone of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						return customer.telephone, nil
					}
					return nil, nil
				},
			},"birthday": &graphql.Field{
				Type:        graphql.String,
				Description: "The birthday of the customer.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					customer, ok := p.Source.(Customer);
					if ok {
						//TODO may format this
						birthday :=customer.birthday.String()
						return birthday, nil
					}
					return nil, nil
				},
			},
		},
	})
}

//mutation {
//  create(input: {name: "blub", surname: "blub", email: "abc@de.y", telephone: "0128326548", skill: 1, address: {street: "bbb", number: 12, zip: 555662, city: "sagewg"}, birthday:"02-03-1994"}) {
//    id
//    name
//  }
//}

func defineCustomerSchema() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"customer": &graphql.Field{
				Type: CustomerType,
				Args: graphql.FieldConfigArgument{
					//Do we need only id?? its to get the whole user, but maybe some selections would be good at customers
					"id": &graphql.ArgumentConfig{
						Description: "id of the user",
						Type:        graphql.NewNonNull(graphql.Int),
					},

				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					return Select(id), nil
				},
			},
			"address": &graphql.Field{
				Type: AddressType,

			},

			//TODO chris mach da, keine lust :P
			"customers": &graphql.Field{
				Type: graphql.NewList(CustomerType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Users are a map, but we need a list!
					customerSlice := make([]Customer, len(Customers))
					idx := 0
					for  _, user := range Customers {
						customerSlice[idx] = user
						idx++
					}
					return customerSlice, nil
				},
			},
		},
	})

	createAddress := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateAddress",
		Fields: graphql.InputObjectConfigFieldMap{
			"street" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
				Description: "street of the customer",
			},"number" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
				Description: "number of the customer",
			},"zip" : &graphql.InputObjectFieldConfig{
				Type:graphql.NewNonNull(graphql.Int),
				Description: "zip of the customer",
			},"city" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "city of the customer",
			},
		},
	})
	createCustomer := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateCustomer",
		Fields:  graphql.InputObjectConfigFieldMap{
			"name" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
				Description: "name of the customer",
			},"surname" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
				Description: "name of the customer",
			},"address" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(createAddress),
				Description: "name of the customer",
			},"skill" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
				Description: "name of the customer",
			},"email" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
				Description: "name of the customer",
			},"telephone" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
				Description: "name of the customer",
			},"birthday" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
				Description: "name of the customer",
			},
		},
	})

	argsCreate := graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description:"An input with the customer details",
			Type: graphql.NewNonNull(createCustomer),
		},
	}

	updateAddress := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UpdateAddress",
		Fields: graphql.InputObjectConfigFieldMap{
			"street" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "street of the customer",
			},"number" : &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
				Description: "number of the customer",
			},"zip" : &graphql.InputObjectFieldConfig{
				Type:graphql.Int,
				Description: "zip of the customer",
			},"city" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "city of the customer",
			},
		},
	})
	updateCustomer := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UpdateCustomer",
		Fields:  graphql.InputObjectConfigFieldMap{
			"id" : &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
				Description: "name of the customer",
			},"name" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "name of the customer",
			},"surname" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "name of the customer",
			},"address" : &graphql.InputObjectFieldConfig{
				Type: updateAddress,
				Description: "name of the customer",
			},"skill" : &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
				Description: "name of the customer",
			},"email" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "name of the customer",
			},"telephone" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "name of the customer",
			},"birthday" : &graphql.InputObjectFieldConfig{
				Type: graphql.String,
				Description: "name of the customer",
			},
		},
	})

	argsUpdate := graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description:"An input with the customer details",
			Type: graphql.NewNonNull(updateCustomer),
		},
	}

	removeCustomer := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "RemoveCustomer",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "street of the customer",
			},
		},
	})

	argsRemove := graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Description:"An input with the customer details",
			Type: removeCustomer,

		},
	}
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"update": &graphql.Field{
				Type: CustomerType,
				Args: argsUpdate,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var inp= p.Args["input"].(map[string]interface{})

					var addressInp= inp["address"].(map[string]interface{})

					var street, city string
					var number, zipcode, id int

					var address Address
					var customer Customer

					if inp["id"] == nil {
						return nil, nil
					} else {
						id = inp["id"].(int)
					}

					if addressInp["street"] != nil {
						street = addressInp["street"].(string)
					}
					if addressInp["number"] != nil {
						number = addressInp["number"].(int)
					}
					if addressInp["zip"] != nil {
						zipcode = addressInp["zip"].(int)
					}
					if addressInp["city"] != nil {
						city = addressInp["city"].(string)
					}
					address = Address{

						street:  street,
						number:  number,
						zipcode: zipcode,
						city:    city,
					}

					customer.address = address

					if inp["name"] != nil {
						customer.name = inp["name"].(string);
					}
					if inp["surname"] != nil {
						customer.surname = inp["surname"].(string);
					}
					if inp["email"] != nil {
						customer.email = inp["email"].(string);
					}
					if inp["telephone"] != nil {
						customer.telephone = inp["telephone"].(string);
					}
					if inp["skill"] != nil {
						customer.skill = Skill(inp["skill"].(int));
					}
					if inp["birthday"] != nil {
						birthday, _ := time.ParseInLocation(time.ANSIC, inp["birthday"].(string), time.Local)
						customer.birthday = birthday
					}

					Update(&customer)

					return Select(id), nil
				},
			},

			"create": &graphql.Field{
				Type: CustomerType,
				Args: argsCreate,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var inp= p.Args["input"].(map[string]interface{})

					var addressInp= inp["address"].(map[string]interface{})
					address := Address{

						street:  addressInp["street"].(string),
						number:  addressInp["number"].(int),
						zipcode: addressInp["zip"].(int),
						city:    addressInp["city"].(string),
					}

					//TODO don't know how to parse this correctly
					birthday, _ := time.ParseInLocation(time.ANSIC, inp["birthday"].(string), time.Local)

					customerToCreate := Customer{
						name:      inp["name"].(string),
						surname:   inp["surname"].(string),
						address:   address,
						skill:     Skill(inp["skill"].(int)),
						email:     inp["email"].(string),
						telephone: inp["telephone"].(string),
						birthday:  birthday,
					}

					customerToCreate.id = InsertCustomer(customerToCreate)
					return customerToCreate, nil
				},
			},

			"remove": &graphql.Field{
				Type: CustomerType,
				Args: argsRemove,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var inp = p.Args["input"].(map[string]interface{})

					customer := Select(inp["id"].(int))

					Remove(&customer)

					return customer, nil
				},
			},
		},
	})


	CustomerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	})
}

