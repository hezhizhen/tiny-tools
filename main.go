package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("HOME"))
	fmt.Println(os.Getenv("USER"))
}
