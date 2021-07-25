package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}
	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	//Palindromos ama, amor a roma
	isPalindromo("casa")
	isPalindromo("ama")
	isPalindromo("amor a roma")
	isPalindromo("Ama")
}
