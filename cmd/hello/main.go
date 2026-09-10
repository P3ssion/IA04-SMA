package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(rand.Intn(100))
	fmt.Println(coucou)
}

func Fills(sl []int) {
	for i := range sl {
		sl[i] = rand.Intn(100)
	}
}
