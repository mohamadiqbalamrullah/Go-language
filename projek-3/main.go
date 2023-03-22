package main

import (
	"fmt"
)

func main() {
	sentc := "selamat malam"

	for _, x := range sentc {
		fmt.Printf("%c \n", x)
	}

	count := make(map[string]int)

	for _, x := range sentc {
		count[string(x)] += 1
	}

	fmt.Println(count)
}
