package repositories

import (
	"main/schemas/domain/models"
)

type FakeSchemasRepository struct {
	Schemas map[string]*models.Schema
}

func (fsr *FakeSchemasRepository) Reset() {
	fsr.Schemas = map[string]*models.Schema{}
}

func (fsr *FakeSchemasRepository) GetByName(name string) (*models.Schema, error) {
	schema, ok := fsr.Schemas[name]
	if !ok {
		return nil, nil
	}

	return schema, nil
}

func (fsr *FakeSchemasRepository) Create(schema *models.Schema) (*models.Schema, error) {
	fsr.Schemas[schema.Name] = schema

	return schema, nil
}
