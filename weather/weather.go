package weather

// GetCurrentWeather returns the current weather for a given location.
// Input is the location and the unit of temperature. Unit can be either "Celsius" or "Fahrenheit".
func GetCurrentWeather(location string, unit string) string {
	if "tokyo" == location {
		return "{\"location\": \"Tokyo\", \"temperature\": \"10\", \"unit\": " + unit + "}"
	} else if "san francisco" == location {
		return "{\"location\": \"San Francisco\", \"temperature\": \"72\", \"unit\": " + unit + "}"
	} else if "paris" == location {
		return "{\"location\": \"Paris\", \"temperature\": \"22\", \"unit\": " + unit + "}"
	} else {
		return "{\"location\": " + location + ", \"temperature\": \"unknown\"}"
	}
}
