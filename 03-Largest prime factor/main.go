/*
Largest Prime Faktor
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	val := 600851475143
	for i := 1; i < val; i++ {
		if val%i == 0 {
			val /= i
		}
	}
	fmt.Println(val)
}
