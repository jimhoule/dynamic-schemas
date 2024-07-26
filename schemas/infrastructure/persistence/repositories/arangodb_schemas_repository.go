package repositories

import (
	"context"
	"main/database"
	"main/database/arango"
	"main/schemas/domain/models"
)

type ArangodbSchemasRepository struct{
	DbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
}

func (asr *ArangodbSchemasRepository) GetById(id string) (*models.Schema, error) {
	asr.DbHandler.Client.Database(context.Background(), id)

	schema := &models.Schema{
		Id: id,
	}

	return schema, nil
}

func (asr *ArangodbSchemasRepository) Create(schema *models.Schema) (*models.Schema, error) {
	_, err := asr.DbHandler.Client.CreateDatabase(context.Background(), schema.Id, nil)
	if err != nil {
		return nil, err
	}

	return schema, nil
}