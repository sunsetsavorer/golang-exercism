// Package weather provides tools to format weather forecast.
package weather

// CurrentCondition contains weather condition.
var CurrentCondition string

// CurrentLocation contains city name.
var CurrentLocation string

// Forecast return formatted forecast string.
func Forecast(city, condition string) string {

	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
