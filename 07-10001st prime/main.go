/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/

package main

import "fmt"

func main() {
	slc := []int{2, 3, 5, 7, 11, 13}
	i := 14
	ix := 14
	x := 0
	for {
		i = ix
		for _, value := range slc {
			if i%value == 0 {
				x = 1
				break
			}
		}
		if x != 1 {
			fmt.Println(i)
			slc = append(slc, i)
		}
		ix++
		x = 0
		if len(slc) == 10001 {
			fmt.Println(slc[10000])
			break
		}
	}
}
