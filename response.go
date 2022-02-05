package trv

type Response struct {
	Result []ResultItem `json:"RESULT"`
}

type ResultItem struct {
	Info  Info  `json:"INFO,omitempty"`
	Error Error `json:"ERROR,omitempty"`
}

type Error struct {
	Source  string `json:"SOURCE,omitempty"`
	Message string `json:"MESSAGE,omitempty"`
}

type Info struct {
}
