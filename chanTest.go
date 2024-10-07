package main

import "fmt"

func main(){
	ch := make(chan int)
	go f(10, ch)
	
	for{
		val, ok := <- ch
		if !ok{
			break
		}
		fmt.Println(val)
	}
	fmt.Println("The Ennd")
}

func f(i int, ch chan int){
	defer close(ch)
	for j := 0; j < i; j++{
		ch <- j
	}
	
}