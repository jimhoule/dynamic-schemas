package repositories

import (
	"context"
	"fmt"
	"main/database"
	"main/database/arango"
	"main/schemas/domain/models"
)

type ArangodbSchemasRepository struct {
	DbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
}

func (asr *ArangodbSchemasRepository) GetByName(name string) (*models.Schema, error) {
	database, err := asr.DbHandler.Client.Database(context.Background(), name)
	if err != nil {
		return nil, err
	}

	info, err := database.Info(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Println(info.ID)

	schema := &models.Schema{
		Id:   info.ID,
		Name: name,
	}

	return schema, nil
}

func (asr *ArangodbSchemasRepository) Create(schema *models.Schema) (*models.Schema, error) {
	_, err := asr.DbHandler.Client.CreateDatabase(context.Background(), schema.Name, nil)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
