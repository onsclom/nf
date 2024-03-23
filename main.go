package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: nf <path/to/file>")
		os.Exit(1)
	}

	filePath := args[0]

	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Printf("Error creating directories for %s: %v\n", filePath, err)
		os.Exit(1)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	file.Close()

	fmt.Printf("File created successfully: %s\n", filePath)
	// TODO: give nice tree view!
}
