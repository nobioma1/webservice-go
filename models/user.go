package models

import (
	"errors"
	"fmt"
)

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
	if user.ID != 0 {
		return User{}, errors.New("User must not include ID")
	}

	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}

// GetUserByID return the user struct with id
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with '%v' not found", id)
}

// UpdateUserByID updates a user with matching id
func UpdateUserByID(update User) (User, error) {
	for index, user := range users {
		if user.ID == update.ID {
			users[index] = &update
			return update, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", update.ID)
}

// RemoveUserByID removes user with ID
func RemoveUserByID(id int) (int, error) {
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			return id, nil
		}
	}
	return -1, fmt.Errorf("User with '%v' not found", id)
}
