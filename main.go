package main

import (
	"fmt"
	"os"
	"strconv"
)

func fac(n int) int {
	if n == 1 {
		return 1
	}
	return n * fac(n-1)
}

func main() {

	valueUI := os.Getenv("RECURSIVE_UI")

	n, _ := strconv.Atoi(valueUI)

	valueString := strconv.Itoa(fac(n))

	fmt.Printf("result %s\n", valueString)
}
