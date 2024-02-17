package main

import "fmt"

func familyName(fname string) {
	fmt.Println("Hello", fname)

}
func familyAges(firstName string, age int) {
	fmt.Println("Hello", age, "year old", firstName)

}

func main() {
	familyName("Khasan")
	familyAges("Khasan", 21)
}
