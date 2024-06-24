package models

type Terrestrial struct {
	Name string  // name of the planet
	Desc string  // description of the planet
	Dist float64 // distance from earth in parsec
	Rdys float64 // radius of planet in km
	Mass float64 // mass of the planet in kg
}

func NewTerrestrialPlanet(name string, desc string, dist float64,
	rdys float64, mass float64) (*Terrestrial, error) {
	err := validateTerrestrialFields(name, desc, dist, rdys, mass)
	if err != nil {
		return nil, err
	}

	return &Terrestrial{
		Name: name, Desc: desc,
		Dist: dist, Rdys: rdys,
		Mass: mass,
	}, nil
}

func (tp *Terrestrial) UpdateName(name string) {
	tp.Name = name
}

func (tp *Terrestrial) GetName() string {
	return tp.Name
}

func (tp *Terrestrial) UpdateDescription(desc string) {
	tp.Desc = desc
}

func (tp *Terrestrial) GetDescription() string {
	return tp.Desc
}

func (tp *Terrestrial) UpdateDistance(dist float64) {
	tp.Dist = dist
}

func (tp *Terrestrial) GetDistance() float64 {
	return tp.Dist
}

func (tp *Terrestrial) UpdateRadius(rdys float64) {
	tp.Rdys = rdys
}

func (tp *Terrestrial) GetRadius() float64 {
	return tp.Rdys
}

func (tp *Terrestrial) UpdateMass(mass float64) {
	tp.Mass = mass
}

func (tp *Terrestrial) GetMass() float64 {
	return tp.Mass
}

func (tp *Terrestrial) GetType() ExoPlanetType {
	return TerrestrialPlanet
}

func validateTerrestrialFields(name string, desc string, dist float64, rdys float64, mass float64) error {
	if name == "" {
		return ErrInvalidName
	}
	if desc == "" {
		return ErrInvalidDescription
	}
	if dist <= 0.0 {
		return ErrInvalidDistance
	}
	if rdys <= 0.0 {
		return ErrInvalidRadius
	}
	if mass <= 0.0 {
		return ErrInvalidMass
	}

	return nil
}
