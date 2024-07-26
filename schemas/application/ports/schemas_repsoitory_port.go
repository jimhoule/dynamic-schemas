package ports

import "main/schemas/domain/models"

type SchemasRepositoryPort interface {
	GetById(id string) (*models.Schema, error)
	Create(schema *models.Schema) (*models.Schema, error)
}