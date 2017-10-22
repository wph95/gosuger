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
	slogan := `${apple} x ${pen}`
	print(slogan)
	println("${apple} x ${pen}".format(apple, pen))
}
