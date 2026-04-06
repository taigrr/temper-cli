package main

import (
	"bytes"
	"encoding/json"
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"freezing point", 0, 32},
		{"boiling point", 100, 212},
		{"body temperature", 37, 98.6},
		{"absolute zero", -273.15, -459.67},
		{"negative", -40, -40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.celsius)
			if math.Abs(got-tt.expected) > 0.01 {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.celsius, got, tt.expected)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"freezing point", 0, 273.15},
		{"boiling point", 100, 373.15},
		{"absolute zero", -273.15, 0},
		{"room temperature", 20, 293.15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToKelvin(tt.celsius)
			if math.Abs(got-tt.expected) > 0.01 {
				t.Errorf("CelsiusToKelvin(%v) = %v, want %v", tt.celsius, got, tt.expected)
			}
		})
	}
}

func TestNewReading(t *testing.T) {
	reading := NewReading("sensor-1", 25.0)

	if reading.Device != "sensor-1" {
		t.Errorf("Device = %q, want %q", reading.Device, "sensor-1")
	}
	if reading.Celsius != 25.0 {
		t.Errorf("Celsius = %v, want %v", reading.Celsius, 25.0)
	}
	if math.Abs(reading.Fahrenheit-77.0) > 0.01 {
		t.Errorf("Fahrenheit = %v, want %v", reading.Fahrenheit, 77.0)
	}
	if math.Abs(reading.Kelvin-298.15) > 0.01 {
		t.Errorf("Kelvin = %v, want %v", reading.Kelvin, 298.15)
	}
}

func TestFormatReading(t *testing.T) {
	reading := NewReading("sensor-1", 25.0)

	tests := []struct {
		name     string
		unit     string
		expected string
	}{
		{"celsius", "celsius", "25.00\n"},
		{"fahrenheit", "fahrenheit", "77.00\n"},
		{"kelvin", "kelvin", "298.15\n"},
		{"default is celsius", "", "25.00\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			FormatReading(&buf, reading, tt.unit)
			if buf.String() != tt.expected {
				t.Errorf("FormatReading(%q) = %q, want %q", tt.unit, buf.String(), tt.expected)
			}
		})
	}
}

func TestFormatReadingsJSON(t *testing.T) {
	readings := []Reading{
		NewReading("sensor-1", 25.0),
		NewReading("sensor-2", 30.0),
	}

	var buf bytes.Buffer
	if err := FormatReadingsJSON(&buf, readings); err != nil {
		t.Fatalf("FormatReadingsJSON error: %v", err)
	}

	var decoded []Reading
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("JSON decode error: %v", err)
	}

	if len(decoded) != 2 {
		t.Fatalf("got %d readings, want 2", len(decoded))
	}

	if decoded[0].Device != "sensor-1" {
		t.Errorf("decoded[0].Device = %q, want %q", decoded[0].Device, "sensor-1")
	}
	if decoded[1].Celsius != 30.0 {
		t.Errorf("decoded[1].Celsius = %v, want %v", decoded[1].Celsius, 30.0)
	}
}
