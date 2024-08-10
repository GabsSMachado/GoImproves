package main

import "fmt"

func main() {
	tempK := 373.0            //Temperatura em Kelvin
	tempC := int(tempK - 273) //Temperatura em graus Celcius

	//Imprime as temperaturas convertidas.
	fmt.Printf("Hoje a temperatura de ebulição em Kelvin é de : %g ºK . Já a temperatura de ebulição em graus celsius é de : %d ºC", tempK, tempC)
}
