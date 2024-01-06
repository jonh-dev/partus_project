package db_test

import (
	"context"
	"os"
	"testing"
	"time"

	db "github.com/jonh-dev/partus_users/internal/config"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestDBConnection(t *testing.T) {
	envGetter := db.NewEnvVarGetter()

	dbService, err := db.NewDBService(envGetter)

	assert.NoError(t, err)
	assert.NotNil(t, dbService)
	assert.NotNil(t, dbService.Client)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = dbService.Client.Ping(ctx, readpref.Primary())
	assert.NoError(t, err)

	assert.Equal(t, os.Getenv("DB_NAME"), dbService.DBName)
}
