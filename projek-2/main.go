package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			// number := "024681012"
			alpa := "C A ш A P B O"
			alparray := []rune(alpa)
			// array := []rune(number)
			for i := 0; i < 13; i++ {
				if i%2 == 0 {
					str1 := string(alparray[i])
					str2 := int(alparray[i])
					fmt.Printf("Character %U '%s' starts at byte position %d \n", str2, str1, i)

				}
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}

}
