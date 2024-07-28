package repositories

import (
	"context"
	"encoding/json"
	"main/collections/domain/models"
	"main/collections/infrastructure/persistence/arangodb/entities"
	"main/collections/infrastructure/persistence/arangodb/mappers"
	"main/database"
	"main/database/arango"
	"strings"
)

type ArangodbCollectionsRepository struct {
	DbHandler        *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
	PropertiesMapper *mappers.PropertiesMapper
}

func (acr *ArangodbCollectionsRepository) GetAll(schemaName string) ([]*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaName)
	if err != nil {
		return nil, err
	}

	databaseCollections, err := database.Collections(context.Background())
	if err != nil {
		return nil, err
	}

	// Filters out all internal collections (they starts with an _)
	filteredCollections := []arango.DriverCollection{}
	for _, databaseCollection := range databaseCollections {
		name := databaseCollection.Name()
		if !strings.HasPrefix(name, "_") {
			filteredCollections = append(filteredCollections, databaseCollection)
		}
	}

	collections := []*models.Collection{}
	for _, filteredCollection := range filteredCollections {
		collectionProperties, err := filteredCollection.Properties(context.Background())
		if err != nil {
			return nil, err
		}

		// Parses entity
		var propertiesEntity entities.Properties
		acr.PropertiesMapper.ToParsedEntity(collectionProperties.Schema.Rule, &propertiesEntity)

		collection := &models.Collection{
			Name:       filteredCollection.Name(),
			Properties: acr.PropertiesMapper.ToDomainModels(propertiesEntity),
		}

		collections = append(collections, collection)
	}

	return collections, nil
}

func (acr *ArangodbCollectionsRepository) GetByName(schemaName string, name string) (*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaName)
	if err != nil {
		return nil, err
	}

	databaseCollection, err := database.Collection(context.Background(), name)
	if err != nil {
		return nil, err
	}

	collectionProperties, err := databaseCollection.Properties(context.Background())
	if err != nil {
		return nil, err
	}

	// Parses entity
	var propertiesEntity entities.Properties
	acr.PropertiesMapper.ToParsedEntity(collectionProperties.Schema.Rule, &propertiesEntity)

	collection := &models.Collection{
		Name:       databaseCollection.Name(),
		Properties: acr.PropertiesMapper.ToDomainModels(propertiesEntity),
	}

	return collection, nil
}

func (acr *ArangodbCollectionsRepository) Create(schemaName string, collection *models.Collection) (*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaName)
	if err != nil {
		return nil, err
	}

	// Maps all models to entity
	propertiesEntity := acr.PropertiesMapper.ToEntity(collection.Properties)

	// Encodes entity
	encodedPropertiesEntity, err := json.Marshal(propertiesEntity)
	if err != nil {
		return nil, err
	}

	// Loads entity for creation
	collectionProperties := arango.CreateCollectionPropertiesOptions{}
	err = collectionProperties.LoadRule(encodedPropertiesEntity)
	if err != nil {
		return nil, err
	}

	_, err = database.CreateCollection(context.Background(), collection.Name, &arango.CreateCollectionOptions{
		Schema: &collectionProperties,
	})
	if err != nil {
		return nil, err
	}

	return collection, nil
}
