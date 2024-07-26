package ports

import "main/tenants/domain/models"

type TenantsRepositoryPort interface {
	Create(tenant *models.Tenant) (*models.Tenant, error)
	GetById(id string) (*models.Tenant, error)
}