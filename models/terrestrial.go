package models

type Terrestrial struct {
	name string  // name of the planet
	desc string  // description of the planet
	dist float64 // distance from earth in parsec
	rdys float64 // radius of planet in km
	mass float64 // mass of the planet in kg
}

const t_type string = "Terrestrial"

func NewTerrestrialPlanet(name string, desc string, dist float64,
	rdys float64, mass float64) *Terrestrial {
	return &Terrestrial{
		name: name, desc: desc,
		dist: dist, rdys: rdys,
		mass: mass,
	}
}

func (tp *Terrestrial) UpdateName(name string) {
	tp.name = name
}

func (tp *Terrestrial) GetName() string {
	return tp.name
}

func (tp *Terrestrial) UpdateDescription(desc string) {
	tp.desc = desc
}

func (tp *Terrestrial) GetDescription() string {
	return tp.desc
}

func (tp *Terrestrial) UpdateDistance(dist float64) {
	tp.dist = dist
}

func (tp *Terrestrial) GetDistance() float64 {
	return tp.dist
}

func (tp *Terrestrial) UpdateRadius(rdys float64) {
	tp.rdys = rdys
}

func (tp *Terrestrial) GetRadius() float64 {
	return tp.rdys
}

func (tp *Terrestrial) UpdateMass(mass float64) {
	tp.mass = mass
}

func (tp *Terrestrial) GetMass() float64 {
	return tp.mass
}

func (tp *Terrestrial) GetType() string {
	return t_type
}
