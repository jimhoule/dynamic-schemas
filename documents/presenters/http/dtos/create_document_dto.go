package dtos

type CreateDocumentDto struct {
	SchemaName     string         `json:"schemaName"`
	CollectionName string         `json:"collectionName"`
	Body           map[string]any `json:"body"`
}
