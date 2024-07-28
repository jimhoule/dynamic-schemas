package repositories

import (
	"fmt"
	"main/collections/domain/models"
)

type FakeCollectionsRepository struct {
	Schemas map[string][]*models.Collection
}

func (fcr *FakeCollectionsRepository) Reset(schemaName string) {
	fcr.Schemas[schemaName] = []*models.Collection{}
}

func (fcr *FakeCollectionsRepository) GetAll(schemaName string) ([]*models.Collection, error) {
	_, ok := fcr.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	return fcr.Schemas[schemaName], nil
}

func (fcr *FakeCollectionsRepository) GetByName(schemaName string, name string) (*models.Collection, error) {
	_, ok := fcr.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	for _, collection := range fcr.Schemas[schemaName] {
		if collection.Name == name {
			return collection, nil
		}
	}

	return nil, nil
}

func (fcr *FakeCollectionsRepository) Create(schemaName string, collection *models.Collection) (*models.Collection, error) {
	_, ok := fcr.Schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	fcr.Schemas[schemaName] = append(fcr.Schemas[schemaName], collection)

	return collection, nil
}
