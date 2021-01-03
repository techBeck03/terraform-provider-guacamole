package types

// ProtocolSchema base schema for protocol choices
type ProtocolSchema struct {
	Name                string               `json:"name"`
	ConnectionForms     []ConnectionForm     `json:"connectionForms"`
	SharingProfileForms []SharingProfileForm `json:"sharingProfileForms"`
}

// ConnectionForm defines the connection shema
type ConnectionForm struct {
	Name   string                `json:"name"`
	Fields []ConnectionFormField `json:"fields"`
}

// ConnectionFormField defines the form field schema
type ConnectionFormField struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

// SharingProfileForm defines the sharing schema
type SharingProfileForm struct {
	Name   string                `json:"name"`
	Fields []ConnectionFormField `json:"fields"`
}
