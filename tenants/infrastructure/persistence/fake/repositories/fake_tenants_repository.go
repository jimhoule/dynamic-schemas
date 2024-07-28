package repositories

import "main/tenants/domain/models"

type FakeTenantsRepository struct {
	Tenants []*models.Tenant
}

func (ftr *FakeTenantsRepository) Reset() {
	ftr.Tenants = []*models.Tenant{}
}

func (ftr *FakeTenantsRepository) GetById(id string) (*models.Tenant, error) {
	for _, tenant := range ftr.Tenants {
		if tenant.Id == id {
			return tenant, nil
		}
	}

	return nil, nil
}

func (ftr *FakeTenantsRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {
	ftr.Tenants = append(ftr.Tenants, tenant)

	return tenant, nil
}
