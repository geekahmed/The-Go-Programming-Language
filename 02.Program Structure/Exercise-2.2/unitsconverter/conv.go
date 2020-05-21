package unitsconverter
func CToF (c Celsius) Fahrenheit {
	return Fahrenheit(c * (9.0/5) + 32)
}

func FToC (f Fahrenheit) Celsius{
	return Celsius((f - 32) * 5/9.0)
}

func CToK(c Celsius) Kelvin  {
	return Kelvin(c + 273.15)
}
func KToC(k Kelvin) Celsius  {
	return Celsius(k - 273.15)
}
func MeterToFeet(m Meter) Feet  {
	return Feet(m * 3.28084)
}
func FeetToMeter(feet Feet) Meter {
	return Meter(feet / 3.28084)
}
func KilogramToPounds(kg Kilogram) Pounds {
	return Pounds(kg * 2.20462)
}
func PoundsToKilogram(po Pounds) Kilogram {
	return Kilogram(po / 2.20462)
}