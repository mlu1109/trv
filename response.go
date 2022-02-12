package trv

type Response struct {
	Response struct {
		Result []ResultItem `json:"RESULT"`
	} `json:"RESPONSE"`
}

type ResultItem struct {
	Info                  *Info         `json:"INFO,omitempty"`
	Error                 *Error        `json:"ERROR,omitempty"`
	RailCrossing          []interface{} `json:",omitempty"`
	ReasonCode            []interface{} `json:",omitempty"`
	TrainAnnouncement     []interface{} `json:",omitempty"`
	TrainMessage          []interface{} `json:",omitempty"`
	TrainStation          []interface{} `json:",omitempty"`
	TrainStationMessage   []interface{} `json:",omitempty"`
	Camera                []interface{} `json:",omitempty"`
	FerryAnnouncement     []interface{} `json:",omitempty"`
	FerryRoute            []interface{} `json:",omitempty"`
	Icon                  []interface{} `json:",omitempty"`
	Parking               []interface{} `json:",omitempty"`
	RoadCondition         []interface{} `json:",omitempty"`
	RoadConditionOverview []interface{} `json:",omitempty"`
	Situation             []interface{} `json:",omitempty"`
	TrafficFlow           []interface{} `json:",omitempty"`
	TrafficSafetyCamera   []interface{} `json:",omitempty"`
	TravelTimeRoute       []interface{} `json:",omitempty"`
	WeatherMeasurepoint   []interface{} `json:",omitempty"`
	WeatherObservation    []interface{} `json:",omitempty"`
	WeatherStation        []interface{} `json:",omitempty"`
	MeasurementData100    []interface{} `json:",omitempty"`
	MeasurementData20     []interface{} `json:",omitempty"`
	PavementData          []interface{} `json:",omitempty"`
	RoadData              []interface{} `json:",omitempty"`
	RoadGeometry          []interface{} `json:",omitempty"`
}

type Error struct {
	Source  string `json:"SOURCE,omitempty" xml:"SOURCE,omitempty"`
	Message string `json:"MESSAGE,omitempty" xml:"MESSAGE,omitempty"`
}

type Info struct {
	LastChangeID string      `json:"LASTCHANGEID,omitempty"`
	EvalResult   interface{} `json:"EVALRESULT,omitempty"`
	SSEURL       string      `json:"SSEURL,omitempty"`

	LastModified struct {
		DateTime string `json:"_attr_datetime"`
	} `json:"LASTMODIFIED,omitempty"`
}
