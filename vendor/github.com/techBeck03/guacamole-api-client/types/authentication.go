package types

// AuthenticationResponse defines connection response
type AuthenticationResponse struct {
	AuthToken            string   `json:"authToken"`
	Username             string   `json:"username"`
	DataSource           string   `json:"dataSource"`
	AvailableDataSources []string `json:"availableDataSources"`
}
