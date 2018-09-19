package main

import (
	"fmt" 
	"strconv"
)

// Define Person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (val receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main(){
	// Init Person using struct
	person1 := Person {
		firstName: "David",
		lastName: "Agaybi",
		city: "Mississauga",
		gender: "M",
		age: 42,

	}

	person2 := Person {
		firstName: "Abelina",
		lastName: "Adams",
		city: "Mississ",
		gender: "F",
		age: 35,
	}
	
	// Alternative
	// person1 := Person{"David", "Agaybi", "Mississauga", "M", 42}

	// fmt.Println(person1)
	// person1.age++
	// fmt.Println(person1.firstName)

	person1.hasBirthday()
	person2.getMarried("Hayes")
	
	fmt.Println(person2.greet())



}