package main

import (
	"fmt"
)

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fiboacci(i)
		fmt.Printf("fiboacci(%d) is: %d\n", i, result)
	}
}

func fiboacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fiboacci(n-1) + fiboacci(n-2)
	}
	return
}
