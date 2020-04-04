package models

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// creating a variable block
var (
	user   []*User // a slice to hold user pointers
	nextID = 1
)
