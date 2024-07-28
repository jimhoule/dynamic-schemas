package repositories

import (
	"fmt"
	collectionModels "main/collections/domain/models"
	"main/documents/domain/models"
)

type Collection struct {
	Properties map[string]*collectionModels.Property
	Documents  []*models.Document
}

type FakeDocumentsRepository struct {
	Schemas map[string]map[string]*Collection
}

func (fdr *FakeDocumentsRepository) Reset(schemaName string, collectionName string) {
	fdr.Schemas[schemaName][collectionName].Documents = []*models.Document{}
}

func (fdr *FakeDocumentsRepository) GetAll(schemaName string, collectionName string) ([]*models.Document, error) {
	_, ok := fdr.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	_, ok = fdr.Schemas[schemaName][collectionName]
	if !ok {
		return nil, fmt.Errorf("error: collection does not exist")
	}

	return fdr.Schemas[schemaName][collectionName].Documents, nil
}

func (fdr *FakeDocumentsRepository) GetByKey(schemaName string, collectionName string, key string) (*models.Document, error) {
	_, ok := fdr.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	_, ok = fdr.Schemas[schemaName][collectionName]
	if !ok {
		return nil, fmt.Errorf("error: collection does not exist")
	}

	for _, document := range fdr.Schemas[schemaName][collectionName].Documents {
		if document.Key == key {
			return document, nil
		}
	}

	return nil, nil
}

func (fdr *FakeDocumentsRepository) Create(schemaName string, collectionName string, document *models.Document) (*models.Document, error) {
	// Checks if schema exists
	_, ok := fdr.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	// Checks if collection exists
	_, ok = fdr.Schemas[schemaName][collectionName]
	if !ok {
		return nil, fmt.Errorf("error: collection does not exist")
	}

	fdr.Schemas[schemaName][collectionName].Documents = append(fdr.Schemas[schemaName][collectionName].Documents, document)

	return document, nil
}
