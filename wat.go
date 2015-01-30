package main

import (
	"fmt"
	"os"
)

func main() {
	doStuff()
}

func doStuff() {
	fmt.Printf("=====> %s <======\n", os.Getenv("TRAVIS_GO_VERSION"))
	fmt.Println("£¢∞§∞¢£¢∞∞¢∞§∞¢∞§")
}
