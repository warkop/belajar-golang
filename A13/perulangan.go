package main

import "fmt"

// func main() {
// 	//for biasa
// 	for i := 0; i < 5; i++ {
// 		// fmt.Println("tod")
// 	}

// 	//for dengan argumen dan hanya kondisi, sama seperti while
// 	k := 0
// 	for k < 6 {
// 		// fmt.Println(k)
// 		k++
// 	}

// 	//for tanpa argumen
// 	j := 0
// 	for {
// 		// fmt.Println(j)

// 		j++
// 		if j == 5 {
// 			break
// 		}
// 	}

// 	//for range
// 	for i := 0; i <= 10; i++ {
// 		if i%2 == 1 {
// 			continue
// 		}

// 		if i > 7 {
// 			break
// 		}

// 		fmt.Println("Angka", i)
// 	}

// 	//pemanfaatan label dalam perulangan
// 	//dibawah ini adalah label
// outerLoop:
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			if i == 3 {
// 				break outerLoop
// 			}
// 			fmt.Print("matriks [", i, "][", j, "]", "\n")
// 		}
// 	}
// }

func main() {
	count := 0
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if i%2 == 0 || j%2 != 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
