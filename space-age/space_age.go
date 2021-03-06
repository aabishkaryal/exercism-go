package space

type Planet string

var EARTH_YEAR float64 = 31557600

var yearConversion map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the number of years for given number of seconds on given planet.
func Age(secs float64, planet Planet) float64 {
	return secs / EARTH_YEAR / yearConversion[planet]
}
