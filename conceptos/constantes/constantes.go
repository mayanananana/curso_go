package main

import "fmt"

const (
	apple, banana, orange = "manzana", 1, "naranja"
)

const (
	place = "Madrid"
	year  = 2024
)

func main() {

	// tengo que declarar y asginar inmediatamente el valor a una constante
	const pi float64 = 3.14159
	const gravedad float64 = 9.81
	const velocidadLuz int = 299792458
	const esProgramador bool = true
	const nombre string = "Go"

	fmt.Println("Valor de pi:", pi)

	// No se puede cambiar el valor de una constante después de su declaración
	// pi = 3.14 // Esto generará un error de compilación
	const os, dos = "Windows", "Linux"

	fmt.Println("Sistema operativo 1:", os)
	fmt.Println("Sistema operativo 2:", dos)
}
