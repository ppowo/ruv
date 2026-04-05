package models

// Station represents a radio station configuration
type Station struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
