package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{"John Wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("Nama panggilan : ", name)
}
