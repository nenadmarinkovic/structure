package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var projectPath string
	if len(os.Args) > 1 {
		projectPath = os.Args[1]
	} else {
		var err error
		projectPath, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
	}

	err := validatePath(projectPath)
	if err != nil {
		fmt.Println("Invalid path:", err)
		return
	}

	tree, err := generateTree(projectPath, "")
	if err != nil {
		fmt.Println("Error generating tree structure:", err)
		return
	}

	readmePath := filepath.Join(projectPath, "README.md")
	file, err := os.OpenFile(readmePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening/creating README.md:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("\n\n### Tree structure of the project:\n\n")
	writer.WriteString(tree)
	writer.Flush()

	fmt.Println("Tree structure created successfully!")
}
