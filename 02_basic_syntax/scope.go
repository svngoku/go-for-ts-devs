package main

func main() {
	var x int = 5
	println(x)
	{
		var x string = "cinq"
		println(x)
	}
	println(x)
}

/*
	println 1: 5
	println 2: cinq
	println 3: 5
*/
