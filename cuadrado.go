package figuras

type Cuadrado struct {
	Lado float32
}

func (cua *Cuadrado) area() float32 {
	return cua.Lado * cua.Lado
}

func (cua *Cuadrado) perimetro() float32 {
	return 4 * cua.Lado
}
