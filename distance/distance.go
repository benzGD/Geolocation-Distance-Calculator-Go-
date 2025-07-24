package distance

import (
	"math"
)

// Exported structs so they can be used in other packages
type Location struct {
	Latitude  float64
	Longitude float64
}

type Coordinate struct {
	D, M, S float64
	H       rune
}

type GeoCoordinate struct {
	Glat  Coordinate
	Glong Coordinate
}

type World struct {
	Radius float64
}

// Converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// Converts coordinate to decimal format
func (c Coordinate) Decimal() float64 {
	sign := 1.0
	if c.H == 'S' || c.H == 'W' || c.H == 's' || c.H == 'w' {
		sign = -1.0
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

// Converts coordinates to Location
func CoordinateToLocation(lat, long Coordinate) Location {
	return Location{lat.Decimal(), long.Decimal()}
}

// Calculates distance between two locations
func (w World) Distance(p1, p2 Location) float64 {
	s1, c1 := math.Sincos(rad(p1.Latitude))
	s2, c2 := math.Sincos(rad(p2.Latitude))
	clong := math.Cos(rad(p1.Longitude - p2.Longitude))
	return w.Radius * math.Acos(s1*s2+c1*c2*clong)
}
