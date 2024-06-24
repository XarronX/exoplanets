package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcGravity(t *testing.T) {
	planet, _ := NewGasGiantPlanet("someName", "someDesc", 0.1, 0.1)
	assert.Equal(t, CalcGravity(planet), 49.99999999999999)
}

func TestRequiredFuel(t *testing.T) {
	planet, _ := NewGasGiantPlanet("someName", "someDesc", 0.1, 0.1)
	assert.Equal(t, RequiredFuel(7, planet), 5.714285714285717e-06)
}
