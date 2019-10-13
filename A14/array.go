package main

import "fmt"

func main() {
	//deklarasi array biasa
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	//inisialisasi nilai awal array
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	//inisialisasi nilai awal array tanpa jumlah elemen
	// var numbers = [...]int{2, 3, 2, 4, 3}

	// numbers[1] = 8

	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))

	//array multidimensi
	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{
	// 	{3, 2, 3}, 
	// 	{3, 4, 5},
	// }

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	//perulangan array menggunakan for
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	//perulangan array menggunakan keyword for-range
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	//penggunaan underscore pada perulangan digunakan apabila kita hanya ingin menampilkan elemennya saja, indexnya tidak diperlukan, ini sama seperti foreach yang mana tidak memerlukan variabel incremental
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for _, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// }

	//alokasi elemen array menggunakan keyword make
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits)  // [apple manggo]
}