package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	if isPrime(n) {
		fmt.Println(1)
		fmt.Println(n)
		return
	}

	for p := 2; p <= n/2; p++ {
		if isPrime(p) && isPrime(n-p) {
			fmt.Println(2)
			fmt.Println(p, n-p)
			return
		}
	}

	for p := 2; p <= (n-3)/2; p++ {
		if isPrime(p) && isPrime(n-3-p) {
			fmt.Println(3)
			fmt.Println(3, p, n-3-p)
			return
		}
	}
}
