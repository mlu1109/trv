package trv_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/mlu1109/trv"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	t.Run("TestMarshalRequest", func(t *testing.T) {
		// Given
		authenticationKey := "yourauthenticationkey"
		sseURL := true
		objectType := "TrainAnnouncement"
		orderBy := "AdvertisedTimeAtLocation"
		schemaVersion := "1.3"
		filterEQName := "AdvertisedTrainIdent"
		filterEQValue := "129"
		filterGTEName := "ScheduledDepartureDateTime"
		filterGTEValue := "2019-05-07T00:00:00.000+02:00"
		filterLTName := "ScheduledDepartureDateTime"
		filterLTValue := "2019-05-08T00:00:00.000+02:00"
		filterExistsName := "TimeAtLocation"
		filterExistsValue := true
		include := []string{"LocationSignature", "ActivityType", "AdvertisedTimeAtLocation", "TimeAtLocation", "ScheduledDepartureDateTime"}
		request := trv.NewRequest(authenticationKey)
		query := request.
			NewQuery(objectType, schemaVersion).
			WithSSEURL(sseURL).
			WithOrderBy(orderBy)
		for _, value := range include {
			query.AddInclude(value)
		}
		query.
			NewFilter().
			NewAnd().
			AddEQ(filterEQName, filterEQValue).
			AddGTE(filterGTEName, filterGTEValue).
			AddLT(filterLTName, filterLTValue).
			AddExists(filterExistsName, filterExistsValue)
		// When
		marshalled, err := xml.MarshalIndent(request, "", "\t")
		actual := string(marshalled)
		// Then
		expected := fmt.Sprintf(`<REQUEST>
	<LOGIN authenticationkey="%s"></LOGIN>
	<QUERY sseurl="%t" objecttype="%s" orderby="%s" schemaversion="%s">
		<FILTER>
			<AND>
				<EQ name="%s" value="%s"></EQ>
				<GTE name="%s" value="%s"></GTE>
				<LT name="%s" value="%s"></LT>
				<EXISTS name="%s" value="%t"></EXISTS>
			</AND>
		</FILTER>
		<INCLUDE>%s</INCLUDE>
		<INCLUDE>%s</INCLUDE>
		<INCLUDE>%s</INCLUDE>
		<INCLUDE>%s</INCLUDE>
		<INCLUDE>%s</INCLUDE>
	</QUERY>
</REQUEST>`, authenticationKey, sseURL, objectType, orderBy, schemaVersion, filterEQName, filterEQValue, filterGTEName, filterGTEValue, filterLTName, filterLTValue, filterExistsName, filterExistsValue, include[0], include[1], include[2], include[3], include[4])
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
