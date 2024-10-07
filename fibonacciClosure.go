package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	res := 0
	x1, x2 := 1, 1
	return func(x int)int{
		if x == 0 {
			return 0
		} else if x == 1 || x == 2{
			return 1
		} else{
			res = x1+x2
			x1 = x2
			x2 = res
			return res
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}