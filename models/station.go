package models

import "fmt"

// Station represents a radio station configuration
type Station struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Validate checks if the station configuration is valid
func (s *Station) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("station name cannot be empty")
	}
	if s.URL == "" {
		return fmt.Errorf("station URL cannot be empty")
	}
	return nil
}