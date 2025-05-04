package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `validate:"min-2,max=32"`
	Email string `validate:"required,email"`
}

func main() {
	user := User{
		Name:  "John Doe",
		Email: "john@test.com",
	}

	t := reflect.TypeOf(user)
	fmt.Printf("Name : %s Kind of user = %s\n", t.Name(), t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")
		fmt.Printf("Field %s, tag = %s\n", field.Name, tag)
	}
}
