package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//arr1 := [...]int{1, 2, 3, 4}
	fmt.Println("Q0 Fonction pair")
	pair()

	fmt.Println("Q1 Fonction fill")
	var tab [5]int
	sl := tab[:]
	Fills(sl)
	fmt.Println(sl)
	fmt.Println("Q2 Moyenne")
	fmt.Println(Moyenne(sl))

	fmt.Println("Q4 Plus1")
	fmt.Println(sl)
	Plus1(sl)

}

// Question intro
func pair() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

// Q1
func Fills(sl []int) {
	for i := range sl {
		sl[i] = rand.Intn(100)
		//fmt.Println(sl[i])
	}
}

// Q2
func Moyenne(nombres []int) float64 {
	var total int = 0
	for _, valeur := range nombres {
		total = total + valeur
	}
	return float64(total) / float64(len(nombres))
}

//Q3

// Q4
func Plus1(sl []int) {
	for i := range sl {
		sl[i] = sl[i] + 1
	}
}
