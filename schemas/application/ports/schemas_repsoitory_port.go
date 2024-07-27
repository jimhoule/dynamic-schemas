package ports

import "main/schemas/domain/models"

type SchemasRepositoryPort interface {
	GetByName(name string) (*models.Schema, error)
	Create(schema *models.Schema) (*models.Schema, error)
}
