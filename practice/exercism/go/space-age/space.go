// Package space handles space calculations
package space

const EarthYearSeconds = 31557600

// Planet years relative to one earth year
var relEarthYear = map[string]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age converts an age in seconds to the equivalent age in Earth years for the specified planet.
func Age(ageSeconds float64, planet string) float64 {
	return (ageSeconds / EarthYearSeconds) / relEarthYear[planet]
}
