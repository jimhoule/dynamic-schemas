package controllers

import (
	"main/documents/application/payloads"
	"main/documents/application/services"
	"main/documents/presenters/http/dtos"
	"main/router"
	"main/router/utils"
	"net/http"
)

type DocumentsController struct {
	DocumentsService *services.DocumentsService
}

func (dc *DocumentsController) GetAll(writer http.ResponseWriter, request *http.Request) {
	schemaId := router.GetUrlParam(request, "schemaId")
	collectionName := router.GetUrlParam(request, "collectionName")

	documents, err := dc.DocumentsService.GetAll(schemaId, collectionName)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, documents)
}

func (dc *DocumentsController) GetByKey(writer http.ResponseWriter, request *http.Request) {
	schemaId := router.GetUrlParam(request, "schemaId")
	collectionName := router.GetUrlParam(request, "collectionName")
	key := router.GetUrlParam(request, "key")

	document, err := dc.DocumentsService.GetByKey(schemaId, collectionName, key)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, document)
}

func (dc *DocumentsController) Create(writer http.ResponseWriter, request *http.Request) {
	var createDocumentDto dtos.CreateDocumentDto
	err := utils.ReadHttpRequestBody(writer, request, &createDocumentDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	document, err := dc.DocumentsService.Create(&payloads.CreateDocumentPayload{
		SchemaId: createDocumentDto.SchemaId,
		CollectionName: createDocumentDto.CollectionName,
		Body: createDocumentDto.Body,
	})
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, document)
}