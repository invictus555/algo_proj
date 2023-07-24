package main

import "fmt"

func main() {
	var array []string

	array = append(array, "hello")
	array = append(array, "world")
	array = append(array, "....")
	fmt.Println(array)

	array = array[0:2]
	fmt.Println(array)

	array = append(array, "!")
	fmt.Println(array)
}
