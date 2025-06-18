package main

import (
	"fmt"
	"mySecondAwesomeProject/person"
	"time"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Println("Main running")

	// Call Try() from new_package
	//new_package.Try()

	fmt.Printf("welcome to second phase \n")

	// Create an instance of the exported Person struct from the 'person' package
	person1 := person.Person{ // Use person.Person
		Name: "sabbir dulabhai sends his salami from person package",
		Age:  30,
		Sex:  "male",
		Dob:  time.Date(1995, 04, 15, 0, 0, 0, 0, time.UTC),
	}

	// You cannot call person1.greet() directly here because 'greet' is a private method
	// within the 'person' package. It can only be called from within the 'person' package itself.
	// You should call the public method 'Meet()' which, if designed to,
	// can internally call 'greet()'.
	person1.Meet() // Call the public method

	// If you want to see the private greet being called, ensure person.Meet() calls person.greet() internally:
	// In person.go:
	// func (p *Person) Meet() {
	//     p.greet() // This calls the private method within the same package
	//     fmt.Println("from public method ,Nice to meet you!")
	// }
}
