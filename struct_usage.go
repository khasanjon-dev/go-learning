package main

import "fmt"

type Person struct {
	name      string
	age       int
	job       string
	salary    int
	isMarried bool
}

func main() {
	var pers1 Person
	var pers2 Person
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	pers1.isMarried = true

	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500
	pers2.isMarried = false
	printPerson(pers1)
	printPerson(pers2)
}
func printPerson(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)

}
