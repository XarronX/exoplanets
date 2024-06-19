package models

type GasGiant struct {
	name string  // name of the planet
	desc string  // description of the planet
	dist float64 // distance from earth in lightyears
	rdys float64 // radius of planet in km
}

const (
	gg_type string  = "Gas Giant"
	gg_mass float64 = 0.5
)

func NewGasGiantPlanet(name string, desc string, dist float64, rdys float64) *GasGiant {
	return &GasGiant{
		name: name, desc: desc,
		dist: dist, rdys: rdys,
	}
}

func (ggp *GasGiant) UpdateName(name string) {
	ggp.name = name
}

func (ggp *GasGiant) GetName() string {
	return ggp.name
}

func (ggp *GasGiant) UpdateDescription(desc string) {
	ggp.desc = desc
}

func (ggp *GasGiant) GetDescription() string {
	return ggp.desc
}

func (ggp *GasGiant) UpdateDistance(dist float64) {
	ggp.dist = dist
}

func (ggp *GasGiant) GetDistance() float64 {
	return ggp.dist
}

func (ggp *GasGiant) UpdateRadius(rdys float64) {
	ggp.rdys = rdys
}

func (ggp *GasGiant) GetRadius() float64 {
	return ggp.rdys
}

func (ggp *GasGiant) GetMass() float64 {
	return gg_mass
}

func (ggp *GasGiant) GetType() string {
	return gg_type
}
