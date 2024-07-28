package services

import (
	"fmt"
	"main/collections/application/services"
	collectionsModels "main/collections/domain/models"
	"main/documents/application/payloads"
	"main/documents/application/ports"
	"main/documents/domain/factories"
	"main/documents/domain/models"
	"reflect"
)

type DocumentsService struct {
	DocumentsFactory    *factories.DocumentsFactory
	DocumentsRepository ports.DocumentsRepositoryPort
	CollectionsService  *services.CollectionsService
}

func (ds *DocumentsService) GetAll(schemaName string, collectionName string) ([]*models.Document, error) {
	return ds.DocumentsRepository.GetAll(schemaName, collectionName)
}

func (ds *DocumentsService) GetByKey(schemaName string, collectionName string, key string) (*models.Document, error) {
	return ds.DocumentsRepository.GetByKey(schemaName, collectionName, key)
}

func (ds *DocumentsService) Create(createDocumentPayload *payloads.CreateDocumentPayload) (*models.Document, error) {
	collection, err := ds.CollectionsService.GetByName(createDocumentPayload.SchemaName, createDocumentPayload.CollectionName)
	if err != nil {
		return nil, err
	}

	if len(collection.Properties) > 0 {
		// Gets properties map
		propertiesMap := map[string]*collectionsModels.Property{}
		for _, property := range collection.Properties {
			propertiesMap[property.Name] = property
		}

		// Validates properties
		err = validate(propertiesMap, createDocumentPayload.Body)
		if err != nil {
			return nil, err
		}
	}

	document := ds.DocumentsFactory.Create(createDocumentPayload.Body)

	return ds.DocumentsRepository.Create(
		createDocumentPayload.SchemaName,
		createDocumentPayload.CollectionName,
		document,
	)
}

func validate(collectionProperties map[string]*collectionsModels.Property, documentBody map[string]any) error {
	for key := range documentBody {
		// If provided body property is defined
		_, ok := collectionProperties[key]
		if !ok {
			return fmt.Errorf("error: %s property in body is not defined", key)
		}
	}

	for key, property := range collectionProperties {
		bodyProperty, ok := documentBody[key]

		// If property is required
		if property.IsRequired {
			if !ok {
				return fmt.Errorf("error: %s property is required", key)
			}
		}

		// If property is not required and was not provided
		if !ok {
			continue
		}

		// If property is of wrong type
		bodyPropertyType := reflect.TypeOf(bodyProperty).Kind().String()
		if property.Type != bodyPropertyType {
			return fmt.Errorf("error: %s property must be of type %s", key, property.Type)
		}
	}

	return nil
}
