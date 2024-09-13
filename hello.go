// Package main provides a simple "Hello, World!" program.
// It demonstrates basic Go syntax and how to create an executable program.
package main

import "fmt"

// Hello returns a friendly greeting message.
//
// This function is used to demonstrate a simple Go function
// that returns a string value.
//
// Returns:
//
//	string: A "Hello, World!" greeting message.
func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}
