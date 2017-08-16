package fuga

import (
	"errors"
	"go-error-handling_pkg-errors/package/hoge"
)

func DoExciting(b bool) error {
	err := hoge.DoExciting(b)

	if err != nil {
		switch err.(type) {

		case *hoge.HogeSomethingError:
			return errors.New("handling hoge.HogeSomethingError at fuga")

		case *hoge.HogeAnythingError:
			return errors.New("handling hoge.HogeAnythingError at fuga")

		}
	}

	return nil
}
