package main

import (
	"github.com/graphql-go/graphql"
)

//
// To actual query this data send a POST to the endpoint (specified in main) with a json body e.g:
// EXAMPLE 1:
// {
// 	"query":"{users{id name}}"
// }
// EXAMPLE 2:
// {
// 	"query":"{user(id:1002){id name}}"
// }
//

var (
	Thomas User
	Lukas  User
	Oli    User
	Basti  User
	Chris  User

	Users      map[int]User
	UserSchema graphql.Schema

	userType *graphql.Object
)

type User struct {
	ID   int
	Name string
}

func InitializeUserDB() {
	// Just initializing some data, later DB should be used!
	Thomas = User{
		ID:   1000,
		Name: "Thomas",
	}
	Lukas = User{
		ID:   1001,
		Name: "Lukas",
	}
	Oli = User{
		ID:   1002,
		Name: "Oli",
	}
	Basti = User{
		ID:   1003,
		Name: "Basti",
	}
	Chris = User{
		ID:   1004,
		Name: "Chris",
	}

	Users = map[int]User{
		1000: Thomas,
		1001: Lukas,
		1002: Oli,
		1003: Basti,
		1004: Chris,
	}
	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A user which is a customer of the company.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// ToDo: How does this Source work exactly? O_o
					user, ok := p.Source.(User);
					if ok {
						return user.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user, ok := p.Source.(User);
					if ok {
						return user.Name, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the user",
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					return GetUser(id), nil
				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
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
	UserSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

func GetUser(id int) User {
	if user, ok := Users[id]; ok {
		return user
	}
	return User{}
}
