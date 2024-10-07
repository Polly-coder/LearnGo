package main

import ("fmt")

type Hello struct {
	A int
}

// у структур есть методы
// скобки перед названием метода - ресивер
func (h *Hello) PrintA(){
	fmt.Println(h.A)
}

func main(){
	s := Hello{11}
	s.PrintA()
}