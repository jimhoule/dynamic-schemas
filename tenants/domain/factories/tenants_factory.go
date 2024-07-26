package factories

import (
	"main/tenants/domain/models"
	"main/uuid/services"
)

type TenantsFactory struct{
	UuidService services.UuidService
}

func (tf *TenantsFactory) Create(name string) *models.Tenant {
	return &models.Tenant{
		Id: tf.UuidService.Generate(),
		Name: name,
	}
}
