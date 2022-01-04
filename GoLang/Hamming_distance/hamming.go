package main

import (
	"fmt"
)

func main() {
	var s1 string
	var s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)

	var d int
	d = 0
	for i := 0; s1[i] <= s2[i]; i++ {
		if len(s1) != len(s2) {
			fmt.Println("Two strings are not of equal length")
		} else if s1[i] == 'a' || s1[i] == 'g' || s1[i] == 'c' || s1[i] == 't' {
			i++

			for i := range s1 {

				if s1[i] != s2[i] {
					d++
				}

			}
		} else {
			fmt.Println("Invalid hamming code")

		}

	}

	// if s1[0] == 'a' || s1[0] == '' {
	// 	fmt.Println("a")
	// } else {
	// 	fmt.Println('b')
	// }
	fmt.Println(d)
}
