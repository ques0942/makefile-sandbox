package main

import (
	"fmt"
	"os"
)

func main() {
	v := os.Getenv("VAR")
	if v != "" {
		fmt.Printf("VAR is %s\n", v)
	} else {
		fmt.Println("VAR is empty")
	}
}
