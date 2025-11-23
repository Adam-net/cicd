package main

import "fmt"

func main() {
	fmt.Println(hello())
	fmt.Println("It's alive!")
}

func hello() string {
	return "Hello Go"
}
