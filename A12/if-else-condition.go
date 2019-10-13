package main

import "fmt"

func main() {
	var point = 8342.34

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	var angka = 3

	switch angka {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{ //pakai kurung kurawal boleh, tidak pakai juga boleh
			fmt.Println("not bad")
			fmt.Println("belajarlah lebih keras lagi")
		}
	}

	var number = 5

	//switch bisa diperlakukan layaknya if else
	switch {
	case number == 8:
		fmt.Println("perfect")
	case (number < 8) && (number > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("belajaro maneh, le")
		}
	}

	//switch dengan keyword fallthrough
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
