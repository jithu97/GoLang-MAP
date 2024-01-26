package main

import "fmt"

func main() {

	shoppingCart := map[string]int{
		"keyboard": 100,
		"mouse":    200,
		"monitor":  300,
	}

	shoppingCart2 := make(map[string]int)

	_,ok := shoppingCart["token"];

	delete(shoppingCart, "monitor");

	fmt.Println(shoppingCart)
	fmt.Println(shoppingCart2)
	fmt.Println(ok)
	fmt.Println(len(shoppingCart))
}