package services

import (
	// "github.com/kvn-media/atgdatastreamer/models"
	"github.com/kvn-media/atgdatastreamer/repositories"
)

// ATGService handles business logic related to the ATG system
type ATGService struct {
	Repository *repositories.ATGRepository
}

// NewATGService creates a new instance of ATGService
func NewATGService(repo *repositories.ATGRepository) *ATGService {
	return &ATGService{
		Repository: repo,
	}
}

// SaveATGData saves ATG data to the database
func (s *ATGService) SaveATGData(data interface{}) error {
	// Implement any business logic or validation before saving data
	// Example: Check if the data is valid before saving

	// Call the repository to save the data
	return s.Repository.SaveATGData(data)
}

// GetLatestATGData retrieves the latest ATG data from the database
func (s *ATGService) GetLatestATGData() (interface{}, error) {
	// Call the repository to get the latest ATG data
	return s.Repository.GetLatestATGData()
}