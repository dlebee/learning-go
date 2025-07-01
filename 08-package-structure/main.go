package main

import (
	"fmt"

	"github.com/dlebee/learning-go/08-package-structure/utils"
)

func main() {
	fmt.Println("1 + 2 =", add(1, 2))
	utils.GreetPerson("David Lebee")
}
