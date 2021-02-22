package formatter

import "fmt"

// Println will print input string
func Println(input string) {
	println(input)
}

// myPrintln will not export
func println(input string) {
	fmt.Println(input)
}
