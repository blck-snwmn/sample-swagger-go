package main

import "fmt"

func main() {
	fmt.Println("hello world")
	returnErr()
}

func returnErr() error {
	return nil
}
