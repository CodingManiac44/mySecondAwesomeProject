package person

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
	Sex  string
	Dob  time.Time
}

func (p *Person) greet() {
	fmt.Println("Hello from private method greet:", p.Name)
}

func (p *Person) Meet() {
	fmt.Println("From public method, nice to meet you!")
	p.greet()
}
