package main

import "fmt"

func main() {
	tempF := 65.0                  //Temperatura em Fahrenheit
	tempC := int(tempF-32) * 5 / 9 //Temperatura em graus Celcius

	//Imprime as temperaturas convertidas.
	fmt.Printf("Hoje a temperatura em Florianópolis é de : %g ºF . Já a temperatura em graus celsius é de : %d ºC", tempF, tempC)

	x := tempC
	if x < 18 {
		fmt.Println("\n Está frio, pegue seu casaquinho!")
	} else {
		fmt.Println("\n Está calor, vá com roupas confortáveis!")
	}
}
