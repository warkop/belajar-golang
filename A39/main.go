package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("start")
	// time.Sleep(time.Second * 4)
	// fmt.Println("after 4 seconds")

	// for true {
	// 	fmt.Println("Hello!!")
	// 	time.Sleep(1 * time.Second)
	// }

	var time = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-time.C
	fmt.Println("finish")
}
