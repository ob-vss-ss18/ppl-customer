package Database

import "github.com/graphql-go/graphql"

var(
	CustomerSchema graphql.Schema
	CustomerType *graphql.Object
	Customers      map[int]Customer
)

func defineCustomerObject() {
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
					//return GetUser(id), nil
					return id, nil
				},
			},
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
	CustomerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

