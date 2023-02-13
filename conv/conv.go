package conv

func FarhenheitToCelsius(value float64) float64 {

	return (value - 32) * (5.0 / 9.0)
}

func KelvinToCelsius(value float64) float64 {

	return (value - 273.15)
}

func CelcsiusToFahrenheit(value float64) float64 {

	return (value*(9/5) + 32)
}

func CelcsiusToKelvin(value float64) float64 {

	return (value + 273.15)
}
