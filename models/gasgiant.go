package models

type GasGiant struct {
	Name string  // name of the planet
	Desc string  // description of the planet
	Dist float64 // distance from earth in lightyears
	Rdys float64 // radius of planet in km
}

const gg_mass float64 = 0.5

func NewGasGiantPlanet(name string, desc string, dist float64, rdys float64) (*GasGiant, error) {
	err := validateGasGiantFields(name, desc, dist, rdys)
	if err != nil {
		return nil, err
	}

	return &GasGiant{
		Name: name, Desc: desc,
		Dist: dist, Rdys: rdys,
	}, nil
}

func (ggp *GasGiant) UpdateName(name string) {
	ggp.Name = name
}

func (ggp *GasGiant) GetName() string {
	return ggp.Name
}

func (ggp *GasGiant) UpdateDescription(desc string) {
	ggp.Desc = desc
}

func (ggp *GasGiant) GetDescription() string {
	return ggp.Desc
}

func (ggp *GasGiant) UpdateDistance(dist float64) {
	ggp.Dist = dist
}

func (ggp *GasGiant) GetDistance() float64 {
	return ggp.Dist
}

func (ggp *GasGiant) UpdateRadius(rdys float64) {
	ggp.Rdys = rdys
}

func (ggp *GasGiant) GetRadius() float64 {
	return ggp.Rdys
}

func (ggp *GasGiant) GetMass() float64 {
	return gg_mass
}

func (ggp *GasGiant) GetType() ExoPlanetType {
	return GasGiantPlanet
}

func validateGasGiantFields(name string, desc string, dist float64, rdys float64) error {
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

	return nil
}
