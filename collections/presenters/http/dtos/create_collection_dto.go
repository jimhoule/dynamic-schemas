package dtos

import "main/collections/application/payloads"

type CreateCollectionDto struct {
	Name                   string                            `json:"name"`
	SchemaName             string                            `json:"schemaName"`
	CreatePropertyPayloads []*payloads.CreatePropertyPayload `json:"createPropertyPayloads"`
}
