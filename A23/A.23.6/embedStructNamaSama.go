package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21        // age of student
	s1.person.age = 22 //	age of person

	//bisa juga seperti ini pengisian nilainya
	// var p1 = person{name: "wick", age: 21}
	// var s1 = student{person: p1, grade: 2}

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
}
