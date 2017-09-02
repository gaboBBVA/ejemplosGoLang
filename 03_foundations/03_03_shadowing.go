package main

import "fmt"

func main() {
	var x = 1
	{
		var x = 2
		fmt.Println("inner block x: ", x)
	}
	fmt.Println("outter block x: ", x)
}
