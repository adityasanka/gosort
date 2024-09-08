// Package main provides a simple "Hello, World" example.
package main

import "fmt"

// Hello returns a standard greeting message.
//
// Returns:
//   - string: A "Hello, World" greeting message.
func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}
