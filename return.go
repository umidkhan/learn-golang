package main

import "fmt"

func sayHi(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(sayHi(12, 12))
}
