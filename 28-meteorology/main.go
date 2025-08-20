package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {
	return fmt.Sprintf("%d %s", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	temp := fmt.Sprintf("%s: %s, Wind %s at %s, %d", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
	return temp + "% Humidity"
}

func main() {
	celsiusUnit := Celsius
	fahrenheitUnit := Fahrenheit

	fmt.Println(celsiusUnit.String())
	// => °C
	fmt.Println(fahrenheitUnit.String())
	// => °F
	fmt.Println(celsiusUnit)
	// Output: °C

	celsiusTemp := Temperature{
		degree: 21,
		unit:   Celsius,
	}
	fmt.Println(celsiusTemp.String())
	// => 21 °C
	fmt.Println(celsiusTemp)
	// Output: 21 °C

	fahrenheitTemp := Temperature{
		degree: 75,
		unit:   Fahrenheit,
	}
	fmt.Println(fahrenheitTemp.String())
	// => 75 °F
	fmt.Println(fahrenheitTemp)
	// Output: 75 °F
	mphUnit := MilesPerHour
	fmt.Println(mphUnit.String())
	// => mph
	fmt.Println(mphUnit)
	// Output: mph

	kmhUnit := KmPerHour
	fmt.Println(kmhUnit.String())
	// => km/h
	fmt.Println(kmhUnit)
	// Output: km/h

	windSpeedNow := Speed{
		magnitude: 18,
		unit:      KmPerHour,
	}
	fmt.Println(windSpeedNow.String())
	// => 18 km/h
	fmt.Println(windSpeedNow)
	// Output: 18 km/h

	windSpeedYesterday := Speed{
		magnitude: 22,
		unit:      MilesPerHour,
	}
	fmt.Println(windSpeedYesterday.String())
	// => 22 mph
	fmt.Println(windSpeedYesterday)
	// Output: 22 mph

	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}

	fmt.Println(sfData.String())
	// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
	fmt.Println(sfData)
	// Output: San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
}
