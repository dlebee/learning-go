package main

import "fmt"

func greet() {
	fmt.Println("Hello!")
}

func greet_with_name(name string) {
	fmt.Printf("Hello %s!\n", name)
}

func add(left int, right int) int {
	return left + right
}

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

func reverse_array(arr []int) []int {
	number_of_records := len(arr)
	result := make([]int, number_of_records)

	for i, j := 0, number_of_records-1; i < number_of_records; i, j = i+1, j-1 {
		result[i] = arr[j]
	}

	return result
}

func reverse_in_place(arr []int) []int {

	number_of_elements := len(arr)
	for i := 0; i < number_of_elements/2; i = i + 1 {
		mid := arr[i]
		arr[i] = arr[number_of_elements-i-1]
		arr[number_of_elements-i-1] = mid
	}

	return arr
}

func main() {
	greet()
	greet_with_name("David Lebee")
	fmt.Println("2 + 3 = ", add(2, 3))
	fmt.Println("sum of 1, 5, 3 is ", sum(1, 5, 3))

	numbers := []int{5, 2, 7, 9, 8}
	fmt.Println("new array | The reverse of", numbers, " is ", reverse_array(numbers))
	numbers2 := []int{5, 2, 7, 9, 8}
	fmt.Print("in place  | The reverse of ", numbers2)
	reverse_in_place(numbers2)
	fmt.Println(" is", numbers2)
}
