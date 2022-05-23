/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
package main

import "fmt"

func main() {
	x := 2000000 // number
	var s []int
	s = append(s, 2) // first number of slice
	res := 0         //result
	a := 0
	for i := 2; i <= x; i++ { // Number loop
		boll1 := false
		boll2 := false
		boll3 := false
		var t int

		for k := range s {
			if i%s[k] == 0 {
				boll3 = true
				break
			}
		}
		if boll3 {
			continue
		}

		for j := 2; j <= i; j++ { // find prime number
			if i%j == 0 {
				boll2 = true
				t = j
				break
			}
		}
		if boll2 { // if not

			for y := range s { //is it in our list
				if s[y] == t {
					boll1 = false
					break
				}
				boll1 = true
			}
			if boll1 { // if out of list append it
				res += s[a]
				s = append(s, t)
				fmt.Println(t)
				a++
			}
		}
	}
	res += s[a]
	/*
		for i := range s {
			res += s[i]
			fmt.Println(s[i])
		}
	*/
	fmt.Println("Result:", res)
}
