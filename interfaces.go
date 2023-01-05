package figuras

import "fmt"

type geometria interface {
	area() float32
	perimetro() float32
}

func Medidas(g geometria) {
	fmt.Println("Medidas: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro: ", g.perimetro())
}
