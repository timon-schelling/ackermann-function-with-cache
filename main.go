package main

import (
	"fmt"
)

func ackermann(a, b, n uint64) (c uint64) {
	if n <= 0 {
		return a + b
	} else if n == 1 {
		return a * b
	} else if n >= 2 {
		if b <= 1{
			return a
		}
		return ackermann(ackermann(a, b-1, n), a, n-1)
	}
	return 0
}

func main() {
	fmt.Println(ackermann(3, 3, 4))
}

