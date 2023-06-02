// Package weather provides tools to get information on the current weather.
package weather

// CurrentCondition represents a weather condition.
var CurrentCondition string

// CurrentLocation represents a location to check the weather.
var CurrentLocation string

// Forecast returns a string containing the weather forcast for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
