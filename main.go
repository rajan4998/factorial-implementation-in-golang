package main

import "fmt"

func main() {
	fmt.Println("==Factorial==")
	factorial(5)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(" * ", n)
	return n * factorial(n-1)
}
