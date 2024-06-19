package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGasGiantPlanet(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"positive":     true,
			"inputName":    "testName",
			"expectedName": "testName",
			"inputDesc":    "testDesc",
			"expectedDesc": "testDesc",
			"inputDist":    0.2,
			"expectedDist": 0.2,
			"inputRyds":    1200.0,
			"expectedRdys": 1200.0,
		},
		{ // Negatice test cases
			"positive":       false,
			"inputName":      "testName",
			"unexpectedName": "notTestName",
			"inputDesc":      "testDesc",
			"unexpectedDesc": "notTestDesc",
			"inputDist":      0.2,
			"unexpectedDist": 0.3,
			"inputRyds":      1200.0,
			"unexpectedRdys": 1400.0,
		},
	}

	for _, testCase := range testCases {
		gasGiantPlanet := NewGasGiantPlanet(testCase["inputName"].(string), testCase["inputDesc"].(string),
			testCase["inputDist"].(float64), testCase["inputRyds"].(float64))

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedName"], gasGiantPlanet.name)
			assert.Equal(t, testCase["expectedDesc"], gasGiantPlanet.desc)
			assert.Equal(t, testCase["expectedDist"], gasGiantPlanet.dist)
			assert.Equal(t, testCase["expectedRdys"], gasGiantPlanet.rdys)
		} else {
			assert.NotEqual(t, testCase["unexpectedName"], gasGiantPlanet.name)
			assert.NotEqual(t, testCase["unexpectedDesc"], gasGiantPlanet.desc)
			assert.NotEqual(t, testCase["unexpectedDist"], gasGiantPlanet.dist)
			assert.NotEqual(t, testCase["unexpectedRdys"], gasGiantPlanet.rdys)
		}
	}
}

func TestGasGiant_GasGiant_UpdateName(t *testing.T) {
	initialName := "initialName"
	updatedName := "updatedName"

	testCases := []map[string]interface{}{
		{
			"positive":     true,
			"inputName":    updatedName,
			"expectedName": updatedName,
		},
		{
			"positive":       false,
			"inputName":      updatedName,
			"unexpectedName": initialName,
		},
	}

	for _, testCase := range testCases {
		gasGiantPlanet := NewGasGiantPlanet(initialName, "someDesc", 0.3, 1400.0)
		gasGiantPlanet.UpdateName(updatedName)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedName"], gasGiantPlanet.name)
		} else {
			assert.NotEqual(t, testCase["unexpectedName"], gasGiantPlanet.name)
		}
	}
}

func TestGasGiant_GetName(t *testing.T) {
	name := "testName"
	gasGiantPlanet := NewGasGiantPlanet(name, "someDesc", 0.3, 1400.0)

	outputName := gasGiantPlanet.GetName()
	assert.Equal(t, name, outputName)
}

func TestGasGiant_UpdateDesc(t *testing.T) {
	initialDesc := "initialDesc"
	updatedDesc := "updatedDesc"

	testCases := []map[string]interface{}{
		{
			"positive":     true,
			"inputDesc":    updatedDesc,
			"expectedDesc": updatedDesc,
		},
		{
			"positive":       false,
			"inputDesc":      updatedDesc,
			"unexpectedDesc": initialDesc,
		},
	}

	for _, testCase := range testCases {
		gasGiantPlanet := NewGasGiantPlanet("someName", initialDesc, 0.3, 1400.0)
		gasGiantPlanet.UpdateDescription(updatedDesc)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedDesc"], gasGiantPlanet.desc)
		} else {
			assert.NotEqual(t, testCase["unexpectedDesc"], gasGiantPlanet.desc)
		}
	}
}

func TestGasGiant_GetDescription(t *testing.T) {
	desc := "testName"
	gasGiantPlanet := NewGasGiantPlanet("someName", desc, 0.3, 1400.0)

	outputDesc := gasGiantPlanet.GetDescription()
	assert.Equal(t, desc, outputDesc)
}

func TestGasGiant_UpdateDistance(t *testing.T) {
	initialDist := 0.3
	updatedDist := 0.4

	testCases := []map[string]interface{}{
		{
			"positive":     true,
			"inputDist":    updatedDist,
			"expectedDist": updatedDist,
		},
		{
			"positive":       false,
			"inputDist":      updatedDist,
			"unexpectedDist": initialDist,
		},
	}

	for _, testCase := range testCases {
		gasGiantPlanet := NewGasGiantPlanet("someName", "someDesc", initialDist, 1400.0)
		gasGiantPlanet.UpdateDistance(updatedDist)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedDist"], gasGiantPlanet.dist)
		} else {
			assert.NotEqual(t, testCase["unexpectedDist"], gasGiantPlanet.dist)
		}
	}
}

func TestGasGiant_GetDistance(t *testing.T) {
	dist := 0.3
	gasGiantPlanet := NewGasGiantPlanet("someName", "someDesc", dist, 1400.0)

	outputDist := gasGiantPlanet.GetDistance()
	assert.Equal(t, dist, outputDist)
}

func TestGasGiant_UpdateRadius(t *testing.T) {
	initialRdys := 0.3
	updatedRdys := 0.4

	testCases := []map[string]interface{}{
		{
			"positive":     true,
			"inputRdys":    updatedRdys,
			"expectedRdys": updatedRdys,
		},
		{
			"positive":       false,
			"inputRdys":      updatedRdys,
			"unexpectedRdys": initialRdys,
		},
	}

	for _, testCase := range testCases {
		gasGiantPlanet := NewGasGiantPlanet("someName", "someDesc", 0.3, initialRdys)
		gasGiantPlanet.UpdateRadius(updatedRdys)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedRdys"], gasGiantPlanet.rdys)
		} else {
			assert.NotEqual(t, testCase["unexpectedRdys"], gasGiantPlanet.rdys)
		}
	}
}

func TestGasGiant_GetRadius(t *testing.T) {
	rdys := 1400.0
	gasGiantPlanet := NewGasGiantPlanet("someName", "someDesc", 0.3, rdys)

	outputRdys := gasGiantPlanet.GetRadius()
	assert.Equal(t, rdys, outputRdys)
}

func TestGasGiant_GetMass(t *testing.T) {
	gasGiantPlanet := NewGasGiantPlanet("someName", "someDesc", 0.3, 1400.0)

	outputMass := gasGiantPlanet.GetMass()
	assert.Equal(t, gg_mass, outputMass)
}

func TestGasGiant_GetType(t *testing.T) {
	gasGiantPlanet := NewGasGiantPlanet("someName", "someDesc", 0.3, 1400.0)

	outputType := gasGiantPlanet.GetType()
	assert.Equal(t, gg_type, outputType)
}
