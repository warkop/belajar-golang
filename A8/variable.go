package main
import "fmt"

func main()  {
	//beserta tipe variabel
	// var firstName string = "john"
	// var lastName string
	// lastName = "wick"
	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	// tanpa tipe variabel
	// var firstName string = "john"
	// lastName := "wick"
	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	//variabel underscore
	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"
	// fmt.Printf("halo %s!\n", name)

	//deklarasi menggunakan keyword new
	name := new(string)

	fmt.Println(name)   // 0x20818a220
	fmt.Println(*name)  // ""
	// fmt.Printf("halo %s!\n", name)
}