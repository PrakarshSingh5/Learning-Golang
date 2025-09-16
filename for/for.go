package main

import "fmt"

// for -> only construct in go for looping
func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

}
