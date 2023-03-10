package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %v\n", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			for i := 0; i <= 6; i++ {
				var karakter = [7]rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
				fmt.Printf("character %U '%c' starts at byte position %v\n", karakter[i], karakter[i], i*2)
			}
		} else {
			fmt.Printf("Nilai j = %v\n", j)
		}
	}
}
