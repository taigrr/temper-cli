package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32.0
}

// CelsiusToKelvin converts a temperature from Celsius to Kelvin.
func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// Reading represents a single temperature sensor reading.
type Reading struct {
	Device     string  `json:"device"`
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

// NewReading creates a Reading from a device name and Celsius temperature.
func NewReading(device string, celsius float64) Reading {
	return Reading{
		Device:     device,
		Celsius:    celsius,
		Fahrenheit: CelsiusToFahrenheit(celsius),
		Kelvin:     CelsiusToKelvin(celsius),
	}
}

// FormatReading writes a single reading to the writer in the specified format.
func FormatReading(w io.Writer, r Reading, unit string) {
	fmt.Fprintln(w, FormatReadingString(r, unit))
}

// FormatReadingString returns a single formatted reading in the specified unit.
func FormatReadingString(r Reading, unit string) string {
	switch unit {
	case "fahrenheit":
		return fmt.Sprintf("%.2f", r.Fahrenheit)
	case "kelvin":
		return fmt.Sprintf("%.2f", r.Kelvin)
	default:
		return fmt.Sprintf("%.2f", r.Celsius)
	}
}

// FormatLabeledReading writes a labeled reading to the writer in the specified format.
func FormatLabeledReading(w io.Writer, r Reading, unit string) {
	fmt.Fprintf(w, "%s: %s\n", r.Device, FormatReadingString(r, unit))
}

// FormatReadingsJSON writes all readings as a JSON array to the writer.
func FormatReadingsJSON(w io.Writer, readings []Reading) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(readings)
}
