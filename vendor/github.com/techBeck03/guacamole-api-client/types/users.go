package types

// GuacUser base user type
type GuacUser struct {
	Username   string             `json:"username"`
	Password   string             `json:"password,omitempty"`
	Attributes GuacUserAttributes `json:"attributes,omitempty"`
	LastActive int                `json:"lastActive,omitempty"`
}

// GuacUserAttributes additional user attributes
type GuacUserAttributes struct {
	GuacOrganizationalRole string `json:"guac-organizational-role,omitempty"`
	GuacFullName           string `json:"guac-full-name,omitempty"`
	Email                  string `json:"guac-email-address,omitempty"`
	Expired                string `json:"expired,omitempty"`
	Timezone               string `json:"timezone,omitempty"`
	AccessWindowStart      string `json:"access-window-start,omitempty"`
	AccessWindowEnd        string `json:"access-window-end,omitempty"`
	Disabled               string `json:"disabled,omitempty"`
	ValidFrom              string `json:"valid-from,omitempty"`
	ValidUntil             string `json:"valid-until,omitempty"`
}
