package lib

import "time"

type Device struct {
	SerialNumber     string    `json:"serialNumber"`
	RegistrationDate time.Time `json:"registrationDate"`
	FirmwareVersion  string    `json:"firmware"`
	Status           string    `json:"status"`
	Temperature      []Sensor  `json:"temperature,omitempty"`
	Humidity         []Sensor  `json:"humidity,omitempty"`
	CO2              []Sensor  `json:"co2,omitempty"`
}

type Sensor struct {
	Value float64   `json:"value"`
	Date  time.Time `json:"date"`
}
