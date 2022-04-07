/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
	"math"
)

func pythagorean() {
	a := 3.0
	b := 4.0
	c := 1.0
	for i := 1; i < 334; i++ { //a<b<c, a max can be 1000/3 = 333.3333
		a++
		for j := a; j < 500; j++ { //a<b<c, b max can be 1000/2 = 500
			b++
			c = math.Sqrt(a*a + b*b)
			if (a + b + c) == 1000.0 {
				break
			}
		}
		if (a + b + c) == 1000.0 {
			break
		}
		b = a
		c = 1
	}
	productnum := int(a * b * c)
	fmt.Printf("a: %v, b: %v, c: %v\nProduct num: %v", a, b, c, productnum)
}

func main() {
	pythagorean()
}
