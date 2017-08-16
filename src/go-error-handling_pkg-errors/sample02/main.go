package main

import (
	"fmt"
	"os"

	"go-error-handling_pkg-errors/package/piyo"
)

func main() {
	if err := piyo.DoExciting(false); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("main success")
}
