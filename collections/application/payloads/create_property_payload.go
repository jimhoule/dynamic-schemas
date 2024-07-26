package payloads

type CreatePropertyPayload struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	IsRequired bool   `json:"isRequired"`
}