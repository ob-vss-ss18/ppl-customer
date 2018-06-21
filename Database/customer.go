package Database

import (
	"time"

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