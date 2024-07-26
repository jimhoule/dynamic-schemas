package services

import (
	"main/tenants/application/payloads"
	"main/tenants/application/ports"
	"main/tenants/domain/factories"
	"main/tenants/domain/models"
)

type TenantsService struct {
	TenantsFactory *factories.TenantsFactory
	TenantsRepository ports.TenantsRepositoryPort
}

func (ts *TenantsService) Create(createTenantPayload *payloads.CreateTenantPayload) (*models.Tenant, error) {
	tenant := ts.TenantsFactory.Create(createTenantPayload.Name)

	return ts.TenantsRepository.Create(tenant)
}

func (ts *TenantsService) GetById(id string) (*models.Tenant, error) {
	return ts.TenantsRepository.GetById(id)
}