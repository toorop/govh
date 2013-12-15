package main

import (
	"fmt"
	"os"
)

func dieError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func debug(msg string) {
	fmt.Println(msg)
}
