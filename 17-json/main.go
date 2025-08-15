package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {

	person := Person{
		Name:    "David Lebee",
		Age:     36,
		IsAdult: true,
	}

	fmt.Println("Hello world!", person)

	result, isError := json.Marshal(person)

	if isError != nil {
		fmt.Println("failed to serialize person:", isError)
	} else {
		fmt.Println("serialized person", string(result))
	}
}
