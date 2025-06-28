package main

import "fmt"

func main() {

	a := []int{1, 2, 3}

	for _, v := range a {
		v *= 2
	}

	fmt.Println(a)
}
