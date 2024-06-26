package models

import "fmt"

/*
	1. name
	2. description
	3. distance from earth
	4. radius
	5. mass (will be provided only in case of Terrestrial type of planet)
	6. type of exoplanet : GasGiant or Terrestrial
*/

var (
	ErrInvalidName        = fmt.Errorf("invalid name")
	ErrInvalidDescription = fmt.Errorf("invalid description")
	ErrInvalidDistance    = fmt.Errorf("invalid distance")
	ErrInvalidRadius      = fmt.Errorf("invalid radius")
	ErrInvalidMass        = fmt.Errorf("invalid mass")
)

const (
	GasGiantPlanet    ExoPlanetType = "gasGiant"
	TerrestrialPlanet ExoPlanetType = "terrestrial"
)

type ExoPlanetType string

type ExoPlanet interface {
	GetName() string        // returns name of the exoplanet
	GetDescription() string // returns Description of the exoplanet
	GetDistance() float64   // returns distance from earth of the exoplanet
	GetRadius() float64     // returns radius of the exoplanet
	GetMass() float64       // returns mass of the exoplanet
	GetType() ExoPlanetType // returns type of the exoplanet
}

// g = (m/r^2)
func CalcGravity(exoPlanet ExoPlanet) float64 {
	m := exoPlanet.GetMass()
	r := exoPlanet.GetRadius()

	return m / (r * r)
}

// f = d / (g^2) * c units
func RequiredFuel(crew_members int, exoPlanet ExoPlanet) float64 {
	d := exoPlanet.GetDistance()
	g := CalcGravity(exoPlanet)

	return d / ((g * g) * float64(crew_members))
}
