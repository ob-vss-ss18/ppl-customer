package Database

import (
	"time"
	_"github.com/lib/pq"
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