package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ohhai")
	fmt.Printf("%s\n", os.Getenv("TRAVIS_GO_VERSION"))
}
