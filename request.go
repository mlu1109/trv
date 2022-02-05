package trv

import (
	"encoding/xml"
)

type Login struct {
	XMLName           xml.Name `xml:"LOGIN"`
	AuthenticationKey string   `xml:"authenticationkey,attr"`
}

type Request struct {
	XMLName xml.Name `xml:"REQUEST"`
	Login   *Login
	Query   *Query
}

func NewRequest(authenticationKey string) *Request {
	r := &Request{}
	r.Login = &Login{AuthenticationKey: authenticationKey}
	return r
}

func (r *Request) NewQuery(objectType, schemaVersion string) *Query {
	r.Query = &Query{ObjectType: objectType, SchemaVersion: schemaVersion}
	return r.Query
}
