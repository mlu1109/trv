package trv

import "encoding/xml"

type Query struct {
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
	Filter                struct{ Filters Filters } `xml:"FILTER"`
	Include               []string                  `xml:"INCLUDE,omitempty"`
	Distinct              []string                  `xml:"DISTINCT,omitempty"`
}

func (q *Query) WithID(id string) *Query {
	q.ID = id
	return q
}

func (q *Query) WithLimit(limit int) *Query {
	q.Limit = limit
	return q
}

func (q *Query) WithIncludDeletedObjects(includeDeletedObjects bool) *Query {
	q.IncludeDeletedObjects = includeDeletedObjects
	return q
}

func (q *Query) WithOrderBy(orderBy string) *Query {
	q.OrderBy = orderBy
	return q
}

func (q *Query) WithSkip(skip int) *Query {
	q.Skip = skip
	return q
}

func (q *Query) WithLastModified(lastModified bool) *Query {
	q.LastModified = lastModified
	return q
}

func (q *Query) WithChangeID(changeID int) *Query {
	q.ChangeID = changeID
	return q
}

func (q *Query) WithSSEURL(sseURL bool) *Query {
	q.SSEURL = sseURL
	return q
}

func (q *Query) AddInclude(field string) *Query {
	q.Include = append(q.Include, field)
	return q
}

func (q *Query) AddDistinct(field string) *Query {
	q.Distinct = append(q.Distinct, field)
	return q
}

func (q *Query) NewFilter() *Filters {
	q.Filter = struct{ Filters Filters }{}
	return &q.Filter.Filters
}
