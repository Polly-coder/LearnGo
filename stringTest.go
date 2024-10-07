package main

import "fmt"

func main(){
	i := "Hello"
	ch := make(chan struct{})
	go func(){
		defer close(ch)
		fmt.Println(i)
	}()
	<- ch
	fmt.Println(i)
}