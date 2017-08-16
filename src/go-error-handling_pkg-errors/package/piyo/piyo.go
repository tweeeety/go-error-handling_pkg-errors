package piyo

import (
	"go-error-handling_pkg-errors/package/hoge"

	"github.com/pkg/errors"
)

func DoExciting(b bool) error {
	err := hoge.DoExciting(b)

	if err != nil {
		switch err.(type) {

		case *hoge.HogeSomethingError:
			return errors.Wrap(err, "error1 at piyo")

		case *hoge.HogeAnythingError:
			return errors.Wrap(err, "error2 at piyo")

		}
	}

	return nil
}
