package main

import (
	"fmt"
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

func printPrimes(n int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// don't edit below this line

func test(n int) {
	fmt.Printf("Primes up to %v:\n", n)
	printPrimes(n)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(50)
}
