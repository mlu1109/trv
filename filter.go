package trv

import (
	"encoding/xml"
	"strconv"
)

// Filter can be an operator (EQ, GT, ...) or a clause (AND, OR, ...).
// Should be constructed with addOp or newOp
type Filter struct {
	XMLName xml.Name
	Name    string `xml:"name,omitempty,attr"`
	Value   string `xml:"value,omitempty,attr"`
	Shape   string `xml:"shape,omitempty"`
	Radius  string `xml:"radius,omitempty"`
	Filters *Filters
}

type Filters []*Filter

// Operators (EQ, GT, ...)

func (f *Filters) addOp(tag, name, value string) *Filters {
	*f = append(*f, &Filter{
		XMLName: xml.Name{Local: tag},
		Name:    name,
		Value:   value,
	})
	return f
}

func (f *Filters) AddEQ(name, value string) *Filters {
	return f.addOp("EQ", name, value)
}

func (f *Filters) AddExists(name string, value bool) *Filters {
	return f.addOp("EXISTS", name, strconv.FormatBool(value))
}

func (f *Filters) AddGT(name, value string) *Filters {
	return f.addOp("GT", name, value)
}

func (f *Filters) AddGTE(name, value string) *Filters {
	return f.addOp("GTE", name, value)
}

func (f *Filters) AddLT(name, value string) *Filters {
	return f.addOp("LT", name, value)
}

func (f *Filters) AddLTE(name, value string) *Filters {
	return f.addOp("LTE", name, value)
}

func (f *Filters) AddNE(name, value string) *Filters {
	return f.addOp("NE", name, value)
}

func (f *Filters) AddLike(name, value string) *Filters {
	return f.addOp("LIKE", name, value)
}

func (f *Filters) AddNotLike(name, value string) *Filters {
	return f.addOp("NOTLIKE", name, value)
}

func (f *Filters) AddIn(name, value string) *Filters {
	return f.addOp("IN", name, value)
}

func (f *Filters) AddNotIn(name, value string) *Filters {
	return f.addOp("NOTIN", name, value)
}

func (f *Filters) AddWithin(name, value, shape, radius string) *Filters {
	*f = append(*f, &Filter{
		XMLName: xml.Name{Local: "WITHIN"},
		Name:    name,
		Value:   value,
		Shape:   shape,
		Radius:  radius,
	})
	return f
}

func (f *Filters) AddIntersects(name, value string) *Filters {
	return f.addOp("INTERSECTS", name, value)
}

func (f *Filters) AddNear(name, value string) *Filters {
	return f.addOp("NEAR", name, value)
}

// Clauses (AND, OR, ...)

func (f *Filters) newOp(tag string) *Filters {
	newOp := &Filters{}
	*f = append(*f, &Filter{
		XMLName: xml.Name{Local: tag},
		Filters: newOp,
	})
	return newOp
}

func (f *Filters) NewOr() *Filters {
	return f.newOp("OR")
}

func (f *Filters) NewAnd() *Filters {
	return f.newOp("AND")
}

func (f *Filters) NewElementMatch() *Filters {
	return f.newOp("ELEMENTMATCH")
}

func (f *Filters) NewNot() *Filters {
	return f.newOp("NOT")
}
