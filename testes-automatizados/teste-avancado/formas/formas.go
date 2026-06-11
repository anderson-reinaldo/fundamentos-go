package formas

import (
	"math"
)

// Forma representa uma forma geometrica
type Forma interface {
	area() float64
}

// Retangulo representa um retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// area retorna a area do retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Circulo representa um circulo
type Circulo struct {
	Raio float64
}

// area retorna a area do circulo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
