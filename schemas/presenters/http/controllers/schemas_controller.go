package controllers

import (
	"main/router"
	"main/router/utils"
	"main/schemas/application/payloads"
	"main/schemas/application/services"
	"main/schemas/presenters/http/dtos"
	"net/http"
)

type SchemasController struct {
	SchemasService *services.SchemasService
}

func (sc *SchemasController) GetById(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")

	schema, err := sc.SchemasService.GetById(id)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, schema)
}

func (sc *SchemasController) Create(writer http.ResponseWriter, request *http.Request) {
	var createSchemaDto dtos.CreateSchemaDto
	err := utils.ReadHttpRequestBody(writer, request, &createSchemaDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	schema, err := sc.SchemasService.Create(&payloads.CreateSchemaPayload{
		Id: createSchemaDto.Id,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, schema)
}