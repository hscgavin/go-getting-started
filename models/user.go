package models

// User is ...
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  = []*User{}
	nextID = 0
)
