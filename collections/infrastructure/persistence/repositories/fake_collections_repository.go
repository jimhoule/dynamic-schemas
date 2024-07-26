package repositories

import (
	"fmt"
	"main/collections/domain/models"
)

type FakeCollectionsRepository struct{
	Schemas map[string][]*models.Collection
}

func (fcr *FakeCollectionsRepository) Reset(schemaId string) {
	fcr.Schemas[schemaId] = []*models.Collection{}
}

func (fcr *FakeCollectionsRepository) GetAll(schemaId string) ([]*models.Collection, error) {
	_, ok := fcr.Schemas[schemaId]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	return fcr.Schemas[schemaId], nil
}

func (fcr *FakeCollectionsRepository) GetByName(schemaId string, name string) (*models.Collection, error) {
	_, ok := fcr.Schemas[schemaId]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	for _, collection := range fcr.Schemas[schemaId] {
		if collection.Name == name {
			return collection, nil
		}
	}

	return nil, nil
}

func (fcr *FakeCollectionsRepository) Create(schemaId string, collection *models.Collection) (*models.Collection, error) {
	_, ok := fcr.Schemas[schemaId]
	if !ok {
		return nil, fmt.Errorf("error: schema does not exist")
	}

	fcr.Schemas[schemaId] = append(fcr.Schemas[schemaId], collection)

	return collection, nil
}