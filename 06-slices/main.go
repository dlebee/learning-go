package main

import "fmt"

func main() {

	nums := []int{5, 10, 15}

	fmt.Println("List of numbers")
	for i, v := range nums {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	nums = append(nums, 20, 25) // this replaces the reference with a new allocation
	// this is somewhat equivalent with malloc and copy into into a single line
	// of what we use to do back in the 90's
	fmt.Println("Altered slices of numbers")
	for i, v := range nums {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	patial_slice := nums[3:5]
	fmt.Println("2:4 slice of the nums array")
	for i, v := range patial_slice {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
