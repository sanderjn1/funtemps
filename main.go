package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
)

func main() {
	var value float64
	var unit string

	flag.Float64Var(&value, "value", 0, "The temperature value to be converted")
	flag.StringVar(&unit, "unit", "celsius", "The unit of the input temperature (fahrenheit, kelvin, or celsius)")

	flag.Parse()

	switch unit {
	case "fahrenheit":
		result := conv.FahrenheitToCelsius(value)
		fmt.Printf("%.2f°F is equal to %.2f°C\n", value, result)
	case "kelvin":
		result := conv.KelvinToCelsius(value)
		fmt.Printf("%.2f K is equal to %.2f°C\n", value, result)
	case "celsius":
		result := conv.CelsiusToFahrenheit(value)
		fmt.Printf("%.2f°C is equal to %.2f°F\n", value, result)
		result = conv.CelsiusToKelvin(value)
		fmt.Printf("%.2f°C is equal to %.2f K\n", value, result)
	default:
		fmt.Printf("Invalid unit type: %s\n", unit)
	}
}
