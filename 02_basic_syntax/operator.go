package main

func operator_string() {
	const a = "a value called a,"
	var b, c string
	c = b
	b = a
	println(a, b, c)
}

func operator_flaot() {
	var a float64 = 2.3
	var pa = &a

	*pa++
	println(a)
}

func main() {
	operator_string()
	println("\ndifferent operator \n")
	operator_flaot()
}

/*
	a value called a, a value called a,

	different operator

	+3.300000e+000
*/
