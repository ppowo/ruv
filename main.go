package main

//go:generate go run install_tools.go

import "fmt"

func main() {
	fmt.Println("Hello from ruv!")
	fmt.Println("This is a Go CLI tool built with mage.")
}