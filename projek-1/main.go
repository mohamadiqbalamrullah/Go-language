package main

import (
	"fmt"
)

func main() {
	// menampilakan nilai i : 21
	var i int = 21
	fmt.Println(i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	// menampilkan tanda %
	fmt.Println("%")

	// menampilkan nilai boolean j : true
	var j bool = true
	fmt.Println(j)

	// menampilkan nilai boolean j : false
	j = false
	fmt.Println(j)

	// menampilkan unicode russia : Я (ya)
	russia := 'Я'
	fmt.Printf("%q \n", russia)

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", i)

	// menampilkan nilai base 8 :25
	var number int = 21
	fmt.Printf("%o \n", number)

	// menampilkan nilai base 16 : f
	beta := 15
	fmt.Printf("%x \n", beta)

	// menampilkan nilai base 16 : F
	delta := 15
	fmt.Printf("%X \n", delta)

	// menampilkan unicode karakter Я : U+042F
	//var Я string = "U+042F"
	fmt.Printf("%U \n", russia)

	// menampilkan float : 123.456000
	var k float64 = 123.456
	fmt.Printf("%F \n", k)

	// menampilkan float scientific : 1.234560E+02
	var scientific float32 = 123.4560
	fmt.Printf("%E \n", scientific)
}
