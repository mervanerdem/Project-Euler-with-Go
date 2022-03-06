/*
Largest palindrome product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import "fmt"

func main() {
	res, result := 0, 0

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			res = i * j
			a := res / 100000
			b := res/10000 - a*10
			c := res/1000 - a*100 - b*10
			d := res/100 - a*1000 - b*100 - c*10
			e := res/10 - a*10000 - b*1000 - c*100 - d*10
			f := res - a*100000 - b*10000 - c*1000 - d*100 - e*10
			first := a*100 + b*10 + c
			second := f*100 + e*10 + d
			if first == second && res > result {
				result = res
				fmt.Println(result)
			}
		}
	}
}
