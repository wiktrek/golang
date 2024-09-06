package main

import "fmt"
func main() {
	loop(5, "wiktrek")
}
func loop(times int, text string) {
	for i := 0; i < times; i++ {
	fmt.Println("", text)
	}
}
