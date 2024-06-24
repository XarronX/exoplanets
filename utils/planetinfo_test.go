package utils

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/million_dollar_space_programme/exoplanets/models"
)

func TestMarshalBinary(t *testing.T) {
	exoPlanet, _ := models.NewGasGiantPlanet("testName", "testDesc", 0.3, 1400.0)
	sep := StorableExoPlanet{
		ID:        uuid.Nil,
		ExoPlanet: exoPlanet,
		Type:      models.GasGiantPlanet,
	}

	_, err := sep.MarshalBinary()
	assert.NoError(t, err)
}

func TestUnmarshalBinary(t *testing.T) {
	t.Run("unmarshalbinary::gasGiant", func(t *testing.T) {
		gasGiant, _ := models.NewGasGiantPlanet("testName", "testDesc", 0.3, 1400.0)
		sep := StorableExoPlanet{
			ID:        uuid.Nil,
			ExoPlanet: gasGiant,
			Type:      models.GasGiantPlanet,
		}
		j, _ := json.Marshal(sep)

		err := sep.UnmarshalBinary(string(j))
		assert.NoError(t, err)
	})

	t.Run("unmarshalbinary::terrestrial", func(t *testing.T) {
		terrestrial, _ := models.NewTerrestrialPlanet("testName", "testDesc", 0.3, 1400.0, 1200.0)
		sep := StorableExoPlanet{
			ID:        uuid.Nil,
			ExoPlanet: terrestrial,
			Type:      models.TerrestrialPlanet,
		}
		j, _ := json.Marshal(sep)

		err := sep.UnmarshalBinary(string(j))
		assert.NoError(t, err)
	})

	t.Run("unmarshalbinary::invalid", func(t *testing.T) {
		terrestrial, _ := models.NewTerrestrialPlanet("testName", "testDesc", 0.3, 1400.0, 1200.0)
		sep := StorableExoPlanet{
			ID:        uuid.Nil,
			ExoPlanet: terrestrial,
			Type:      models.TerrestrialPlanet,
		}
		j, _ := json.Marshal(sep)

		str := string(j)
		str = strings.Replace(str, "00000000-0000-0000-0000-000000000000", "invalid-uuid-string", 1)

		err := sep.UnmarshalBinary(str)
		assert.Error(t, err)
	})

}
