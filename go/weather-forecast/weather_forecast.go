// Package weather provides local weather conditions.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents current city.
var CurrentLocation string

// Forecast returns a string containing the current weather condition of a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
