package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindFone(expression string) string {

	var fone string

	for _, c := range expression {
		char := fmt.Sprintf("%c", c)
		switch strings.ToUpper(char) {
		case "1", "0", "-", " ":
			fone += string(c)
		case "A", "B", "C":
			fone += "2"
		case "D", "E", "F":
			fone += "3"
		case "G", "H", "I":
			fone += "4"
		case "J", "K", "L":
			fone += "5"
		case "M", "N", "O":
			fone += "6"
		case "P", "Q", "R", "S":
			fone += "7"
		case "T", "U", "V":
			fone += "8"
		case "W", "X", "Y", "Z":
			fone += "9"
		default:
			return "invalid character"
		}
	}
	return fone
}

func main() {
	fmt.Print("Enter expression: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	input = strings.TrimSuffix(input, "\n")
	fone := FindFone(input)
	fmt.Println(fone)
}
