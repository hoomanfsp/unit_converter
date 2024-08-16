package proceed

import (
	"fmt"
)

// ConvertWeight converts weight between different units.
func ConvertWeight(value float64, fromUnit, toUnit string) (string, error) {
	// Define conversion factors to grams
	unitToGrams := map[string]float64{
		"gram":     1,
		"kilogram": 1000,
		"mg":       0.001,
		"cg":       0.01,
	}

	// Check if the provided units are valid
	fromFactor, fromExists := unitToGrams[fromUnit]
	toFactor, toExists := unitToGrams[toUnit]
	if !fromExists || !toExists {
		return "", fmt.Errorf("unsupported unit conversion: from %s to %s", fromUnit, toUnit)
	}

	// Convert value to grams
	valueInGrams := value * fromFactor

	// Convert grams to target unit
	result := valueInGrams / toFactor

	// Format result to 2 decimal places
	return fmt.Sprintf("%.2f", result), nil
}

func ConvertLength(value float64, fromUnit, toUnit string) (string, error) {
	// Define conversion factors to meters
	unitToMeters := map[string]float64{
		"meter":      1,
		"kilometer":  1000,
		"yard":       0.9144,
		"foot":       0.3048,
		"centimeter": 0.01,
		"millimeter": 0.001,
		"mile":       1609.344,
		"inch":       0.0254,
	}

	// Check if the provided units are valid
	fromFactor, fromExists := unitToMeters[fromUnit]
	toFactor, toExists := unitToMeters[toUnit]
	if !fromExists || !toExists {
		return "", fmt.Errorf("unsupported unit conversion: from %s to %s", fromUnit, toUnit)
	}

	// Convert value to meters
	valueInMeters := value * fromFactor

	// Convert meters to target unit
	result := valueInMeters / toFactor

	// Format result to 2 decimal places
	return fmt.Sprintf("%.2f", result), nil
}
func ConvertTemperature(value float64, fromUnit, toUnit string) (string, error) {
	// Convert value to Celsius first
	var tempInCelsius float64

	switch fromUnit {
	case "Celsius":
		tempInCelsius = value
	case "Fahrenheit":
		tempInCelsius = (value - 32) * 5 / 9
	case "Kelvin":
		tempInCelsius = value - 273.15
	case "Rankine":
		tempInCelsius = (value - 491.67) * 5 / 9
	case "Reaumur":
		tempInCelsius = value * 5 / 4
	default:
		return "", fmt.Errorf("unsupported fromUnit: %s", fromUnit)
	}

	// Convert Celsius to target unit
	var result float64

	switch toUnit {
	case "Celsius":
		result = tempInCelsius
	case "Fahrenheit":
		result = tempInCelsius*9/5 + 32
	case "Kelvin":
		result = tempInCelsius + 273.15
	case "Rankine":
		result = (tempInCelsius + 273.15) * 9 / 5
	case "Reaumur":
		result = tempInCelsius * 4 / 5
	default:
		return "", fmt.Errorf("unsupported toUnit: %s", toUnit)
	}

	// Format result to 2 decimal places
	return fmt.Sprintf("%.2f", result), nil
}
