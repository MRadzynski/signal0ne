package opensearch

import (
	"signal0ne/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_GetLogOccurrences tests the function getLogOccurrences

// You need to run lab01 -  https://portal.azure.com/#@szymonst2808gmail.onmicrosoft.com/resource/subscriptions/fb775820-301c-4a7d-af99-83285b864825/resourceGroups/rg01-lab01/providers/Microsoft.Compute/virtualMachines/lab01/overview
// In order to run this test without mocking
func Test_GetLogOccurrences(t *testing.T) {
	mockConn := utils.ConnectToSocket()
	defer mockConn.Close()
	mockedGetLogOccurrencesInput := map[string]string{
		"service":    "adservice",
		"query":      "{\"query\": {\"match_all\": {}}}",
		"compare_by": "resource.service.name",
	}

	//Create integration object
	inventory := NewOpenSearchIntegrationInventory(mockConn)
	integration := OpenSearchIntegration{
		Inventory: inventory,
		Config: Config{
			Host:  "20.127.192.216",
			Index: "otel",
			Port:  "9200",
			Ssl:   false,
		},
	}

	output, err := getLogOccurrences(mockedGetLogOccurrencesInput, integration)
	assert.NoError(t, err)
	assert.NotNil(t, output)

}
