package main

import (
	"fmt"
	"os"
	"strconv"
)

func fac(n int) int {
	if n == 0 {
		return 1
	}
	return n * fac(n-1)
}

func main() {

	inputUI := os.Getenv("RECURSIVE_UI")

	requestToInt, _ := strconv.Atoi(inputUI)

	responseToString := strconv.Itoa(fac(requestToInt))

	fmt.Printf("result %s\n", responseToString)
}
