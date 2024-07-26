package factories

import "main/collections/domain/models"

type PropertiesFactory struct{}

func (*PropertiesFactory) Create(name string, isRequired bool, propertyType string) *models.Property {
	return &models.Property{
		Name:       name,
		IsRequired: isRequired,
		Type:       propertyType,
	}
}
