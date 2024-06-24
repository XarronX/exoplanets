package repository

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/million_dollar_space_programme/exoplanets/errorhandler"
	"github.com/million_dollar_space_programme/exoplanets/models"
	"github.com/million_dollar_space_programme/exoplanets/repository/db"
	"github.com/million_dollar_space_programme/exoplanets/utils"
)

var dbClient *redis.Client

func init() {
	var err error
	dbClient, err = db.GetDBClient(context.Background())
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}
}

func CreateExoPlanet(ctx context.Context, exoPlanet models.ExoPlanet) (uuid.UUID, error) {
	newUUID := utils.GenerateUUID()
	storable := utils.StorableExoPlanet{
		ID:        newUUID,
		ExoPlanet: exoPlanet,
		Type:      exoPlanet.GetType(),
	}

	storableBytes, err := storable.MarshalBinary()
	if err != nil {
		return utils.NilUUID(), errorhandler.New("internal server error", 500, err.Error())
	}

	txn := dbClient.TxPipeline() // create transaction

	resp := txn.SetNX(ctx, newUUID.String(), storableBytes, 0) // set exoplanet info against uuid in the txn
	if err := resp.Err(); err != nil {
		txn.Discard()
		return utils.NilUUID(), errorhandler.New("internal server error", 500, err.Error())
	}

	if err := txn.SAdd(ctx, "exoplanets", newUUID.String()).Err(); err != nil { // add uuid to exoplanets set in the txn
		txn.Discard()
		return utils.NilUUID(), errorhandler.New("internal server error", 500, err.Error())
	}

	if _, err := txn.Exec(ctx); err != nil { // now execute the txn as a whole
		txn.Discard()
		return utils.NilUUID(), errorhandler.New("internal server error", 500, err.Error())
	}

	return newUUID, nil
}

func GetExoPlanets(ctx context.Context, input utils.FindAllInput) (utils.FindAllOutPut, error) {
	resp := dbClient.SScan(ctx, "exoplanets", input.Offset, "*", int64(input.Size)) // get exoplanets from exoplanets set in the txn
	ids, cursor, err := resp.Result()
	if err != nil {
		return utils.FindAllOutPut{}, errorhandler.New("internal server error", 500, err.Error())
	}

	if len(ids) == 0 { // no exoplanets found in the db
		return utils.FindAllOutPut{
			AllExoPlanets: []models.ExoPlanet{},
		}, nil
	}

	vals, err := dbClient.MGet(ctx, ids...).Result() // get exoplanets from exoplanets set in the txn
	if err != nil {
		return utils.FindAllOutPut{}, errorhandler.New("internal server error", 500, err.Error())
	}

	ValsLen := len(vals)

	allExoPlanets := make([]models.ExoPlanet, ValsLen)
	respectiveIds := make([]uuid.UUID, ValsLen)

	for i, val := range vals {
		storable := new(utils.StorableExoPlanet)
		if val == nil {
			continue
		}

		value := val.(string)

		err := storable.UnmarshalBinary(value)
		if err != nil {
			return utils.FindAllOutPut{}, errorhandler.New("internal server error", 500, err.Error())
		}

		allExoPlanets[i] = storable.ExoPlanet
		respectiveIds[i] = storable.ID
	}

	return utils.FindAllOutPut{
		AllExoPlanets: allExoPlanets,
		RespectiveIDs: respectiveIds,
		CurrentCursor: cursor,
	}, nil
}

func GetExoPlanet(ctx context.Context, id string) (models.ExoPlanet, uuid.UUID, error) {
	var exoplanet models.ExoPlanet

	entity, err := dbClient.Get(ctx, id).Result()
	if errors.Is(err, redis.Nil) {
		return exoplanet, utils.NilUUID(), errorhandler.New("entity not found", 404, err.Error())
	} else if err != nil {
		return exoplanet, utils.NilUUID(), errorhandler.New("internal server error", 500, err.Error())
	}

	storable := new(utils.StorableExoPlanet)
	err = storable.UnmarshalBinary(entity)
	if err != nil {
		return exoplanet, utils.NilUUID(), errorhandler.New("internal server error", 500, err.Error())
	}

	return storable.ExoPlanet, storable.ID, nil
}

func UpdateExoPlanet(ctx context.Context, id string, exoPlanet models.ExoPlanet) error {
	_, _id, err := GetExoPlanet(ctx, id)
	if err != nil {
		return err
	}

	storable := utils.StorableExoPlanet{
		ID:        _id,
		ExoPlanet: exoPlanet,
		Type:      exoPlanet.GetType(),
	}

	storableBytes, err := storable.MarshalBinary()
	if err != nil {
		return errorhandler.New("internal server error", 500, err.Error())
	}

	return dbClient.SetXX(ctx, id, storableBytes, 0).Err()
}

func DeleteExoPlanet(ctx context.Context, id string) error {
	_, _, err := GetExoPlanet(ctx, id)
	if err != nil {
		return err
	}

	return dbClient.Del(ctx, id).Err()
}

func FuelEstimation(ctx context.Context, id string, crew_members int) (float64, error) {
	exoplanet, _, err := GetExoPlanet(ctx, id)
	if err != nil {
		return 0.0, err
	}

	return models.RequiredFuel(crew_members, exoplanet), nil
}
