package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	showTheTaskMenu()
}

// display welcome menu
func showTheTaskMenu() {
	// create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=======================================")
	fmt.Println("Edspert.id Bootcamp Backend Dev with GO")
	fmt.Println("Tugas Pertemuan Ke-1")
	fmt.Println("=======================================")
	fmt.Println("")
	fmt.Println("1. Program Palindrome")
	fmt.Println("2. Program Menentukan Vokal atau Konsonan")
	fmt.Println("3. Exit Program")
	fmt.Println("")
	fmt.Println("Pilih program 1/2/3:")

	choice, error := reader.ReadString('\n')
	choice = removeNewlineChar(choice)

	if error == nil {
		switch choice {
		case "1":
			palindromeMenu()
		case "2":
			fmt.Println("Anda telah memilih no. 2")
		case "3":
			fmt.Println("Terima Kasih, sampai jumpa kembali di program ini :)")
			os.Exit(0)
		default:
			fmt.Println("Anda tidak memilih pilihan yang benar, silahkan mulai kembali program ini.")
			os.Exit(0)
		}
	} else {
		fmt.Println("Ada kesalahan dalam input.")
		os.Exit(0)
	}

}

func palindromeMenu() {
	// create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n\n=======================================")
	fmt.Println("Edspert.id Bootcamp Backend Dev with GO")
	fmt.Println("1. Program Palindrome")
	fmt.Println("=======================================")

	fmt.Println("Silahkan masukkan kata:")
	answer, _ := reader.ReadString('\n')
	answer = removeNewlineChar(answer)
	isPal := isPalindrome(answer)

	fmt.Printf("\nInput: %s\n", answer)
	fmt.Printf("Is Palindrome: %t\n\n", isPal)

	showTheTaskMenu()
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		charFromFront := str[i]

		j := len(str) - 1 - i
		charFromBack := str[j]

		if charFromFront != charFromBack {
			return false
		}
	}
	return true
}

func removeNewlineChar(input string) string {
	newInput := strings.Replace(input, "\r\n", "", -1)
	newInput = strings.Replace(newInput, "\n", "", -1)
	return newInput
}
