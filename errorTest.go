package main

import "fmt"

func main(){
	fmt.Println(handle())
	fmt.Println(handleNeg(-3))
}

func handle() error {
	return &myError{text: "error"}
}

func handleNeg(x int) error{
	return NegError(x)
}

// для типа
type NegError int

// для структуры
type myError struct{
	text string
}

func (m *myError) Error() string {
	return m.text
}

func (e NegError) Error() string {
	return fmt.Sprintf("Отрицательное значение: %d", int(e))
}