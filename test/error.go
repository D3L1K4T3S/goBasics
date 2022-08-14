package test

import "fmt"

type appError struct {
	Err    error
	Custom string
	Field  int
}

type AppError interface {
	Error() string
	Unwrap() error
}

func (e *appError) Error() string {
	return e.Custom
}

func (e *appError) Unwrap() error {
	return e.Err
}

func main() {
	err := message()
	if err != nil {
		fmt.Println(err.Unwrap())
	}
}

func message() *appError {
	return &appError{
		Err:    fmt.Errorf("my error"),
		Custom: "value",
		Field:  2,
	}
}
