package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata: %.2f", avg)

	//cara mengirim kedalam variabel biasa dan variadic, parameter pertama adalah parameter biasa, selanjutnya adalah parameter variadic
	yourHobbies("Moden", "sepakbola", "basket", "bersepeda")

	//bisa juga menggunakan slice untuk dikirim
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)

	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

//fungsi dengan parameter biasa dan variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
