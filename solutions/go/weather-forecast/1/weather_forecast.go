//Package weather provides tools forecasting the current weather condition of various cities in Goblinocus.
package weather
// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents cities in Goblinocus.
var CurrentLocation string
// Forecast returns a string value which represent an information about the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
