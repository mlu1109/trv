package trv_test

import (
	"encoding/json"
	"testing"

	"github.com/mlu1109/trv"
	"github.com/stretchr/testify/assert"
)

func TestReadTrainStationV14(t *testing.T) {
	// Given
	resJSON := trainStationResponseJSON
	var response trv.Response
	err := json.Unmarshal([]byte(resJSON), &response)
	assert.Nil(t, err)
	// When
	tss, err := trv.ReadTrainStationV14(&response)
	// Then
	assert.Nil(t, err)
	assert.NotNil(t, tss)
}
