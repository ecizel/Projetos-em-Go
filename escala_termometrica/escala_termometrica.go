package main

import "fmt"

func main() {

	ebulicao_agua_kelvin := 373.0

	// C = K - 273
	ebulicao_agua_celcius := ebulicao_agua_kelvin - 273.0

	fmt.Printf("O ponto de ebulicação da água é %g K ou %g ºC", ebulicao_agua_kelvin, ebulicao_agua_celcius)

}