package main

import "fmt"

func main() {
	// menampilkan nilai i : 21
	i := 21
	fmt.Printf("%v \n", i)

	//menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)

	// menampilkan tanda %
	fmt.Printf("%%\n")

	// menampilkan nilai boolean j : true
	var j bool = true
	fmt.Printf("%t \n", j)

	// menampilkan nilai boolean j : true
	fmt.Printf("%t \n", j)

	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%U\n", 'Я')

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d\n", 21)

	// menampilkan nilai base 8 :25
	fmt.Printf("%o\n", 21)

	// menampilkan nilai base 16 : f
	f := 15
	fmt.Printf("%x\n", f)

	// menampilkan nilai base 16 : F
	F := 15
	fmt.Printf("%X\n", F)

	// menampilkan unicode karakter Я : U+042F
	fmt.Printf("%U\n", 'Я')

	// var k float64 = 123.456;
	var k float64 = 123.456
	fmt.Printf("%v\n", k)

	// menampilkan float : 123.456000
	fmt.Printf("%.6f\n", k)

	// menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e\n", k)
}
