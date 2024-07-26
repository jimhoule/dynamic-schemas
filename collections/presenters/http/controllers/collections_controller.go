package controllers

import (
	"main/collections/application/payloads"
	"main/collections/application/services"
	"main/collections/presenters/http/dtos"
	"main/router"
	"main/router/utils"
	"net/http"
)

type CollectionsController struct {
	CollectionsService *services.CollectionsService
}

func (cc *CollectionsController) GetAll(writer http.ResponseWriter, request *http.Request) {
	schemaId := router.GetUrlParam(request, "schemaId")

	collection, err := cc.CollectionsService.GetAll(schemaId)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, collection)
}

func (cc *CollectionsController) GetByName(writer http.ResponseWriter, request *http.Request) {
	schemaId := router.GetUrlParam(request, "schemaId")
	name := router.GetUrlParam(request, "name")

	collection, err := cc.CollectionsService.GetByName(schemaId, name)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, collection)
}

func (cc *CollectionsController) Create(writer http.ResponseWriter, request *http.Request) {
	var createCollectionDto dtos.CreateCollectionDto
	err := utils.ReadHttpRequestBody(writer, request, &createCollectionDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	collection, err := cc.CollectionsService.Create(&payloads.CreateCollectionPayload{
		Name: createCollectionDto.Name,
		SchemaId: createCollectionDto.SchemaId,
		CreatePropertyPayloads: createCollectionDto.CreatePropertyPayloads,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, collection)
}