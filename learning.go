package main

import "fmt"

func main() {
	time := 6
	if time < 6 && time < 18 {
		fmt.Println("Good day.")
	} else if time == 6 {
		fmt.Println("Good morning")
	} else {
		fmt.Println("Good evening.")
	}

}
