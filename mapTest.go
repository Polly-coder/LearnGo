package main

func main(){
	m := make(map[string]int)
	println("%v", m)
	m["Answer"] = 42
	println("The value", m["Answer"])
	println(m)
	m["Answer"] = 48
	println("The value", m["Answer"])
	delete(m, "Answer")

	v, ok := m["Answer"]
	println("The value", v, "Present?", ok)
}