package repositories

import (
	"main/schemas/domain/models"
)

type FakeSchemasRepository struct{
	Schemas map[string]*models.Schema
}

func (fsr *FakeSchemasRepository) Reset() {
	fsr.Schemas = map[string]*models.Schema{}
}


func (fsr *FakeSchemasRepository) GetById(id string) (*models.Schema, error) {
	schema, ok := fsr.Schemas[id]
	if !ok {
		return nil, nil
	}

	return schema, nil
}

func (fsr *FakeSchemasRepository) Create(schema *models.Schema) (*models.Schema, error) {
	fsr.Schemas[schema.Id] = schema

	return schema, nil
}