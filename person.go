package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
	sex  string
	dob  time.Time
}

// name start with
func (p *person) greet() {
	fmt.Println("Hello from private method greet: no1 Dulabhai", p.name)
}

// public method (accessible from other packages)
func (p *person) Meet() {
	fmt.Println("from public method ,Nice to meet you!")
}
