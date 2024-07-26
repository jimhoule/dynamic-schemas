package repositories

import "main/tenants/domain/models"

type FakeTenantsRepository struct{}

var tenants []*models.Tenant = []*models.Tenant{}

func ResetFakeTenantsRepository() {
	tenants = []*models.Tenant{}
}

func (ftr *FakeTenantsRepository) GetById(id string) (*models.Tenant, error) {
	for _, tenant := range tenants {
		if tenant.Id == id {
			return tenant, nil
		}
	}

	return nil, nil
}

func (ftr *FakeTenantsRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {
	tenants = append(tenants, tenant)

	return tenant, nil
}