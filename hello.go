package main

import "fmt"

func main() {

	fmt.Printf("Hello World\n")

	a := 0
	b := "hello"
	fmt.Printf("%d", a)

	fmt.Print(a, a, a)
	fmt.Print(b)

	for i := 0; i < 5; i++ {
		fmt.Print(i)
		if i == 3 {
			break
		}
	}

}
