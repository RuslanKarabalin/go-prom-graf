package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	array := make([]int, 0, n)

	for range n {
		var v int

		fmt.Fscan(in, &v)
		array = append(array, v)
	}

	for i, v := range array {
		fmt.Printf("%v", v)
		if i+1 != n {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
