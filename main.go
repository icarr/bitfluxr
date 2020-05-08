package main

import (
	"fmt"
)

var (
	author  string
	version string
)

func main() {
	author = "Ian Carr"
	version = "0.0.1"
	fmt.Printf("BitFluxR v.%s by %s\n", version, author)
}
