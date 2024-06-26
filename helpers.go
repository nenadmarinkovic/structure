package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func isHidden(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

func validatePath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}

func generateTree(path string, prefix string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	list := ""

	for _, file := range files {
		if isHidden(file.Name()) || file.Name() == "node_modules" || file.Name() == "README.md" {
			continue
		}

		list += fmt.Sprintf("%s- %s\n", prefix, file.Name())
		if file.IsDir() {
			subList, err := generateTree(filepath.Join(path, file.Name()), prefix+"  ")
			if err != nil {
				return "", err
			}
			list += subList
		}
	}

	return list, nil
}
