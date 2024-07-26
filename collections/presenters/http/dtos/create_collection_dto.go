package dtos

import "main/collections/application/payloads"

type CreateCollectionDto struct {
	Name                   string                            `json:"name"`
	SchemaId               string                            `json:"schemaId"`
	CreatePropertyPayloads []*payloads.CreatePropertyPayload `json:"createPropertyPayloads"`
}