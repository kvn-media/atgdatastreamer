package managers

import (
	"fmt"
	"net"

	"github.com/kvn-media/atgdatastreamer/configs"
)

type ATGManager struct {
	Conn net.Conn // Exported field
}

func NewATGManager(config *configs.Config) (*ATGManager, error) {
	// Open connection to ATG
	conn, err := net.Dial("tcp", config.ATGAddress+":"+fmt.Sprintf("%d", config.ATGPort))
	if err != nil {
		return nil, err
	}

	manager := &ATGManager{
		Conn: conn, // Assign the connection to the ATGManager
	}

	return manager, nil
}

func (manager *ATGManager) Close() {
	// Close connection to ATG
	manager.Conn.Close()
}
