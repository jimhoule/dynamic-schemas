package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"main/collections/domain/models"
	"main/database"
	"main/database/arango"
	"strings"
)

type ArangodbCollectionsRepository struct {
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

		var arangoProperties arango.Properties
		transcode(collectionProperties.Schema.Rule, &arangoProperties)

		collection := &models.Collection{
			Name:       filteredCollection.Name(),
			Properties: mapPropertiesToModels(arangoProperties),
		}

		collections = append(collections, collection)
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

	collectionProperties, err := databaseCollection.Properties(context.Background())
	if err != nil {
		return nil, err
	}

	var arangoProperties arango.Properties
	transcode(collectionProperties.Schema.Rule, &arangoProperties)

	collection := &models.Collection{
		Name:       databaseCollection.Name(),
		Properties: mapPropertiesToModels(arangoProperties),
	}

	return collection, nil
}

func (acr *ArangodbCollectionsRepository) Create(schemaId string, collection *models.Collection) (*models.Collection, error) {
	database, err := acr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	// Maps all properties to arnago properties model
	arangoProperties := mapPropertiesToEntities(collection.Properties)

	// Encodes arango properties
	encodedArangoProperties, err := json.Marshal(arangoProperties)
	if err != nil {
		return nil, err
	}

	// Loads arango properties for creation
	collectionProperties := arango.CreateCollectionPropertiesOptions{}
	err = collectionProperties.LoadRule(encodedArangoProperties)
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

func mapPropertiesToEntities(properties []*models.Property) arango.Properties {
	arangoProperties := arango.Properties{
		Rule: arango.Rule{
			Properties: map[string]arango.Property{},
			Required:   []string{},
		},
		Level: "strict",
	}

	for _, property := range properties {
		// If is required
		if property.IsRequired {
			arangoProperties.Rule.Required = append(arangoProperties.Rule.Required, property.Name)
		}

		arangoProperty := arango.Property{
			Type: property.Type,
		}
		// If is of type array, adds type of each item in the array
		if property.Type == "array" {
			arangoProperty.Items.Type = property.ItemType
		}

		arangoProperties.Rule.Properties[property.Name] = arangoProperty
	}

	return arangoProperties
}

func mapPropertiesToModels(arangoProperties arango.Properties) []*models.Property {
	properties := []*models.Property{}
	for key, arangoProperty := range arangoProperties.Rule.Properties {
		property := &models.Property{
			Type:       arangoProperty.Type,
			IsRequired: false,
		}

		// If is of type array
		if arangoProperty.Type == "array" {
			property.ItemType = arangoProperty.Items.Type
		}

		// If is required
		for _, requiredPropertyName := range arangoProperties.Rule.Required {
			if requiredPropertyName == key {
				property.IsRequired = true
			}
		}

		properties = append(properties, property)
	}

	return properties
}

func transcode(in any, out any) {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(in)
	json.NewDecoder(&buffer).Decode(out)
}
