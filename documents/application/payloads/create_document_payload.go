package payloads

type CreateDocumentPayload struct {
	SchemaName     string         `json:"schemaName"`
	CollectionName string         `json:"collectionName"`
	Body           map[string]any `json:"body"`
}
