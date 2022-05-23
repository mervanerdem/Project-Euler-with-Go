/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
package main

import "fmt"

func SieveOfEratosthenes(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}
	for p := 2; p*p < n; p++ {
		if integers[p] {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	primes := SieveOfEratosthenes(2000000)
	fmt.Println(len(primes))
	sum := 0
	for _, prime := range primes {
		sum1 := sum + prime
		if sum1 < sum {
			fmt.Println(sum)
		} else {
			sum = sum1
		}
	}
	fmt.Println("Result:", sum)
}
