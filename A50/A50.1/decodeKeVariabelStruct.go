package main

import (
	"encoding/json"
	"fmt"
)

// User is ...
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user: ", data.FullName)
	fmt.Println("age: ", data.Age)

	//decode JSON ke map[string]interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user : ", data1["Name"])
	fmt.Println("age : ", data1["Age"])

	//decode JSON ke interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user : ", decodedData["Name"])
	fmt.Println("age : ", decodedData["Age"])
}
