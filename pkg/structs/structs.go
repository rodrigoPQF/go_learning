package structs

import "fmt"

// Struct é usado para criar objetos 💣
type User struct {
	name string
	age  int
}

// Struct com campos privados 🦹
type UserLogin struct {
	Email    string
	password int // Esse campo não aparece
}

func Structs() {
	var user User = User{
		name: "Rodrigo",
		age:  22,
	}

	var userLogin UserLogin = UserLogin{
		Email:    "mail",
		password: 123,
	}

	fmt.Println(user)
	fmt.Println(userLogin)

}
