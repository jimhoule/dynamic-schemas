package factories

import "main/collections/domain/models"

type PropertiesFactory struct{}

func (*PropertiesFactory) Create(isRequired bool, propertyType string) *models.Property {
	return &models.Property{
		IsRequired: isRequired,
		Type: propertyType,
	}
}