package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	// Get the list of files in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	// Create subdirectories for each day of creation
	for _, file := range files {
		if !file.IsDir() {
			creationTime := file.ModTime()
			creationDay := creationTime.Format("2006-01-02")
			subdir := filepath.Join(dir, creationDay)
			if err := os.MkdirAll(subdir, os.ModePerm); err != nil {
				fmt.Println("Error creating subdirectory:", err)
				os.Exit(1)
			}

			// Move the file to the subdirectory
			oldPath := filepath.Join(dir, file.Name())
			newPath := filepath.Join(subdir, file.Name())
			if err := os.Rename(oldPath, newPath); err != nil {
				fmt.Println("Error moving file:", err)
				os.Exit(1)
			}
		}
	}
}
