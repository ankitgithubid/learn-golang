package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	u1 := NewUserDetail()
	u1.SetName("Ankit")
	u1.SetAge(25)

	u1.SetHouse("206")
	u1.SetCity("Gurgaon")
	u1.SetPincode(122101)

	fmt.Println("User Name:", u1.Name())
	fmt.Println("User Age:", u1.Age())
	fmt.Println("User House:", u1.House())
	fmt.Println("User City:", u1.City())
	fmt.Println("User Pincode:", u1.Pincode())
}
