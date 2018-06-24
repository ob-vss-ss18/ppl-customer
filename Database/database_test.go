package Database

import (
	"testing"
	"time"
)

var customer = []Customer{
	{name: "test",
	surname: "user",
	address: Address{
		street: "Hauptstraße",
		number: 1337,
		zipcode: 123456,
		city: "VSSCity",
	},
	skill:PRO,
	email: "testuser@vsscity.com",
	telephone: "+49 1762 1337331",
	birthday: time.Now().Local(),
	},

	{name: "Cooler",
	surname: "Name",
	address: Address{
		street: "Nebenstraße",
		number: 80,
		zipcode: 80082,
		city: "Noob City",
	},
	skill:BEGINNER,
	email: "cooler-user@noobcity.com",
	telephone: "+49 1762 7862334",
	birthday: time.Now().Local(),
	},
}

func TestInsert(t *testing.T) {
	InitializeCustomerDB()
	currentCustomer := customer[0]
	id := InsertCustomer(currentCustomer)
	createdCustomer, _ := Select(id)


	if createdCustomer.name != currentCustomer.name{
		t.Errorf("Tried to insert name %s, but got %s", currentCustomer.name, createdCustomer.name)
	}
	if createdCustomer.surname != currentCustomer.surname{
		t.Errorf("Tried to insert surname %s, but got %s", currentCustomer.surname, createdCustomer.surname)
	}
	if createdCustomer.address.street != currentCustomer.address.street{
		t.Errorf("Tried to insert street %s, but got %s", currentCustomer.address.street, createdCustomer.address.street)
	}
	if createdCustomer.address.number != currentCustomer.address.number{
		t.Errorf("Tried to insert number %d, but got %d", currentCustomer.address.number, createdCustomer.address.number)
	}
	if createdCustomer.address.zipcode != currentCustomer.address.zipcode{
		t.Errorf("Tried to insert zipcode %d, but got %d", currentCustomer.address.zipcode, createdCustomer.address.zipcode)
	}
	if createdCustomer.address.city != currentCustomer.address.city{
		t.Errorf("Tried to insert city %s, but got %s", currentCustomer.address.city, createdCustomer.address.city)
	}
	if createdCustomer.skill != currentCustomer.skill{
		t.Errorf("Tried to insert skill %v, but got %v", currentCustomer.skill, createdCustomer.skill)
	}
	if createdCustomer.email != currentCustomer.email{
		t.Errorf("Tried to insert email %s, but got %s", currentCustomer.email, createdCustomer.email)
	}
	if createdCustomer.telephone != currentCustomer.telephone{
		t.Errorf("Tried to insert telephone %s, but got %s", currentCustomer.telephone, createdCustomer.telephone)
	}
}

func TestRemove(t *testing.T) {
	InitializeCustomerDB()
	currentCustomer := customer[0]
	id := InsertCustomer(currentCustomer)
	createdCustomer, _ := Select(id)

	if createdCustomer.id == 0 {
		t.Errorf("Insertion failed")
	}

	Remove(&createdCustomer)

	createdCustomer, _ = Select(createdCustomer.id)

	if createdCustomer.id != 0 {
		t.Errorf("Remove failed")
	}

}

func TestUpdate(t *testing.T) {
	InitializeCustomerDB()
	currentCustomer := customer[0]
	id := InsertCustomer(currentCustomer)
	createdCustomer, _ := Select(id)

	if createdCustomer.id == 0 {
		t.Errorf("Insertion failed")
	}

	currentCustomer = customer[1]
	currentCustomer.id = id

	Update(&currentCustomer)

	updatedCustomer, _ := Select(id)

	if updatedCustomer.id != currentCustomer.id {
		t.Errorf("Got ID %d but it should be ID %d", updatedCustomer.id, currentCustomer.id)
	}

	if updatedCustomer.name != currentCustomer.name{
		t.Errorf("Tried to update name %s, but got %s", currentCustomer.name, updatedCustomer.name)
	}
	if updatedCustomer.surname != currentCustomer.surname{
		t.Errorf("Tried to update surname %s, but got %s", currentCustomer.surname, updatedCustomer.surname)
	}
	if updatedCustomer.address.street != currentCustomer.address.street{
		t.Errorf("Tried to update street %s, but got %s", currentCustomer.address.street, updatedCustomer.address.street)
	}
	if updatedCustomer.address.number != currentCustomer.address.number{
		t.Errorf("Tried to update number %d, but got %d", currentCustomer.address.number, updatedCustomer.address.number)
	}
	if updatedCustomer.address.zipcode != currentCustomer.address.zipcode{
		t.Errorf("Tried to update zipcode %d, but got %d", currentCustomer.address.zipcode, updatedCustomer.address.zipcode)
	}
	if updatedCustomer.address.city != currentCustomer.address.city{
		t.Errorf("Tried to update city %s, but got %s", currentCustomer.address.city, updatedCustomer.address.city)
	}
	if updatedCustomer.skill != currentCustomer.skill{
		t.Errorf("Tried to update skill %v, but got %v", currentCustomer.skill, updatedCustomer.skill)
	}
	if updatedCustomer.email != currentCustomer.email{
		t.Errorf("Tried to update email %s, but got %s", currentCustomer.email, updatedCustomer.email)
	}
	if updatedCustomer.telephone != currentCustomer.telephone{
		t.Errorf("Tried to update telephone %s, but got %s", currentCustomer.telephone, updatedCustomer.telephone)
	}

}