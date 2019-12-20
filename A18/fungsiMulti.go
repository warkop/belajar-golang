package main

import "fmt"
import "math"

func calculated(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d / 2,2)

	//hitung keliling
	var circumference = math.Pi * d

	//kembalikan 2 nilai
	return area, circumference
}

//bisa juga seperti dibawah ini, jadi variabelnya langsung didefinisikan diawal
// func calculate(d float64) (area float64, circumference float64) {
//     area = math.Pi * math.Pow(d / 2, 2)
//     circumference = math.Pi * d

//     return
// }

func main() {
	var diameter float64 = 15

	//nama variabel penerima bebas, tidak harus sama seperti di fungsi
	var area, circumference = calculated(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t\t: %.2f \n", circumference)
}