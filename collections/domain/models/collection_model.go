package models

import documentModels "main/documents/domain/models"

type Collection struct {
	Name       string               `json:"name"`
	Properties map[string]*Property `json:"properties"`
	Documents  map[string]*documentModels.Document `json:"documents"`
}