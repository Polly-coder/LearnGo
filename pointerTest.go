package main

func main (){
	first := []int{1, 2, 3, 4}
	second := make([]*int, len(first))
	println(&first[0])
	for i, v := range first{
		// адрес в памяти i всегда один и тот же 
		// для v каждый раз создаётся новая ячейка памяти
		println(&i, &v)
		second[i] = &v
	}
	println(*second[0], *second[1])
}