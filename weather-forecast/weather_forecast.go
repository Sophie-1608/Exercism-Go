// Package weather provides tools to forecast the current weather condition.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string 
	// CurrentLocation represents the current location studied.
	CurrentLocation  string
)

// Forecast returns a string giving the current weather condition for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
