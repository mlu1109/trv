package trv_test

import (
	"encoding/json"
	"testing"

	"github.com/mlu1109/trv"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalResponseWithOneItem(t *testing.T) {
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

func TestUnmarshalResponseWithTwoItems(t *testing.T) {
	unpack := func(r trv.ResultItem) []string {
		tas := r.TrainAnnouncement
		assert.Equal(t, 1, len(tas))
		ta := tas[0].(map[string]interface{})
		assert.Equal(t, 1, len(ta))
		value, ok := ta["ScheduledDepartureDateTime"]
		assert.True(t, ok)
		items := value.([]interface{})
		var sddts []string
		for _, item := range items {
			sddt := item.(string)
			sddts = append(sddts, sddt)
		}
		return sddts
	}
	// Given
	j := trainAnnouncementResponseWithSeveralResultsJSON
	// When
	var response trv.Response
	err := json.Unmarshal([]byte(j), &response)
	// Then
	assert.Nil(t, err)
	assert.Equal(t, 2, len(response.Response.Result))
	expected := [][]string{
		{
			"2022-02-10T00:00:00.000+01:00",
			"2022-02-11T00:00:00.000+01:00",
			"2022-02-12T00:00:00.000+01:00",
			"2022-02-13T00:00:00.000+01:00",
			"2022-02-14T00:00:00.000+01:00",
		},
		{
			"2022-01-30T00:00:00.000+01:00",
			"2022-02-04T00:00:00.000+01:00",
			"2022-02-06T00:00:00.000+01:00",
			"2022-02-07T00:00:00.000+01:00",
			"2022-02-08T00:00:00.000+01:00",
			"2022-02-09T00:00:00.000+01:00",
		},
	}
	for i, resultItem := range response.Response.Result {
		actual := unpack(resultItem)
		assert.Equal(t, expected[i], actual)
	}
}
