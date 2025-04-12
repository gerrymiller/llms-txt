package main

import (
	"fmt"
	"os"

	"github.com/llms-txt/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: llmsgen <path-to-docs>")
		os.Exit(1)
	}

	docPath := os.Args[1]
	manifest, err := parser.GenerateManifest(docPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("llms.txt", []byte(manifest), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write llms.txt: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("llms.txt generated successfully.")
}
