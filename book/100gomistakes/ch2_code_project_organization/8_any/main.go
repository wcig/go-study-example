package main

func main() {
	var a any
	a = 10
	a = "ok"
	_ = a

	var i interface{}
	i = 10
	i = "ok"
	_ = i
}
