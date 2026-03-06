package main

import "fmt"

func main() {
	var edad int = 30
	var nombre string = "Juan"
	var altura float64 = 1.75
	var esEstudiante bool = true
	fmt.Println("Edad:", edad)
	fmt.Println("Nombre:", nombre)
	fmt.Println("Altura:", altura)
	fmt.Println("Es estudiante:", esEstudiante)

	var fruta string
	fruta = "manzana"
	fmt.Println("Fruta:", fruta)

	// Declaración e inicialización en una sola línea
	var numero1, numero2 int = 10, 20
	fmt.Println("Número 1:", numero1)
	fmt.Println("Número 2:", numero2)

	// Declaración e inicialización con inferencia de tipo
	var (
		ciudad    string = "Madrid"
		pais      string = "España"
		poblacion int    = 3223000
	)
	fmt.Println("Ciudad:", ciudad)
	fmt.Println("País:", pais)
	fmt.Println("Población:", poblacion)

	// Declaración e inicialización con inferencia de tipo en una sola línea
	maquillaje1, maquillaje2 := "labial", "rímel"
	fmt.Println("Maquillaje 1:", maquillaje1)
	fmt.Println("Maquillaje 2:", maquillaje2)

}
