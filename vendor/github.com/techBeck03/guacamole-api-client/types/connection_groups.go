package types

// GuacConnectionGroup defines base values of a connection group
type GuacConnectionGroup struct {
	Name              string                        `json:"name"`
	Identifier        string                        `json:"identifier"`
	ParentIdentifier  string                        `json:"parentIdentifier"`
	Type              string                        `json:"type"`
	ActiveConnections int                           `json:"activeConnections"`
	ChildConnections  []GuacConnection              `json:"childConnections"`
	ChildGroups       []GuacConnectionGroup         `json:"childConnectionGroups"`
	Attributes        GuacConnectionGroupAttributes `json:"attributes"`
}

// GuacConnectionGroupAttributes defines attributes of a connection group
type GuacConnectionGroupAttributes struct {
	MaxConnections        string `json:"max-connections"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
	EnableSessionAffinity string `json:"enable-session-affinity"`
}

// ValidTypes returns list of valid types
func (GuacConnectionGroup) ValidTypes() []string {
	return []string{
		"ORGANIZATIONAL",
		"BALANCING",
	}
}
