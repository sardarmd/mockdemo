package domain

import "fmt"

var usersList []User

type UserStub struct {
	users []User
}

func (stub UserStub) AddUser(user User) {

	usersList = append(usersList, user)
	fmt.Printf("users %v", usersList)

}

func (stub UserStub) FindUserByPosition(index int64) User {
	return usersList[index]
}

func NewUserStub() UserStub {
	return UserStub{}
}
