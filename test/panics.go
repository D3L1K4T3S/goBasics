package test

import (
	"errors"
	"fmt"
	"log"
)

//Кастомная ошибка

type AppErr struct {
	Message string
	Err     error
}

func (ae *AppErr) Error() string {
	return ae.Message
}

type name struct {
	A, B int
}

func (n *name) method() {
	fmt.Println("all okey")
}

// Обработка паники
func devide(a, b int) {
	defer func() {
		var appErr *AppErr
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("panic", err)
				} else {
					panic("!_!")
				}
			default:
				panic("some panic")
			}
			log.Println("panic happened: ", err)
		}
	}()
	fmt.Println(div(a, b))
}

func div(x, y int) int {
	if y == 0 {
		panic(&AppErr{
			Message: "this is divide by zero custom error",
			Err:     nil,
		})
	}
	return x / y
}

func main() {
	n := &name{1, 2}
	n = nil
	n.method()
	devide(1, 0)
	fmt.Println("Hello world")
}
