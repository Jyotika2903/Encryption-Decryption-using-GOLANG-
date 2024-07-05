package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashletterfn(key int, letter string) string {
	runes := []rune(letter)
	lastletterkey := string(runes[len(letter)-key : len(letter)])
	leftoverletter := string(runes[0 : len(letter)-key])
	return fmt.Sprintf("%s%s", lastletterkey, leftoverletter)
}

func encrypt(key int, plaintext string) string {
	hashletter := hashletterfn(key, originalLetter)
	var hashedstring string = ""

	findone := func(r rune) rune {
		pos := strings.Index(originalLetter, string([]rune{r}))
		if pos != -1 {
			letterposition := (pos + len(originalLetter)) % len(originalLetter)
			hashedstring = hashedstring + string(hashletter[letterposition])
			return r
		}
		return r
	}

	strings.Map(findone, plaintext)
	return hashedstring
}

func decrypt(key int, encryptedtext string) string {
	hashletter := hashletterfn(key, originalLetter)
	var hashedstring = ""

	findone := func(r rune) rune {
		pos := strings.Index(hashletter, string([]rune{r}))
		if pos != -1 {
			letterposition := (pos + len(originalLetter)) % len(originalLetter)
			hashedstring = hashedstring + string(originalLetter[letterposition])
			return r
		}
		return r
	}

	strings.Map(findone, encryptedtext)
	return hashedstring
}

func main() {
	var plaintext string

	fmt.Print("Enter the string: ")
	fmt.Scanln(&plaintext)

	fmt.Println("The entered string is:", plaintext)

	encrypted := encrypt(5, plaintext)
	fmt.Println("Encrypted text is:", encrypted)

	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted text is:", decrypted) // 5 is the key here; it can be modified as you wish
}
