package main

import (
	"fmt"
	"os"
)

func main() {
	debug, ok := os.LookupEnv("DEBUG")
	if !ok {
		fmt.Println("DEBUG environment variable is not set")
		os.Exit(1)
	}
	fmt.Println("DEBUG:", debug)
}
