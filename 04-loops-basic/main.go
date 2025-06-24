package main

import "fmt"

func main() {
	fmt.Println("basic loops, in GoLang all loops use a for loop keyword with variations of ways of using it..")

	fmt.Println("1. Traditional for loops")
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("\n2. While loops")
	j := 0
	for j < 3 {
		fmt.Println("j = ", j)
		j++
	}

	fmt.Println("\n3. Infinite loop with a break")
	k := 0
	for {

		fmt.Println("k = ", k)

		k += 2

		if k == 10 {
			break
		}
	}
}
