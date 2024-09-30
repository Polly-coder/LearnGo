package main
import ("fmt"; 
 //"math"
 )

func main(){
	myArr := [3]int{}
	mySlice := []string{"a"}
	mySlice2 := [][]int{} // слайс слайсов
	mySlice2 = append(mySlice2, []int{3, 3, 4, 5})
	mySlice3 := make([]bool, 2, 3)
	fmt.Println(myArr, mySlice, mySlice2, mySlice3)
	mySlice3 = append(mySlice3, true)
	fmt.Println(mySlice3, len(mySlice3), cap(mySlice3))
	mySlice3 = append(mySlice3, true, false)
	fmt.Println(mySlice3, len(mySlice3), cap(mySlice3))
}