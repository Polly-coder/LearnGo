package main

import ("fmt"
	"math")

type T interface{
	Abs()float64
}

type Hey struct{
	x, y float64
}

func (h *Hey) Abs()float64{
	return math.Sqrt(h.x*h.y)
}

func main(){
	var i T
	h := Hey{1, 23}
	i = &h
	fmt.Println(i.Abs())

	i = &Hey{12, 22}
	fmt.Println(i.Abs())
}