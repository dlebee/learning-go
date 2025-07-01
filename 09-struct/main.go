package main

import "fmt"

func main() {

	david := Person{
		Name: "David Lebee",
		Age:  36,
	}

	// not recommended to use order building, could be dangerous when introducing new fields to the struct.
	johnDoe := Person{"John doe", 25}

	david.Greet()
	johnDoe.Greet()

	fmt.Println("David's having a birthday")
	david.haveBirthday()
	david.Greet()

	fmt.Println("Combined age of david and john is", david.CombinedAge(johnDoe))
}
