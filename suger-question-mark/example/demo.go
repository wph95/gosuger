package main

import (
	"fmt"
)


func errs() error {
	return nil
}

func main() {
	err := errs()
	b := `11111`
	b = "333333".format()
	fmt.Printf(b)

	// ? err
	{
		fmt.Print(err)
	}
	fmt.Printf("Hello, Golang\n")
}
