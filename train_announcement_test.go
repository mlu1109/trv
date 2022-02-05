package trv_test

import (
	"encoding/json"
	"testing"
	"trv"

	"github.com/stretchr/testify/assert"
)

func TestReadTrainAnnouncementsV16(t *testing.T) {
	// Given
	resJSON := trainAnnouncementResponseJSON
	var response trv.Response
	err := json.Unmarshal([]byte(resJSON), &response)
	assert.Nil(t, err)
	// When
	tas, err := trv.ReadTrainAnnouncementV16(&response)
	// Then
	assert.Nil(t, err)
	assert.NotNil(t, tas)
}
