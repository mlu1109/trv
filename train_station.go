package trv

import "encoding/json"

type TrainStationV14 struct {
	Advertised                  bool
	AdvertisedLocationName      string
	AdvertisedShortLocationName string
	CountryCode                 string
	CountyNo                    []int
	Deleted                     bool
	Geometry                    Geometry
	LocationInformationText     string
	LocationSignature           string
	ModifiedTime                string
	OfficialLocationName        string
	PlatformLine                []string
	Prognosticated              bool
}

type Geometry struct {
	SWEREF99TM string
	WGS84      string
}

func ReadTrainStationV14(response *Response) ([]TrainStationV14, error) {
	var res []TrainStationV14
	items := response.Response.Result
	for _, item := range items {
		j, err := json.Marshal(item.TrainStation)
		if err != nil {
			return res, err
		}
		var tss []TrainStationV14
		err = json.Unmarshal(j, &tss)
		if err != nil {
			return res, err
		}
		res = append(res, tss...)
	}
	return res, nil
}
