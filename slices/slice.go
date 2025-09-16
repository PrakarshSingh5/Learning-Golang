package main

import "fmt"

func main() {
	//uninitialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil)

	var nums = make([]int, 2)
	fmt.Println(nums)

}
