package controllers

import (
	"main/router"
	"main/router/utils"
	"main/schemas/application/services"
	"net/http"
)

type SchemasController struct {
	SchemasService *services.SchemasService
}

func (sc *SchemasController) GetByName(writer http.ResponseWriter, request *http.Request) {
	name := router.GetUrlParam(request, "name")

	schema, err := sc.SchemasService.GetByName(name)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, schema)
}
