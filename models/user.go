package models

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// creating a variable block
var (
	users  []*User // a slice to hold user pointers
	nextID = 1
)

// GetUsers return all users
func GetUsers() []*User {
	return users
}

// AddUser add a new user to users slice passing the user object
func AddUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}
