package services

import (
	"main/schemas/application/payloads"
	"main/schemas/application/ports"
	"main/schemas/domain/factories"
	"main/schemas/domain/models"
)

type SchemasService struct{
	SchemasFactory *factories.SchemasFactory
	SchemasRepository ports.SchemasRepositoryPort
}

func (ss *SchemasService) GetById(id string) (*models.Schema, error) {
	return ss.SchemasRepository.GetById(id)
}

func (ss *SchemasService) Create(createSchemaPayload *payloads.CreateSchemaPayload) (*models.Schema, error) {
	schema := ss.SchemasFactory.Create(createSchemaPayload.Id)

	return ss.SchemasRepository.Create(schema)
}