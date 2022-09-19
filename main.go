package main

import (
	"fmt"

	"github.com/sardarmd/mockgendemo/domain"
)

func main() {
	user := domain.User{
		Id: "123", Name: "Sardar", City: "Bangalore",
	}
	stub := domain.NewUserStub()
	stub.AddUser(user)

	usr := stub.FindUserByPosition(0)
	fmt.Printf("User fetched --- %v", usr)
}
