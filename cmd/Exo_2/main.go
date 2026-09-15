package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//arr1 := [...]int{1, 2, 3, 4}

	fmt.Println("DEBUT DES TESTS")
	slt := make([]int, 5)
	Fills(slt)
	moyenneAvant := Moyenne(slt)
	centreAvant := ValeursCentrales(slt)

	fmt.Println("Slice initial: ", slt)
	fmt.Println("Moyenne avant : ", moyenneAvant)
	fmt.Println("Valeurs centrales avant : ", centreAvant)

	fmt.Println("---Application de Plus1---")
	Plus1(slt)

	moyenneApres := Moyenne(slt)
	centreApres := ValeursCentrales(slt)

	fmt.Println("Slice final : ", slt)
	fmt.Println("Moyenne apres: ", moyenneApres)
	fmt.Println("Centres apres : ", centreApres)

	fmt.Println("Fin des tests")

	maListe := make([]int, 5)

	monSlice := make([]int, 5)
	Fills(monSlice)
	fmt.Println("Résultat du slice que j'ai Fills :", monSlice)

	fmt.Println("Valeurs centrales :", ValeursCentrales(monSlice))

	var monTableau [6]int
	Fills(monTableau[:])
	fmt.Println("Résultat du tableau :", monTableau)
	fmt.Println("Valeurs centrales tableau cases paires :", ValeursCentrales(monTableau[:]))

	RemplirAvecDesZero(maListe)
	fmt.Println(maListe)

	fmt.Println("Q0 Fonction pair")
	pair()

	fmt.Println("Q1 Fonction fill")
	var tab [5]int
	sl := tab[:]
	Fills(sl)
	fmt.Println(sl)
	fmt.Println("Q2 Moyenne")
	fmt.Println(Moyenne(sl))

	// On calcule la moyenne de monSlice
	maMoyenne := Moyenne(monSlice)
	fmt.Println("La moyenne est :", maMoyenne)

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
		//Pas besoin car on peut juste print le slice directement
	}
}

func RemplirAvecDesZero(sl []int) {
	for index := range sl {
		sl[index] = 0
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

// Q5
func ValeursCentrales(sl []int) []int {
	taille := len(sl)
	milieu := taille / 2

	if taille%2 == 0 {
		return []int{sl[milieu-1], sl[milieu]}
	} else {
		return []int{sl[milieu]}
	}
}
