package Database

var counter int = 0

type Person struct {
	ID int
	Name string
}
func CreatePerson(personName string) *Person{
	shipID := counter
	counter++
	newPerson := &Person{
		ID:shipID,
		Name:personName,
		}

	return newPerson
}