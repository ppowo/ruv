package main

//go:generate go run install_tools.go

import (
	"fmt"
	"os"

	"github.com/ppowo/ruv/cmd"
)

func main() {
	// Execute Cobra command tree
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}