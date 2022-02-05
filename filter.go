package trv

import (
	"encoding/xml"
	"strconv"
)

// filter can be an operator (EQ, GT, ...) or a clause (AND, OR, ...).
// Should be constructed with addOp or newOp
type filter struct {
	XMLName xml.Name
	Name    string `xml:"name,omitempty,attr"`
	Value   string `xml:"value,omitempty,attr"`
	Shape   string `xml:"shape,omitempty"`
	Radius  string `xml:"radius,omitempty"`
	Filters *filters
}

type filters []*filter

// Operators (EQ, GT, ...)

func (f *filters) addOp(tag, name, value string) *filters {
	*f = append(*f, &filter{
		XMLName: xml.Name{Local: tag},
		Name:    name,
		Value:   value,
	})
	return f
}

func (f *filters) AddEQ(name, value string) *filters {
	return f.addOp("EQ", name, value)
}

func (f *filters) AddExists(name string, value bool) *filters {
	return f.addOp("EXISTS", name, strconv.FormatBool(value))
}

func (f *filters) AddGT(name, value string) *filters {
	return f.addOp("GT", name, value)
}

func (f *filters) AddGTE(name, value string) *filters {
	return f.addOp("GTE", name, value)
}

func (f *filters) AddLT(name, value string) *filters {
	return f.addOp("LT", name, value)
}

func (f *filters) AddLTE(name, value string) *filters {
	return f.addOp("LTE", name, value)
}

func (f *filters) AddNE(name, value string) *filters {
	return f.addOp("NE", name, value)
}

func (f *filters) AddLike(name, value string) *filters {
	return f.addOp("LIKE", name, value)
}

func (f *filters) AddNotLike(name, value string) *filters {
	return f.addOp("NOTLIKE", name, value)
}

func (f *filters) AddIn(name, value string) *filters {
	return f.addOp("IN", name, value)
}

func (f *filters) AddNotIn(name, value string) *filters {
	return f.addOp("NOTIN", name, value)
}

func (f *filters) AddWithin(name, value, shape, radius string) *filters {
	*f = append(*f, &filter{
		XMLName: xml.Name{Local: "WITHIN"},
		Name:    name,
		Value:   value,
		Shape:   shape,
		Radius:  radius,
	})
	return f
}

func (f *filters) AddIntersects(name, value string) *filters {
	return f.addOp("INTERSECTS", name, value)
}

func (f *filters) AddNear(name, value string) *filters {
	return f.addOp("NEAR", name, value)
}

// Clauses (AND, OR, ...)

func (f *filters) newOp(tag string) *filters {
	newOp := &filters{}
	*f = append(*f, &filter{
		XMLName: xml.Name{Local: tag},
		Filters: newOp,
	})
	return newOp
}

func (f *filters) NewOr() *filters {
	return f.newOp("OR")
}

func (f *filters) NewAnd() *filters {
	return f.newOp("AND")
}

func (f *filters) NewElementMatch() *filters {
	return f.newOp("ELEMENTMATCH")
}

func (f *filters) NewNot() *filters {
	return f.newOp("NOT")
}
