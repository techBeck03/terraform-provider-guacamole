package types

// GuacPermissionItem defines a basic permissions operation item
type GuacPermissionItem struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

// GuacPermissionData defines connection permission data
type GuacPermissionData struct {
	ConnectionPermissions       map[string][]string `json:"connectionPermissions"`
	ConnectionGroupPermissions  map[string][]string `json:"connectionGroupPermissions"`
	SharingProfilePermissions   map[string][]string `json:"sharingProfilePermissions"`
	UserPermissions             map[string][]string `json:"userPermissions"`
	UserGroupPermissions        map[string][]string `json:"userGroupPermissions"`
	SystemPermissions           []string            `json:"systemPermissions"`
	ActiveConnectionPermissions map[string][]string `json:"activeConnectionPermissions"`
}

// SystemPermissions defines a class type for system permissions
type SystemPermissions struct {
}

// Administer constant value option
func (SystemPermissions) Administer() string {
	return "ADMINISTER"
}

// CreateUser constant value option
func (SystemPermissions) CreateUser() string {
	return "CREATE_USER"
}

// CreateUserGroup constant value option
func (SystemPermissions) CreateUserGroup() string {
	return "CREATE_USER_GROUP"
}

// CreateConnection constant value option
func (SystemPermissions) CreateConnection() string {
	return "CREATE_CONNECTION"
}

// CreateConnectionGroup constant value option
func (SystemPermissions) CreateConnectionGroup() string {
	return "CREATE_CONNECTION_GROUP"
}

// CreateSharingProfile constant value option
func (SystemPermissions) CreateSharingProfile() string {
	return "CREATE_SHARING_PROFILE"
}

// ValidChoices returns an array of valid system permission choices
func (SystemPermissions) ValidChoices() []string {
	return []string{
		"ADMINISTER",
		"CREATE_USER",
		"CREATE_USER_GROUP",
		"CREATE_CONNECTION",
		"CREATE_CONNECTION_GROUP",
		"CREATE_SHARING_PROFILE",
	}
}
