package main

import (
	"fmt"
	"golangplayground/unitsconverter"
	"os"
	"strconv"
)

func main() {

		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := unitsconverter.Fahrenheit(t)
			c := unitsconverter.Celsius(t)

			meter := unitsconverter.Meter(t)
			foot := unitsconverter.Feet(t)

			kilogram := unitsconverter.Kilogram(t)
			pound := unitsconverter.Pounds(t)

			fmt.Printf("%s = %s, %s = %s\n", f, unitsconverter.FToC(f), c, unitsconverter.CToF(c))
			fmt.Printf("%s = %s, %s = %s\n", meter, unitsconverter.MeterToFeet(meter), foot, unitsconverter.FeetToMeter(foot))
			fmt.Printf("%s = %s, %s = %s\n", kilogram, unitsconverter.KilogramToPounds(kilogram), pound, unitsconverter.PoundsToKilogram(pound))


	}
}
