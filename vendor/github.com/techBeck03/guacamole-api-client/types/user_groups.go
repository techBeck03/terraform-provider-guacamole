package types

// GuacUserGroup defines user group properties
type GuacUserGroup struct {
	Identifier string                  `json:"identifier"`
	Attributes GuacUserGroupAttributes `json:"attributes"`
}

// GuacUserGroupAttributes defines user group atribute properties
type GuacUserGroupAttributes struct {
	Disabled string `json:"disabled"`
}
