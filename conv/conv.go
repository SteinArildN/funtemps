package conv

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(Fahrenheit float64) float64 {
	return (Fahrenheit - 32.0) * (5.0 / 9.0)
}
func FahrenheitToKelvin(Fahrenheit float64) float64 {
	return (Fahrenheit-32.0)*(5.0/9.0) + 273.15
}

func CelsiusToFahrenheit(Celsius float64) float64 {
	return Celsius*(9.0/5.0) + 32.0
}
func CelsiusToKelvin(Celsius float64) float64 {
	return Celsius + 273.15
}

func KelvintoFahrenheit(Kelvin float64) float64 {
	return (Kelvin-273.15)*(9.0/5.0) + 32.0
}
func KelvintoCelsius(Kelvin float64) float64 {
	return Kelvin - 273.15
}
