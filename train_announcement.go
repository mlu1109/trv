package trv

import "encoding/json"

type ActivityType string

const (
	Departure ActivityType = "Avgang"
	Arrival   ActivityType = "Ankomst"
)

type TrainAnnouncementV16 struct {
	ActivityId                            string
	ActivityType                          ActivityType
	Advertised                            bool
	AdvertisedTimeAtLocation              string
	AdvertisedTrainIdent                  string
	Booking                               []TrainInformation
	Canceled                              bool
	Deleted                               bool
	Deviation                             []TrainInformation
	EstimatedTimeAtLocation               string
	EstimatedTimeIsPreliminary            bool
	FromLocation                          []LocationInformation
	InformationOwner                      string
	LocationSignature                     string
	MobileWebLink                         string
	ModifiedTime                          string
	NewEquipment                          int
	Operator                              string
	OtherInformation                      []TrainInformation
	PlannedEstimatedTimeAtLocation        string
	PlannedEstimatedTimeAtLocationIsValid bool
	ProductInformation                    []TrainInformation
	ScheduledDepartureDateTime            string
	Service                               []TrainInformation
	TechnicalDateTime                     string
	TechnicalTrainIdent                   string
	TimeAtLocation                        string
	TimeAtLocationWithSeconds             string
	ToLocation                            []LocationInformation
	TrackAtLocation                       string
	TrainComposition                      []TrainInformation
	TrainOwner                            string
	TypeOfTraffic                         []TrainInformation
	WebLink                               string
	WebLinkName                           string
	ViaFromLocation                       []LocationInformation
	ViaToLocation                         []LocationInformation
}

type TrainInformation struct {
	Code        string
	Description string
}

type LocationInformation struct {
	LocationName string
	Order        int
	Priority     int
}

func ReadTrainAnnouncementV16(response *Response) ([]TrainAnnouncementV16, error) {
	var res []TrainAnnouncementV16
	items := response.Response.Result
	for _, item := range items {
		j, err := json.Marshal(item.TrainAnnouncement)
		if err != nil {
			return res, err
		}
		var tas []TrainAnnouncementV16
		err = json.Unmarshal(j, &tas)
		if err != nil {
			return res, err
		}
		res = append(res, tas...)
	}
	return res, nil
}
