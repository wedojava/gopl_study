package main

import "fmt"

func max(nums ...int) int {
	m := 0
	if len(nums) == 0 {
		return m
	}
	m = nums[0]
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}

func min(nums ...int) int {
	m := 0
	if len(nums) == 0 {
		return m
	}
	m = nums[0]
	for _, n := range nums {
		if m > n {
			m = n
		}
	}
	return m
}

func main() {
	print("max() test start:\n")
	fmt.Println("It should be 0! we got:", max())
	fmt.Println("It should be 3! we got:", max(3))
	fmt.Println("It should be 4! we got:", max(-5, 2, 3, 4))
	print("min() test start:\n")
	fmt.Println("It should be 0! we got:", min())
	fmt.Println("It should be 3! we got:", min(3))
	fmt.Println("It should be -5! we got:", min(2, 1, 3, 4, -5))

}
