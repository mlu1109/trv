package trv

import (
	"encoding/xml"
)

type login struct {
	XMLName           xml.Name `xml:"LOGIN"`
	AuthenticationKey string   `xml:"authenticationkey,attr"`
}

type Request struct {
	XMLName xml.Name `xml:"REQUEST"`
	Login   *login
	Query   *query
}

func NewRequest(authenticationKey string) *Request {
	r := &Request{}
	r.Login = &login{AuthenticationKey: authenticationKey}
	return r
}

func (r *Request) NewQuery(objectType, schemaVersion string) *query {
	r.Query = &query{ObjectType: objectType, SchemaVersion: schemaVersion}
	return r.Query
}
