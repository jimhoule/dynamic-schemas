package services

import (
	"main/collections/application/payloads"
	"main/collections/application/ports"
	"main/collections/domain/factories"
	"main/collections/domain/models"
)

type CollectionsService struct {
	CollectionsFactory    *factories.CollectionsFactory
	PropertiesFactory     *factories.PropertiesFactory
	CollectionsRepository ports.CollectionsRepositoryPort
}

func (cs *CollectionsService) GetAll(schemaName string) ([]*models.Collection, error) {
	return cs.CollectionsRepository.GetAll(schemaName)
}

func (cs *CollectionsService) GetByName(schemaName string, name string) (*models.Collection, error) {
	return cs.CollectionsRepository.GetByName(schemaName, name)
}

func (cs *CollectionsService) Create(createCollectionPayload *payloads.CreateCollectionPayload) (*models.Collection, error) {
	properties := []*models.Property{}
	for _, createPropertyPayload := range createCollectionPayload.CreatePropertyPayloads {
		property := cs.PropertiesFactory.Create(
			createPropertyPayload.Name,
			createPropertyPayload.IsRequired,
			createPropertyPayload.Type,
		)

		properties = append(properties, property)
	}

	collection := cs.CollectionsFactory.Create(createCollectionPayload.Name, properties)

	return cs.CollectionsRepository.Create(createCollectionPayload.SchemaName, collection)
}
