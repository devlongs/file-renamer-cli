package main

import (
	"fmt"
	"os"
)

func main() {
    // Get the current file name and the new file name from command-line arguments
    oldFile := os.Args[1]
    newFile := os.Args[2]

    // Use the os.Rename function to rename the file
    err := os.Rename(oldFile, newFile)
    if err != nil {
        fmt.Println("Error renaming file:", err)
    } else {
        fmt.Println("File renamed successfully!")
    }
}
