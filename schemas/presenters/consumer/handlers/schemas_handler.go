package handlers

import (
	"encoding/json"
	"main/schemas/application/payloads"
	"main/schemas/application/services"
	"main/schemas/presenters/consumer/dtos"
)

type SchemasHandler struct {
	SchemasService *services.SchemasService
}

func (sh *SchemasHandler) Create(body []byte) error {
	var createSchemaDto dtos.CreateSchemaDto
	err := json.Unmarshal(body, &createSchemaDto)
	if err != nil {
		return err
	}

	_, err = sh.SchemasService.Create(&payloads.CreateSchemaPayload{
		Name: createSchemaDto.Name,
	})

	if err != nil {
		return err
	}

	return nil
}
