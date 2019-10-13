package main

import "fmt"

func main() {
	var data map[string]int
	data["one"] = 1
	// akan muncul error!

	data = map[string]int{}
	data["one"] = 1

	fmt.Println(data)
}