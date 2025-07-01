package utils

import "fmt"

func greetWithMessage(message string) {
	fmt.Printf("hello %s\n", message)
}

func GreetPerson(fullname string) {
	greetWithMessage(fullname)
}
