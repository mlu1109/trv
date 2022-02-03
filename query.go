package trafikverket

import "encoding/xml"

type query struct {
	XMLName               xml.Name `xml:"QUERY"`
	Objecttype            string   `xml:"objecttype,attr"`
	Schemaversion         string   `xml:"schemaversion,attr"`
	ID                    string   `xml:"id,attr,omitempty"`
	Limit                 int      `xml:"limit,attr,omitempty"`
	IncludeDeletedObjects bool     `xml:"includedeletedobjects,attr,omitempty"`
	OrderBy               string   `xml:"orderby,attr,omitempty"`
	Skip                  int      `xml:"skip,attr,omitempty"`
	LastModified          bool     `xml:"lastmodified,attr,omitempty"`
	ChangeID              int      `xml:"changeid,attr,omitempty"`
	Filter                filter
}
