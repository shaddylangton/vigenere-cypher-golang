package main

import (
	"fmt"
	"strings"
)

func encodeString(cipher, key rune) rune {
	const asciiA rune = 65
	const numLetters = 26

	plainTextIndex := cipher + key
	asciiLetter := (plainTextIndex+numLetters)%numLetters + asciiA
	return asciiLetter
}

func encode(message, kw string) string {
	var plainText strings.Builder
	kwChars := []rune(kw)

	for i, cipherChar := range message {
		key := i % len(kwChars)
		plainText.WriteRune(encodeString(cipherChar, kwChars[key]))
	}

	return plainText.String()
}
func decipherString(cipher, key rune) rune {
	const asciiA rune = 65
	const numLetters = 26

	plainTextIndex := cipher - key
	asciiLetter := (plainTextIndex+numLetters)%numLetters + asciiA
	return asciiLetter
}

func decipher(message, kw string) string {
	var plainText strings.Builder
	kwChars := []rune(kw)

	for i, cipherChar := range message {
		key := i % len(kwChars)
		plainText.WriteRune(decipherString(cipherChar, kwChars[key]))
	}

	return plainText.String()
}

func main() {
	fmt.Println("Enter Your string: ")

	var first string

	fmt.Scanln(&first)
	fmt.Println("Enter your KEY: ")
	var second string
	fmt.Scanln(&second)
	cipherText := first
	keyword := second
	fmt.Print("Do you want to  1. Encrypt or 2. Decrypt")
	var option int
	fmt.Scanln(&option)
	if option == 1 {
		fmt.Println(encode(cipherText, keyword))
	} else if option == 2 {
		fmt.Println(decipher(cipherText, keyword))
	} else {
		fmt.Println("please choose the right option")
	}

}
