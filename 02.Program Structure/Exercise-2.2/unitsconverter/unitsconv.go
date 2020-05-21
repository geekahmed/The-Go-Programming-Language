package unitsconverter

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Feet float64
type Pounds float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	AbsoluteZeroK Kelvin = 0
	FreezingK Kelvin = 273.15
	BoilingK Kelvin = 373.15
)
func (c Celsius) String() string  {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahrenheit) String() string  {
	return fmt.Sprintf("%g°F", f)
}
func (k Kelvin) String() string  {
	return fmt.Sprintf("%g°K", k)
}

func (m Meter) String() string  {
	return fmt.Sprintf("%gm", m)
}
func (f Feet) String() string  {
	return fmt.Sprintf("%g feet", f)
}
func (k Kilogram) String() string  {
	return fmt.Sprintf("%g Kg", k)
}
func (p Pounds) String() string  {
	return fmt.Sprintf("%g lb", p)
}
