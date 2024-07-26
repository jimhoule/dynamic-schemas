package payloads

type CreateCollectionPayload struct {
	Name                   string                   `json:"name"`
	SchemaId               string                   `json:"schemaId"`
	CreatePropertyPayloads []*CreatePropertyPayload `json:"createPropertyPayloads"`
}