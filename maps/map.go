package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "golang"
	
	m["area"] = "backend"

	//seting an element
	fmt.Println(m)


	// for matching the two map we have a different function which is 
	// maps.Equal(m1, m2)
}
