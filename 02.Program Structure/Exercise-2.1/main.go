package main

import (
	"fmt"
	"golangplayground/tempconv"
)

func main() {

	// Convert 30 degrees celsius to each of kelvin and Fahrenheit.
	degree30C := tempconv.Celsius(30)
	fmt.Printf("Temprature degree is: %s\n", degree30C.String())
	fmt.Printf("Temprature degree is: %s\n", tempconv.CToF(degree30C))
	fmt.Printf("Temprature degree is: %s\n", tempconv.CToK(degree30C))
	// Getting Boiling degree in different scale.
	fmt.Printf("----Boiling Degree in different scale----\n")
	fmt.Printf("Celsius: %s", tempconv.BoilingC)
	fmt.Printf("\nKelvin: %s", tempconv.BoilingK)
	fmt.Printf("\nFahrenheit: %s", tempconv.CToF(tempconv.BoilingC))
	// Getting Freezing degree in different scale.
	fmt.Printf("\n----Freezing Degree in different scale----\n")
	fmt.Printf("Celsius: %s", tempconv.FreezingC)
	fmt.Printf("\nKelvin: %s", tempconv.FreezingK)
	fmt.Printf("\nFahrenheit: %s", tempconv.CToF(tempconv.FreezingC))
}


