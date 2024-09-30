package main

import (
	"fmt"
)

func main() {
	a1 := make([]int, 0, 10)
	a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	a2 := append(a1, 6)
	a3 := append(a1, 7)
	// a1, a2, a3 ссылаются на одну область памяти
	fmt.Println("a", a1, a2, a3)
	fmt.Println(a1[:7])
	a1[1] = 99
	fmt.Println(a1, a2, a3)

	b1 := make([]int, 0)
	b1 = append(b1, []int{1, 2, 3, 4, 5}...)
	fmt.Println("Capacity b1", cap(b1))
	// если при аппенде не хватает capacity - аллоцируется новая память
	b2 := append(b1, 6)
	fmt.Println("Capacity b2", cap(b2))
	b3 := append(b1, 7)

	fmt.Println("b: ", b1, b2, b3)
	//fmt.Println(b1[:7]) // runtime error: slice bounds out of range [:7] with capacity 5
	b1[1] = 99
	fmt.Println("b changed: ", b1, b2, b3)
}