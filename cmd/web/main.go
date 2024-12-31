package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	listOfWords := strings.Split(strings.TrimSpace(s), " ")
	fmt.Println(listOfWords)
	return len(listOfWords[len(listOfWords)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}
