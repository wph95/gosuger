package main

import (
	"fmt"
)


func errs() error {
	return nil
}

// ? 2345
func main() {
	err  := errs()
	// ? hello
	{
		fmt.Print(err)
	}
	// ? asdf
	fmt.Printf("Hello, Golang\n")
}
