package trv

type ActivityType string

const (
	Avgang  ActivityType = "Avgang"
	Ankomst ActivityType = "Ankomst"
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
	EstimatedTimeIsPreliminary            string
	FromLocation                          []LocationInformation
	InformationOwner                      string
	LocationSignature                     string
	MobileWebLink                         string
	ModifiedTime                          string
	NewEquipment                          string
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
