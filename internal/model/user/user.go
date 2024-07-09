// Package user. There is user domain definition here.
package user

type User struct {
	Name
	Login string
	ID    int
	Age   int8
}

type Name struct {
	FirstName string
	LastName  string
}

func CreateUser(id int, firstName, lastName string, age int8) *User {
	return &User{
		ID:  id,
		Age: age,
		Name: Name{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}
