package services

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/million_dollar_space_programme/exoplanets/errorhandler"
	"github.com/million_dollar_space_programme/exoplanets/models"
	"github.com/million_dollar_space_programme/exoplanets/repository"
	"github.com/million_dollar_space_programme/exoplanets/utils"
)

func CreateExoPlanet(ctx context.Context, exoPlanetInfo utils.ExoPlanetInfo) (uuid.UUID, error) {
	var (
		exoPlanet models.ExoPlanet
		err       error
	)

	switch exoPlanetInfo.Type {
	case models.GasGiantPlanet:
		exoPlanet, err = models.NewGasGiantPlanet(
			exoPlanetInfo.Name, exoPlanetInfo.Description, exoPlanetInfo.Distance, exoPlanetInfo.Radius,
		)
		if err != nil {
			return utils.NilUUID(), errorhandler.New("bad request", http.StatusBadRequest, err.Error())
		}
	case models.TerrestrialPlanet:
		exoPlanet, err = models.NewTerrestrialPlanet(
			exoPlanetInfo.Name, exoPlanetInfo.Description, exoPlanetInfo.Distance, exoPlanetInfo.Radius, exoPlanetInfo.Mass,
		)
		if err != nil {
			return utils.NilUUID(), errorhandler.New("bad request", http.StatusBadRequest, err.Error())
		}
	default:
		return utils.NilUUID(), errorhandler.New("bad request", http.StatusBadRequest, "unknown planet type")
	}

	return repository.CreateExoPlanet(ctx, exoPlanet)
}

func GetExoPlanets(ctx context.Context, input utils.FindAllInput) ([]utils.ExoPlanetInfo, error) {
	findAllOutPut, err := repository.GetExoPlanets(ctx, input)
	if err != nil {
		return []utils.ExoPlanetInfo{}, err
	}

	allExoPlanets := make([]utils.ExoPlanetInfo, 0)

	for i, exoPlanet := range findAllOutPut.AllExoPlanets {
		if exoPlanet == nil {
			continue
		}

		exoPlanetInfo := utils.ExoPlanetInfo{
			ID:          findAllOutPut.RespectiveIDs[i],
			Name:        exoPlanet.GetName(),
			Description: exoPlanet.GetDescription(),
			Distance:    exoPlanet.GetDistance(),
			Radius:      exoPlanet.GetRadius(),
			Mass:        exoPlanet.GetMass(),
		}

		switch exoPlanet.GetType() {
		case models.TerrestrialPlanet:
			exoPlanetInfo.Type = models.TerrestrialPlanet
		case models.GasGiantPlanet:
			exoPlanetInfo.Type = models.GasGiantPlanet
		}

		allExoPlanets = append(allExoPlanets, exoPlanetInfo)
	}

	return allExoPlanets, nil
}

func GetExoPlanet(ctx context.Context, id string) (utils.ExoPlanetInfo, error) {
	exoPlanet, _id, err := repository.GetExoPlanet(ctx, id)
	if err != nil {
		return utils.ExoPlanetInfo{}, err
	}

	return utils.ExoPlanetInfo{
		ID:          _id,
		Name:        exoPlanet.GetName(),
		Description: exoPlanet.GetDescription(),
		Distance:    exoPlanet.GetDistance(),
		Radius:      exoPlanet.GetRadius(),
		Mass:        exoPlanet.GetMass(),
		Type:        exoPlanet.GetType(),
	}, nil
}

func UpdateExoPlanet(ctx context.Context, id string, exoPlanetInfo utils.ExoPlanetInfo) error {
	var (
		exoPlanet models.ExoPlanet
		err       error
	)

	switch exoPlanetInfo.Type {
	case models.GasGiantPlanet:
		exoPlanet, err = models.NewGasGiantPlanet(
			exoPlanetInfo.Name, exoPlanetInfo.Description, exoPlanetInfo.Distance, exoPlanetInfo.Radius,
		)
		if err != nil {
			return errorhandler.New("bad request", http.StatusBadRequest, err.Error())
		}
	case models.TerrestrialPlanet:
		exoPlanet, err = models.NewTerrestrialPlanet(
			exoPlanetInfo.Name, exoPlanetInfo.Description, exoPlanetInfo.Distance, exoPlanetInfo.Radius, exoPlanetInfo.Mass,
		)
		if err != nil {
			return errorhandler.New("bad request", http.StatusBadRequest, err.Error())
		}
	default:
		return errorhandler.New("bad request", http.StatusBadRequest, "unknown planet type")
	}

	return repository.UpdateExoPlanet(ctx, id, exoPlanet)
}

func DeleteExoPlanet(ctx context.Context, id string) error {
	return repository.DeleteExoPlanet(ctx, id)
}

func FuelEstimation(ctx context.Context, id string, crew_members int) (float64, error) {
	return repository.FuelEstimation(ctx, id, crew_members)
}
