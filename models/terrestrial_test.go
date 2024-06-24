package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTerrestrialPlanet(t *testing.T) {
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
			"inputMass":    1300.0,
			"expectedMass": 1300.0,
		},
		{
			"positive":       false,
			"inputName":      "testName",
			"unexpectedName": "notTestName",
			"inputDesc":      "testDesc",
			"unexpectedDesc": "notTestDesc",
			"inputDist":      0.2,
			"unexpectedDist": 0.3,
			"inputRyds":      1200.0,
			"unexpectedRdys": 1400.0,
			"inputMass":      1300.0,
			"unexpectedMass": 1600.0,
		},
	}

	for _, testCase := range testCases {
		terrestrialPlanet, _ := NewTerrestrialPlanet(testCase["inputName"].(string), testCase["inputDesc"].(string),
			testCase["inputDist"].(float64), testCase["inputRyds"].(float64), testCase["inputMass"].(float64))

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedName"], terrestrialPlanet.Name)
			assert.Equal(t, testCase["expectedDesc"], terrestrialPlanet.Desc)
			assert.Equal(t, testCase["expectedDist"], terrestrialPlanet.Dist)
			assert.Equal(t, testCase["expectedRdys"], terrestrialPlanet.Rdys)
			assert.Equal(t, testCase["expectedMass"], terrestrialPlanet.Mass)
		} else {
			assert.NotEqual(t, testCase["unexpectedName"], terrestrialPlanet.Name)
			assert.NotEqual(t, testCase["unexpectedDesc"], terrestrialPlanet.Desc)
			assert.NotEqual(t, testCase["unexpectedDist"], terrestrialPlanet.Dist)
			assert.NotEqual(t, testCase["unexpectedRdys"], terrestrialPlanet.Rdys)
			assert.NotEqual(t, testCase["unexpectedMass"], terrestrialPlanet.Mass)
		}
	}
}

func TestTerrestrial_UpdateName(t *testing.T) {
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
		terrestrialPlanet, _ := NewTerrestrialPlanet(initialName, "someDesc", 0.3, 1400.0, 1200.0)
		terrestrialPlanet.UpdateName(updatedName)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedName"], terrestrialPlanet.Name)
		} else {
			assert.NotEqual(t, testCase["unexpectedName"], terrestrialPlanet.Name)
		}
	}
}

func TestTerrestrial_GetName(t *testing.T) {
	name := "testName"
	terrestrialPlanet, _ := NewTerrestrialPlanet(name, "someDesc", 0.3, 1400.0, 1200.0)

	outputName := terrestrialPlanet.GetName()
	assert.Equal(t, name, outputName)
}

func TestTerrestrial_UpdateDesc(t *testing.T) {
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
		terrestrialPlanet, _ := NewTerrestrialPlanet("someName", initialDesc, 0.3, 1400.0, 1200.0)
		terrestrialPlanet.UpdateDescription(updatedDesc)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedDesc"], terrestrialPlanet.Desc)
		} else {
			assert.NotEqual(t, testCase["unexpectedDesc"], terrestrialPlanet.Desc)
		}
	}
}

func TestTerrestrial_GetDescription(t *testing.T) {
	desc := "testName"
	terrestrialPlanet, _ := NewTerrestrialPlanet("someName", desc, 0.3, 1400.0, 1200.0)

	outputDesc := terrestrialPlanet.GetDescription()
	assert.Equal(t, desc, outputDesc)
}

func TestTerrestrial_UpdateDistance(t *testing.T) {
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
		terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", initialDist, 1400.0, 1200.0)
		terrestrialPlanet.UpdateDistance(updatedDist)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedDist"], terrestrialPlanet.Dist)
		} else {
			assert.NotEqual(t, testCase["unexpectedDist"], terrestrialPlanet.Dist)
		}
	}
}

func TestTerrestrial_GetDistance(t *testing.T) {
	dist := 0.3
	terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", dist, 1400.0, 1200.0)

	outputDist := terrestrialPlanet.GetDistance()
	assert.Equal(t, dist, outputDist)
}

func TestTerrestrial_UpdateRadius(t *testing.T) {
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
		terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", 0.3, initialRdys, 1200.0)
		terrestrialPlanet.UpdateRadius(updatedRdys)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedRdys"], terrestrialPlanet.Rdys)
		} else {
			assert.NotEqual(t, testCase["unexpectedRdys"], terrestrialPlanet.Rdys)
		}
	}
}

func TestTerrestrial_GetRadius(t *testing.T) {
	rdys := 1400.0
	terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", 0.3, rdys, 1200.0)

	outputRdys := terrestrialPlanet.GetRadius()
	assert.Equal(t, rdys, outputRdys)
}

func TestTerrestrial_UpdateMass(t *testing.T) {
	initialMass := 1600.0
	updatedMass := 1300.0

	testCases := []map[string]interface{}{
		{
			"positive":     true,
			"inputMass":    updatedMass,
			"expectedMass": updatedMass,
		},
		{
			"positive":       false,
			"inputMass":      updatedMass,
			"unexpectedMass": initialMass,
		},
	}

	for _, testCase := range testCases {
		terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", 0.3, 1400.0, initialMass)
		terrestrialPlanet.UpdateMass(updatedMass)

		if testCase["positive"].(bool) {
			assert.Equal(t, testCase["expectedMass"], terrestrialPlanet.Mass)
		} else {
			assert.NotEqual(t, testCase["unexpectedMass"], terrestrialPlanet.Mass)
		}
	}
}

func TestTerrestrial_GetMass(t *testing.T) {
	mass := 1200.0
	terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", 0.3, 1400.0, mass)

	outputMass := terrestrialPlanet.GetMass()
	assert.Equal(t, mass, outputMass)
}

func TestTerrestrial_GetType(t *testing.T) {
	terrestrialPlanet, _ := NewTerrestrialPlanet("someName", "someDesc", 0.3, 1400.0, 1200.0)

	outputType := terrestrialPlanet.GetType()
	assert.Equal(t, TerrestrialPlanet, outputType)
}

func TestTerrestrialInvalidFields(t *testing.T) {
	_, err := NewTerrestrialPlanet("", "someDesc", 0.3, 1400.0, 1200.0)
	assert.Equal(t, ErrInvalidName, err)

	_, err = NewTerrestrialPlanet("someName", "", 0.3, 1400.0, 1200.0)
	assert.Equal(t, ErrInvalidDescription, err)

	_, err = NewTerrestrialPlanet("someName", "someDesc", 0.0, 1400.0, 1200.0)
	assert.Equal(t, ErrInvalidDistance, err)

	_, err = NewTerrestrialPlanet("someName", "someDesc", 0.3, 0.0, 1200.0)
	assert.Equal(t, ErrInvalidRadius, err)

	_, err = NewTerrestrialPlanet("someName", "someDesc", 0.3, 1400.0, 0.0)
	assert.Equal(t, ErrInvalidMass, err)
}
