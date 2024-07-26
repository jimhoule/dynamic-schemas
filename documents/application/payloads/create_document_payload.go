package payloads

type CreateDocumentPayload struct {
	SchemaId       string         `json:"schemaId"`
	CollectionName string         `json:"collectionName"`
	Body           map[string]any `json:"body"`
}