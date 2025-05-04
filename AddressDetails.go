package main

type Address struct {
	house   string
	city    string
	pincode int
}

func NewAddress() *Address {
	return &Address{}
}

func (a *Address) House() string {
	return a.house
}

func (a *Address) SetHouse(house string) {
	a.house = house
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) SetCity(city string) {
	a.city = city
}

func (a *Address) Pincode() int {
	return a.pincode
}

func (a *Address) SetPincode(pincode int) {
	a.pincode = pincode
}
