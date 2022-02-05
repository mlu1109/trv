package trv

import "encoding/xml"

type query struct {
	XMLName               xml.Name                  `xml:"QUERY"`
	SSEURL                bool                      `xml:"sseurl,attr,omitempty"`
	ObjectType            string                    `xml:"objecttype,attr"`
	OrderBy               string                    `xml:"orderby,attr,omitempty"`
	SchemaVersion         string                    `xml:"schemaversion,attr"`
	ChangeID              int                       `xml:"changeid,attr,omitempty"`
	ID                    string                    `xml:"id,attr,omitempty"`
	IncludeDeletedObjects bool                      `xml:"includedeletedobjects,attr,omitempty"`
	LastModified          bool                      `xml:"lastmodified,attr,omitempty"`
	Limit                 int                       `xml:"limit,attr,omitempty"`
	Skip                  int                       `xml:"skip,attr,omitempty"`
	Filter                struct{ Filters filters } `xml:"FILTER"`
	Include               []string                  `xml:"INCLUDE,omitempty"`
}

func (q *query) WithID(id string) *query {
	q.ID = id
	return q
}

func (q *query) WithLimit(limit int) *query {
	q.Limit = limit
	return q
}

func (q *query) WithIncludDeletedObjects(includeDeletedObjects bool) *query {
	q.IncludeDeletedObjects = includeDeletedObjects
	return q
}

func (q *query) WithOrderBy(orderBy string) *query {
	q.OrderBy = orderBy
	return q
}

func (q *query) WithSkip(skip int) *query {
	q.Skip = skip
	return q
}

func (q *query) WithLastModified(lastModified bool) *query {
	q.LastModified = lastModified
	return q
}

func (q *query) WithChangeID(changeID int) *query {
	q.ChangeID = changeID
	return q
}

func (q *query) WithSSEURL(sseURL bool) *query {
	q.SSEURL = sseURL
	return q
}

func (q *query) AddInclude(field string) *query {
	q.Include = append(q.Include, field)
	return q
}

func (q *query) NewFilter() *filters {
	q.Filter = struct{ Filters filters }{}
	return &q.Filter.Filters
}
