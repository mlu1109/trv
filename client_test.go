package trv_test

import (
	"os"
	"testing"

	"github.com/mlu1109/trv"
	"github.com/stretchr/testify/assert"
)

func TestClientIntegrationTests(t *testing.T) {
	authenticationKey := os.Getenv("TRV_API_KEY")
	if authenticationKey == "" {
		t.Skip("Client tests are skipped as TRV_API_KEY is not set")
	}
	t.Run("Get TrainAnnouncements (JSON)", func(t *testing.T) {
		// Given
		request := trv.
			NewRequest(authenticationKey)
		request.
			NewQuery("TrainAnnouncement", "1.6").
			WithLimit(20).
			NewFilter().
			AddEQ("LocationSignature", "Cst")
		// When
		response, err := trv.
			NewClient().
			Post(request)
		// Then ...
		// ... response is OK
		assert.Nil(t, err)
		// ... items are returned according to Query
		items, err := trv.ReadTrainAnnouncementV16(response)
		t.Logf("Fetched %d items from Trafikverket", len(items))
		assert.Nil(t, err)
		assert.Equal(t, 20, len(items))
		for _, item := range items {
			assert.Equal(t, "Cst", item.LocationSignature)
		}
	})
}
