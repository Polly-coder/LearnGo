package main

import (
	"strings"
	"fmt"
)

func wordCount(s string) map[string]int {
	sl := strings.Fields(s)
	res := make(map[string]int)
	for _, v := range sl{
		res[v]++
	}
	return res
}

func main(){
	println(wordCount("hello word"))
	fmt.Println(wordCount("hello world world Hello"))
}