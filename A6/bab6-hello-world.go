package main

import "fmt"

func main() {
	//dibawah ini adalah variabel tampung, tidak bisa ditampilkan, hanya untuk mengirim nilai balik
	_ = "aneh"

	var angka1 = 34
	var angka2 = 23

	//pointer
	nama := new(string)
	*nama = "tod"

	fmt.Println("halo %s %s!\n", angka1+angka2)
	fmt.Println(*nama)
}
