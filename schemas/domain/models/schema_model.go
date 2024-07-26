package models

import "main/collections/domain/models"

type Schema struct {
	Id          string                        `json:"id"`
	Collections map[string]*models.Collection `json:"collections"`
}