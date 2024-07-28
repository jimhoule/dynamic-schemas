package mappers

import (
	"bytes"
	"encoding/json"
	"main/collections/domain/models"
	"main/collections/infrastructure/persistence/arangodb/entities"
)

type PropertiesMapper struct{}

func (pm *PropertiesMapper) ToDomainModels(propertiesEntity entities.Properties) []*models.Property {
	properties := []*models.Property{}
	for key, arangoProperty := range propertiesEntity.Rule.Properties {
		property := &models.Property{
			Name:       key,
			Type:       arangoProperty.Type,
			IsRequired: false,
		}

		// If is of type array
		if arangoProperty.Type == "array" {
			property.ItemType = arangoProperty.Items.Type
		}

		// If is required
		for _, requiredPropertyName := range propertiesEntity.Rule.Required {
			if requiredPropertyName == key {
				property.IsRequired = true
			}
		}

		properties = append(properties, property)
	}

	return properties
}

func (pm *PropertiesMapper) ToEntity(propertyModels []*models.Property) entities.Properties {
	propertyEntities := entities.Properties{
		Rule: entities.Rule{
			Properties: map[string]entities.Property{},
			Required:   []string{},
		},
		Level: "strict",
	}

	for _, propertyModel := range propertyModels {
		// If is required
		if propertyModel.IsRequired {
			propertyEntities.Rule.Required = append(propertyEntities.Rule.Required, propertyModel.Name)
		}

		propertyEntity := entities.Property{
			Type: propertyModel.Type,
		}
		// If is of type array, adds type of each item in the array
		if propertyModel.Type == "array" {
			propertyEntity.Items.Type = propertyModel.ItemType
		}

		propertyEntities.Rule.Properties[propertyModel.Name] = propertyEntity
	}

	return propertyEntities
}

func (pm *PropertiesMapper) ToParsedEntity(in any, parsedPropertiesEntity *entities.Properties) {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(in)
	json.NewDecoder(&buffer).Decode(parsedPropertiesEntity)
}
