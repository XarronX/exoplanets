package utils

import (
	"bytes"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/million_dollar_space_programme/exoplanets/models"
)

type ExoPlanetInfo struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Distance    float64              `json:"distance"`
	Radius      float64              `json:"radius"`
	Mass        float64              `json:"mass"`
	Type        models.ExoPlanetType `json:"type"`
}

type StorableExoPlanet struct {
	ID        uuid.UUID
	ExoPlanet models.ExoPlanet
	Type      models.ExoPlanetType
}

func (sep StorableExoPlanet) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(sep)
	return buf.Bytes(), nil
}

func (sep *StorableExoPlanet) UnmarshalBinary(data string) error {
	mp := make(map[string]interface{})
	buf := bytes.NewBuffer([]byte(data))
	json.NewDecoder(buf).Decode(&mp)

	id, err := ParseUUID(mp["ID"].(string))
	if err != nil {
		return err
	}

	switch mp["Type"].(string) {
	case string(models.TerrestrialPlanet):
		sep.ID = id
		sep.Type = models.TerrestrialPlanet
		sep.ExoPlanet, _ = models.NewTerrestrialPlanet(
			mp["ExoPlanet"].(map[string]interface{})["Name"].(string),
			mp["ExoPlanet"].(map[string]interface{})["Desc"].(string),
			mp["ExoPlanet"].(map[string]interface{})["Dist"].(float64),
			mp["ExoPlanet"].(map[string]interface{})["Rdys"].(float64),
			mp["ExoPlanet"].(map[string]interface{})["Mass"].(float64),
		)
	case string(models.GasGiantPlanet):
		sep.ID = id
		sep.Type = models.GasGiantPlanet
		sep.ExoPlanet, _ = models.NewGasGiantPlanet(
			mp["ExoPlanet"].(map[string]interface{})["Name"].(string),
			mp["ExoPlanet"].(map[string]interface{})["Desc"].(string),
			mp["ExoPlanet"].(map[string]interface{})["Dist"].(float64),
			mp["ExoPlanet"].(map[string]interface{})["Rdys"].(float64),
		)
	}

	return nil
}

type FindAllInput struct { // Find multiple exoplanets input
	Size   uint64
	Offset uint64
}

type FindAllOutPut struct { // Find multiple exoplanets output
	AllExoPlanets []models.ExoPlanet
	RespectiveIDs []uuid.UUID
	CurrentCursor uint64
}
