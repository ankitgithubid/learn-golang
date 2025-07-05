package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `validate:"min-2,max=32"`
	Email string `validate:"required,email"`
}

func validate(val interface{}) error {
	v := reflect.ValueOf(val)
	t := reflect.TypeOf(val)

	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			fieldValue := v.Field(i)
			fieldType := t.Field(i)
			tag := fieldType.Tag.Get("validate")

			fmt.Printf("Field %d: %s = %v, tag = %s\n", i, fieldType.Name, fieldValue, tag)
		}
	} else {
		return fmt.Errorf("expected struct, got %s", v.Kind())
	}

	return nil
}

func main() {
	user := User{
		Name:  "John Doe",
		Email: "john@test.com",
	}

	fmt.Printf("Validate = %v\n", validate(user))

	//t := reflect.TypeOf(user)
	//fmt.Printf("Name : %s Kind of user = %s\n", t.Name(), t.Kind())
	//
	//for i := 0; i < t.NumField(); i++ {
	//	field := t.Field(i)
	//	tag := field.Tag.Get("validate")
	//	fmt.Printf("Field %s, tag = %s\n", field.Name, tag)
	//}
}
