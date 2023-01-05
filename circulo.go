package figuras

import "math"

type Circulo struct {
	Radio float32
}

func (cir *Circulo) area() float32 {
	return math.Pi * (cir.Radio * cir.Radio)
}

func (cir *Circulo) perimetro() float32 {
	return 2 * math.Pi * cir.Radio
}
