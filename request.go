package trafikverket

import (
	"encoding/xml"
)

type login struct {
	XMLName           xml.Name `xml:"LOGIN"`
	AuthenticationKey string   `xml:"authenticationkey,attr"`
}

type request struct {
	XMLName xml.Name `xml:"REQUEST"`
	Login   login
	Query   query
}

func newRequest(authenticationKey string) *request {
	return &request{
		Login: login{AuthenticationKey: authenticationKey},
	}
}
