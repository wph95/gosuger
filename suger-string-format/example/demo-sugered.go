package main

import (
	"fmt"
)


func errs() error {
	return nil
}

func main() {
	apple := "Golang China Foundation"
	pen := "PingCAP"
	slogan := fmt.Sprintf("%s x %s",apple, pen)
	println(slogan)
	println(fmt.Sprintf("%s x %s",apple, pen))
}
