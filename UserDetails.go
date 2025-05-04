package main

type UserDetail struct {
	name string
	age  int
	Address
}

func NewUserDetail() *UserDetail {
	return &UserDetail{
		Address: Address{},
	}
}

func (u *UserDetail) Name() string {
	return u.name
}

func (u *UserDetail) Age() int {
	return u.age
}

func (u *UserDetail) SetName(name string) {
	u.name = name
}

func (u *UserDetail) SetAge(age int) {
	u.age = age
}
