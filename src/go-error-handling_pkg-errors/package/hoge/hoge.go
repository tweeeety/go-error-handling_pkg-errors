package hoge

import "fmt"

type HogeSomethingError struct{}

func (f *HogeSomethingError) Error() string {
	return fmt.Sprintf("this is HogeSomethingError")
}

type HogeAnythingError struct{}

func (f *HogeAnythingError) Error() string {
	return fmt.Sprintf("this is HogeAnythingError")
}

func DoSomething() error {
	return &HogeSomethingError{}
}

func DoAnything() error {
	return &HogeAnythingError{}
}

func DoExciting(b bool) error {
	if b {
		return DoSomething()
	} else {
		return DoAnything()
	}
	return nil
}
