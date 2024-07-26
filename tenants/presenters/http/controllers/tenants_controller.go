package controllers

import (
	"main/router"
	"main/router/utils"
	"main/tenants/application/payloads"
	"main/tenants/application/services"
	"main/tenants/presenters/http/dtos"
	"net/http"
)

type TenantsController struct {
	TenantsService *services.TenantsService
}

func (ts *TenantsController) GetById(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	tenant, err := ts.TenantsService.GetById(id)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, tenant)
}

func (ts *TenantsController) Create(writer http.ResponseWriter, request *http.Request) {
	var createTenantDto dtos.CreateTenantDto
	err := utils.ReadHttpRequestBody(writer, request, &createTenantDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusBadRequest, err)
		return
	}

	tenant, err := ts.TenantsService.Create(&payloads.CreateTenantPayload{
		Name: createTenantDto.Name,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, tenant)
}