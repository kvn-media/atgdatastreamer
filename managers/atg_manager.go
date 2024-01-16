package managers

import "github.com/kvn-media/atgdatastreamer/services"

// ATGManager manages the flow of data between the communication layer and the business logic layer
type ATGManager struct {
	ATGService *services.ATGService
}

// NewATGManager creates a new instance of ATGManager
func NewATGManager(atgService *services.ATGService) *ATGManager {
	return &ATGManager{
		ATGService: atgService,
	}
}

// ProcessATGData receives ATG data from the communication layer and processes it
func (m *ATGManager) ProcessATGData(data interface{}) error {
	// Implement any processing logic before passing the data to the service
	// Example: Validate data format, transform data if needed

	// Pass the data to the ATG service to handle business logic and database interaction
	return m.ATGService.SaveATGData(data)
}

// GetLatestATGData retrieves the latest ATG data from the service
func (m *ATGManager) GetLatestATGData() (interface{}, error) {
	// Call the ATG service to get the latest ATG data
	return m.ATGService.GetLatestATGData()
}
