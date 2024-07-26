package repositories

import (
	"fmt"
	collectionModels "main/collections/domain/models"
	"main/documents/domain/models"
	"reflect"
)

type Collection struct{
	Properties map[string]*collectionModels.Property
	Documents []*models.Document
}

type FakeDocumentsRepository struct{
	Schemas map[string]map[string]*Collection
}

func (fdr *FakeDocumentsRepository) Reset(schemaId string, collectionName string) {
	fdr.Schemas[schemaId][collectionName].Documents = []*models.Document{}
}

func (fdr *FakeDocumentsRepository) GetAll(schemaId string, collectionName string) ([]*models.Document, error) {
	_, ok := fdr.Schemas[schemaId]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	_, ok = fdr.Schemas[schemaId][collectionName]
	if !ok {
		return nil, fmt.Errorf("error: collection does not exist")
	}

	return fdr.Schemas[schemaId][collectionName].Documents, nil
}

func (fdr *FakeDocumentsRepository) GetByKey(schemaId string, collectionName string, key string) (*models.Document, error) {
	_, ok := fdr.Schemas[schemaId]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	_, ok = fdr.Schemas[schemaId][collectionName]
	if !ok {
		return nil, fmt.Errorf("error: collection does not exist")
	}

	for _, document := range fdr.Schemas[schemaId][collectionName].Documents {
		if document.Key == key {
			return document, nil
		}
	}

	return nil, nil
}

func (fdr *FakeDocumentsRepository) Create(schemaId string, collectionName string, document *models.Document) (*models.Document, error) {
	// Checks if schema exists
	_, ok := fdr.Schemas[schemaId]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	// Checks if collection exists
	_, ok = fdr.Schemas[schemaId][collectionName]
	if !ok {
		return nil, fmt.Errorf("error: collection does not exist")
	}

	// If some properties have been defined
	if len(fdr.Schemas[schemaId][collectionName].Properties) > 0 {
		for key := range document.Body {
			// If provided body property is defined
			_, ok := fdr.Schemas[schemaId][collectionName].Properties[key]
			if !ok {
				return nil, fmt.Errorf("error: %s property in body is not defined", key)
			}
		}

		for key, property := range fdr.Schemas[schemaId][collectionName].Properties {
			bodyProperty, ok := document.Body[key]

			// If property is required
			if property.IsRequired {
				if !ok {
					return nil, fmt.Errorf("error: %s property is required", key)
				}
			}

			// If property is not required and was not provided
			if !ok {
				continue
			}

			// If property is of wrong type
			bodyPropertyType := reflect.TypeOf(bodyProperty).Kind().String()
			if property.Type != bodyPropertyType {
				return nil, fmt.Errorf("error: %s property must be of type %s", key, property.Type)
			}
		}
	}

	fdr.Schemas[schemaId][collectionName].Documents = append(fdr.Schemas[schemaId][collectionName].Documents, document)

	return document, nil
}