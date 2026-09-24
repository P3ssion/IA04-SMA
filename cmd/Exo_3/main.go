package main

import (
	"fmt"
	"math"
	"sync"
)

type Point2D struct {
	x float64
	y float64
}

func NewPoint2D(x float64, y float64) *Point2D {
	nouveauPoint := Point2D{x: x, y: y}
	return &nouveauPoint
}

func (p Point2D) X() float64 {
	return p.x
}

func (p *Point2D) SetX(nouvelleValeur float64) {
	p.x = nouvelleValeur
}

func (p Point2D) Y() float64 {
	return p.y
}

func (p *Point2D) SetY(nouvelleValeur float64) {
	p.y = nouvelleValeur
}

func (p Point2D) Clone() Point2D {
	copie := Point2D{x: p.x, y: p.y}
	return copie
}

func (p Point2D) CloneLeger() *Point2D {
	return NewPoint2D(p.x, p.y)
}

func (p Point2D) Module() float64 {
	return math.Sqrt(math.Pow(p.X(), 2) + math.Pow(p.Y(), 2))
}

// Q2
type Rectangle struct {
	sommetHaut *Point2D
	sommetBas  *Point2D
}

func NewRectanglePointeurs(p1 *Point2D, p2 *Point2D) *Rectangle {
	nouveauRectangle := Rectangle{sommetHaut: p1, sommetBas: p2}
	return &nouveauRectangle
}

func NewRectangleAdr(p1 Point2D, p2 Point2D) *Rectangle {
	nouveauRectangle := Rectangle{sommetHaut: &p1, sommetBas: &p2}
	return &nouveauRectangle
}

func main() {
	monPoint := NewPoint2D(2.5, 4.0)
	fmt.Println("Mon nv point est : ", monPoint)

	monPoint.SetX(9.9)
	fmt.Println("Nv X lu par le getter est :", monPoint.X())
	fmt.Println(monPoint)

	p1 := NewPoint2D(3.0, 4.0)
	fmt.Println("P1 initial : ", p1)

	distance := p1.Module()
	fmt.Println("Distance à l'origine de p1 est :", distance)
	fmt.Println()

	p2 := p1.Clone()
	fmt.Println("-----fonction clone-----")
	fmt.Println("Clone de p1 qui est p2 = ", p2)
	fmt.Println()

	p2.SetX(99.0)
	fmt.Println("--Après modification du clone--")
	fmt.Println("Le x de l'original P1 vaut toujours : ", p1.X())
	fmt.Println("Le x du clone p2 vaut maintenant: ", p2.X())
	fmt.Println("p2=", p2)

	fmt.Println()
	fmt.Println("-----fonction clone leger-----")
	p3 := p2.Clone()
	fmt.Println("Clonne de p2 qui est p3 = ", p3)
	p4 := p1.CloneLeger()
	fmt.Println("Clone leger de P1 qui est p4 = ", p4)
	p5 := p3.CloneLeger()
	fmt.Println("Clone leger de p3 qui est p5 =", p5)
	fmt.Println()

	fmt.Println("----Rectangle----")
	fmt.Println("----Rectangle avec pointeurs ----")
	recPointeurs := NewRectanglePointeurs(p1, p1)
	fmt.Println(recPointeurs)
	fmt.Println("----Rectangle avec pointeurs ----")
	recAdr := NewRectangleAdr(p2, p3)
	fmt.Println(recAdr)

	fmt.Println("----Rectangle avec copie physique PAS BON ----")
	pointHaut := Point2D{x: 0, y: 10}
	pointBas := Point2D{x: 10, y: 0}

	// 2. On appelle ton usine
	monRect := NewRectangleAdr(pointHaut, pointBas)

	// 3. On affiche le contenu du rectangle
	// (Le %+v est une astuce Go pour afficher tout le contenu d'une structure)
	fmt.Printf("Mon rectangle contient : %+v\n", monRect)
	fmt.Printf("Le point haut est : %+v\n", monRect.sommetHaut)

	/*fmt.Println("----FONCTION COMPTE----")
	go compte(11)
	compte(10)
	go compte(12)

	fmt.Println("----Fonction compte Msg----")
	go compteMsg(10, "hello")
	go compteMsg(10, "Miam")
	go compteMsg(10, "Lolo")

	fmt.Println("----Fonction compte Msg finale----")
	//go compteMsgFromTo(9, 15, "coco")
	go compteMsgFromTo(2, 7, "freidn")
	//fmt.Scanln()

	for i := range 10 {
		msg := fmt.Sprintf("#%d:", i)
		go compteMsgFromTo(i*10, (i+1)*10, msg)
	}

	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 10000; i++ {
		go f()
	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)

	*/
}

func compte(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func compteMsg(n int, msg string) {
	for i := 0; i < n; i++ {
		fmt.Println(msg, i)
	}
}

func compteMsgFromTo(start int, end int, msg string) {
	for i := start; i < end; i++ {
		fmt.Println(msg, i)
	}
}

var n = 0
var l sync.Mutex

func f() {
	l.Lock()
	defer l.Unlock()
	n++
}
