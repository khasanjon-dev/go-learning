package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	for k, v := range a {
		fmt.Printf("%v: %v, ", k, v)
	}
	//b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	//fmt.Printf("a\t%v\n", a)
	//fmt.Printf("b\t%v\n\n", b)
	//var a = make(map[string]string)
	//a["brand"] = "Ford"
	//a["model"] = "Mustang"
	//a["year"] = "1964"
	//
	//b := make(map[string]int)
	//b["Oslo"] = 1
	//b["Bergen"] = 2
	//b["Trondheim"] = 3
	//b["Stavanger"] = 4
	//
	val1, ok1 := a["brand"]
	fmt.Println(val1, ok1)
	// delete
	//delete(a, "year")
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(a["brand"])
}
