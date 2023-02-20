package conv

//Fra Fahrenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {

	return (value - 32) * (5.0 / 9.0)
}

//Fra Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {

	return (value - 273.15)
}

//Fra Celcius til Fahrenheit
func CelcsiusToFahrenheit(value float64) float64 {

	return (value*(9.0/5.0) + 32.0)
}

//Fra Celcsius til Kelvin
func CelcsiusToKelvin(value float64) float64 {

	return (value + 273.15)
}

//Fra Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	return (value-273.15)*9/5 + 32
}

//Fra Fahrenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}
