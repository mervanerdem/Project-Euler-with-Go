/*
Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import "fmt"

func main() {
	a := 2520 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19 * 20
	b := 0
	result := 0
	for i := 2520; i < a; i += 2520 { //ilk 10 sayı için zaten en küçük bölen sayı 2520 bunun katlarını kontrol etmemiz gerekmekte
		for j := 1; j <= 20; j++ {
			if i%j == 0 {
				b++
				if b == 20 {
					result = i
					break
				}
			} else {
				b = 0
				break
			}
		}
		if b == 20 {
			result = i
			break
		}
	}
	fmt.Println(result)

}
