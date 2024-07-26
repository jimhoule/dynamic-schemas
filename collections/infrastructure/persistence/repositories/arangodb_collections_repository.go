package repositories

import (
	"context"
	"fmt"
	"main/collections/domain/models"
	"main/database"
	"main/database/arango"
)

type ArangodbCollectionsRepository struct{
	DbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
}

func (acr *ArangodbCollectionsRepository) GetAll(schemaId string) ([]*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	databaseCollections, err := database.Collections(context.Background())
	if err != nil {
		return nil, err
	}

	collections := []*models.Collection{}
	for _, databaseCollection := range databaseCollections {
		collection := &models.Collection{
			Name: databaseCollection.Name(),
		}

		collections = append(collections, collection)

		collectionProperties, err := databaseCollection.Properties(context.Background())
		if err != nil {
			return nil, err
		}
		
		fmt.Println(collectionProperties.Schema.Rule)
	}



	return collections, nil
}

func (acr *ArangodbCollectionsRepository) GetByName(schemaId string, name string) (*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	databaseCollection, err := database.Collection(context.Background(), name)
	if err != nil {
		return nil, err
	}

	collection := &models.Collection{
		Name: databaseCollection.Name(),
	}

	return collection, nil
}

func (acr *ArangodbCollectionsRepository) Create(schemaId string, collection *models.Collection) (*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	databaseCollection, err := database.CreateCollection(context.Background(), collection.Name, nil)
	if err != nil {
		return nil, err
	}

	collectionProperties, err := databaseCollection.Properties(context.Background())
	if err != nil {
		return nil, err
	}
	
	// Maps properties to rule
	fmt.Println(collectionProperties.Schema.Rule)

	return collection, nil
}