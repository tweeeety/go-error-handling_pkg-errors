package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"go-error-handling_pkg-errors/package/hoge"
	"go-error-handling_pkg-errors/package/piyo"
)

func main() {
	if err := piyo.DoExciting(false); err != nil {
		switch errors.Cause(err).(type) {
		case *hoge.HogeSomethingError:
			fmt.Fprintln(os.Stderr, "case1 at main:", err)
		case *hoge.HogeAnythingError:
			fmt.Fprintln(os.Stderr, "case1 at main:", err)
		}
	}

	fmt.Println("main success")
}
