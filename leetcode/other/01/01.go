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
	var k int
	fmt.Fscan(in, &k)

	i := 0
	for range n {
		var v int

		fmt.Fscan(in, &v)
		if v >= k {
			i++
		}
	}

	fmt.Println(i)
}
