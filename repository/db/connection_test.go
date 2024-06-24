package db

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func SetRedisClient(_client *redis.Client) {
	client = _client
}

func TestGetDBClient_Success(t *testing.T) {
	_client, mock := redismock.NewClientMock()
	mock.ExpectPing().SetVal("pong")

	SetRedisClient(_client)

	actualClient, err := GetDBClient(context.Background())

	assert.NoError(t, err)

	assert.Equal(t, client, actualClient)
}

func TestGetDBClient_Failure(t *testing.T) {
	_client, mock := redismock.NewClientMock()
	mock.ExpectPing().RedisNil()

	SetRedisClient(_client)

	_, err := GetDBClient(context.Background())

	assert.Error(t, err)
}
