package wightconv

import "fmt"

type Kg float64
type G 	float64

func (kg Kg) String() string {
	return fmt.Sprintf("%gKg", kg)
}

func (g G) String() string {
	return fmt.Sprintf("%gg", g)
}

func KToG(kg Kg) G {
	return G(kg * 1000)
}

func GToK(g G) Kg {
	return Kg(g / 1000)
}