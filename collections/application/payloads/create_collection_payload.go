package payloads

type CreateCollectionPayload struct {
	Name                   string                   `json:"name"`
	SchemaName             string                   `json:"schemaName"`
	CreatePropertyPayloads []*CreatePropertyPayload `json:"createPropertyPayloads"`
}
