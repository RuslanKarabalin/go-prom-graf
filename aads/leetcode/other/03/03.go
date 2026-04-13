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

	arr := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &arr[i])
	}

	var k, m int
	found := 0
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] != 2 {
			if found == 0 {
				k = i
			} else {
				m = i
			}
			found++
		}
	}
	if found == 1 {
		m = k
	}

	fmt.Printf("%d %d\n", k, m)
}
