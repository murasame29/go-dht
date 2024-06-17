package dht

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// TemperatureUnit is the temperature unit wanted, either Celsius or Fahrenheit
type TemperatureUnit int

const (
	// Celsius temperature unit
	Celsius TemperatureUnit = iota
	// Fahrenheit temperature unit
	Fahrenheit
)

// DHT struct to interface with the sensor.
// Call NewDHT to create a new one.
type DHT struct {
	pin             rpio.Pin
	temperatureUnit TemperatureUnit
	sensorType      string
	numErrors       int
	lastRead        time.Time
}
