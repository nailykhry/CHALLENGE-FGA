package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calc(statement string) {
	charmap := make(map[string]int)
	sliceOutput := strings.Split(statement, "")
	for _, value := range sliceOutput {
		charmap[value]++
		fmt.Println(value)
	}

	fmt.Print(charmap)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimRight(sentence, "\n")
	calc(sentence)
}
