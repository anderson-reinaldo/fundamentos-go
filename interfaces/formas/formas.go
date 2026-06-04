package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreveArea(nomeForma string, f forma) {
	fmt.Printf("A area da forma %s é: %f\n", nomeForma, f.area())
}

func main() {
	r := retangulo{10, 15}
	c := circulo{10}
	escreveArea("retangulo", r)
	escreveArea("circulo", c)
}
