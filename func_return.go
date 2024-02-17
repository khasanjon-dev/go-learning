package main

func myFunction(x int, y int) int {
	return x + y
}
func myFunction1(x int, y int) (result int) {
	result = x + y
	return result

}

func multipleReturn(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
func main() {
	//println(myFunction(1, 2))
	//total := myFunction1(2, 3)
	//println(total)
	//result, result1 := multipleReturn(20, "Hello")
	_, result := multipleReturn(20, "Hello")
	println(result)
}
