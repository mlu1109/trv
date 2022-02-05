package trv_test

import (
	"encoding/json"
	"testing"
	"trv"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalResponse(t *testing.T) {
	// Given
	j := trainAnnouncementResponseJSON
	// When
	var response trv.Response
	err := json.Unmarshal([]byte(j), &response)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, 1, len(response.Response.Result))
	taInterface := response.Response.Result[0].TrainAnnouncement
	taJSON, err := json.Marshal(taInterface)
	assert.Nil(t, err)
	var ta []trv.TrainAnnouncementV16
	err = json.Unmarshal(taJSON, &ta)
	assert.Nil(t, err)
	assert.Equal(t, 20, len(ta))
	first := ta[0]
	assert.Equal(t, trv.Departure, first.ActivityType)
	assert.Equal(t, "13", first.TrackAtLocation)
	assert.Equal(t, "Hpbg", first.ToLocation[0].LocationName)
	last := ta[19]
	assert.Equal(t, trv.Arrival, last.ActivityType)
	assert.Equal(t, "4", last.TrackAtLocation)
	assert.Equal(t, "Ös", last.ToLocation[0].LocationName)
}
