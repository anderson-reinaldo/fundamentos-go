package formas_test

import (
	"math"
	"testing"

	"teste-avancado/formas"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		retangulo := formas.Retangulo{10, 15}
		areaEsperada := float64(150)
		areaRecebida := retangulo.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("esperava %f, recebeu %f", areaEsperada, areaRecebida)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circulo := formas.Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()
		if areaEsperada != areaRecebida {
			t.Errorf("esperava %f, recebeu %f", areaEsperada, areaRecebida)
		}
	})
}
