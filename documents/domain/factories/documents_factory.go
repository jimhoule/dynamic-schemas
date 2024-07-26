package factories

import (
	"main/documents/domain/models"
	"main/uuid/services"
)

type DocumentsFactory struct{
	UuidService services.UuidService
}

func (df *DocumentsFactory) Create(body map[string]any) *models.Document {
	return &models.Document{
		Key: df.UuidService.Generate(),
		Body: body,
	}
}