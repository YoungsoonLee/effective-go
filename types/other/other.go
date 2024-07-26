package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	val, ok := <-ch
	fmt.Println(val, ok)

	var i any = "hi"
	// n := i.(int) // panic
	if n, ok := i.(int); ok {
		fmt.Println(n)
	} else {
		fmt.Println("not an int")
	}

}
