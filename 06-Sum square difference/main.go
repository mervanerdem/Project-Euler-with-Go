/*
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import (
	"fmt"
)

func main() {
	sum := 0
	sum2 := 0

	for i := 1; i <= 100; i++ {
		sum = sum + i*i
		sum2 = sum2 + i
	}
	sum2Pow := sum2 * sum2

	fmt.Println(sum2Pow - sum)
}
