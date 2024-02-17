package main

import "fmt"

//func main() {
//	fruits := [3]string{"apple", "orange", "banana"}
//	for idx, val := range fruits {
//		fmt.Printf("%v %v\n", idx, val)
//	}
//}

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
}
