package senml

import (
	"encoding/xml"
	"sort"
)

// Pack defines a SenML pack (a list of Records).
type Pack []Record

// Equals checks if 2 packs are equal.
func (p Pack) Equals(p2 Pack) bool {
	if (p == nil && p2 != nil) || (p != nil && p2 == nil) {
		return false
	}
	if len(p) != len(p2) {
		return false
	}
	for i := range p {
		if !p[i].Equals(&p2[i]) {
			return false
		}
	}
	return true
}

// Record is a SenML Record.
type Record struct {
	BaseName  string   `json:"bn,omitempty"  xml:"bn,attr,omitempty"`
	BaseTime  float64  `json:"bt,omitempty"  xml:"bt,attr,omitempty"`
	BaseUnit  string   `json:"bu,omitempty"  xml:"bu,attr,omitempty"`
	BaseValue *float64 `json:"bv,omitempty"  xml:"bv,attr,omitempty"`
	BaseSum   *float64 `json:"bs,omitempty"  xml:"bs,attr,omitempty"`

	Version int `json:"bver,omitempty"  xml:"bver,attr,omitempty"`

	Name string `json:"n,omitempty"  xml:"n,attr,omitempty"`
	Unit string `json:"u,omitempty"  xml:"u,attr,omitempty"`

	Time       float64 `json:"t,omitempty"  xml:"t,attr,omitempty"`
	UpdateTime float64 `json:"ut,omitempty"  xml:"ut,attr,omitempty"`

	Value       *float64 `json:"v,omitempty"  xml:"v,attr,omitempty"`
	StringValue string   `json:"vs,omitempty"  xml:"vs,attr,omitempty"`
	DataValue   []byte   `json:"vd,omitempty"  xml:"vd,attr,omitempty"`
	BoolValue   *bool    `json:"vb,omitempty"  xml:"vb,attr,omitempty"`
	Sum         *float64 `json:"s,omitempty"  xml:"sum,attr,omitempty"`
}

// Equals checks if two records are equal
func (r *Record) Equals(r2 *Record) bool {
	if (r == nil && r2 != nil) || (r != nil && r2 == nil) {
		return false
	}
	if r.BaseName != r2.BaseName {
		return false
	}
	if r.BaseTime != r2.BaseTime {
		return false
	}
	if r.BaseUnit != r2.BaseUnit {
		return false
	}
	if r.Version != r2.Version {
		return false
	}
	if r.Name != r2.Name {
		return false
	}
	if r.Unit != r2.Unit {
		return false
	}
	if r.Time != r2.Time {
		return false
	}
	if r.UpdateTime != r2.UpdateTime {
		return false
	}

	if r.Value != nil && r2.Value != nil && *r.Value == *r2.Value {
		return true
	}
	if r.StringValue != "" && r.StringValue == r2.StringValue {
		return true
	}
	if r.BoolValue != nil && r2.BoolValue != nil && *r.BoolValue == *r2.BoolValue {
		return true
	}
	if r.Sum != nil && r2.Sum != nil && *r.Sum == *r2.Sum {
		return true
	}
	if r.DataValue != nil && r2.DataValue != nil && len(r.DataValue) == len(r2.DataValue) {
		for i := range r.DataValue {
			if r.DataValue[i] != r2.DataValue[i] {
				return false
			}
		}
		return true
	}
	return false
}

// Normalize resolves the SenML Records, as explained in https://tools.ietf.org/html/draft-ietf-core-senml-16#section-4.6.
// All base items are removed, and records are sorted in chronological order.
func (p Pack) Normalize() Pack {
	var bname, bunit string
	var btime, bval, bsum float64
	var bver int

	n := make(Pack, 0, len(p))
	for i := range p {
		if p[i].BaseTime != 0 {
			btime = p[i].BaseTime
		}
		if p[i].Version != 0 {
			bver = p[i].Version
		}
		if p[i].BaseUnit != "" {
			bunit = p[i].BaseUnit
		}
		if p[i].BaseName != "" {
			bname = p[i].BaseName
		}
		if p[i].BaseValue != nil {
			bval = *p[i].BaseValue
		}
		if p[i].BaseSum != nil {
			bsum = *p[i].BaseSum
		}
		r := Record{
			Name:    bname + p[i].Name,
			Time:    btime + p[i].Time,
			Unit:    bunit,
			Version: bver,
		}
		if p[i].Unit != "" {
			r.Unit = p[i].Unit
		}
		switch {
		case p[i].Value != nil:
			nval := bval + *p[i].Value
			r.Value = &nval
		case p[i].BoolValue != nil:
			r.BoolValue = p[i].BoolValue
		case p[i].StringValue != "":
			r.StringValue = p[i].StringValue
		case len(p[i].DataValue) > 0:
			r.DataValue = p[i].DataValue
		case p[i].Sum != nil:
			nsum := bsum + *p[i].Sum
			r.Sum = &nsum
		default:
			continue
		}
		n = append(n, r)
	}
	sort.Sort(&n)
	return n
}

// Len implements sort.Interface.
func (p Pack) Len() int {
	return len(p)
}

// Less implements sort.Interface.
func (p Pack) Less(i, j int) bool {
	return p[i].Time < p[j].Time
}

// Swap implements sort.Interface.
func (p Pack) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// MarshalXML implements xml.Marshaler. It encodes the SenML Pack to XML.
func (p Pack) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n := xmlPack{
		Xmlns:   "urn:ietf:params:xml:ns:senml",
		Records: make([]xmlRecord, len(p)),
	}
	for i := range p {
		n.Records[i].Record = p[i]
	}
	return e.Encode(n)
}

// UnmarshalXML implements xml.Unmarshaler. It decodes a XML encoded SenML Pack.
func (p *Pack) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n := xmlPack{}
	err := d.Decode(&n)
	if err != nil {
		return err
	}
	*p = make(Pack, len(n.Records))
	for i := range *p {
		(*p)[i] = n.Records[i].Record
	}
	return nil
}

type xmlPack struct {
	XMLName *bool       `xml:"sensml"`
	Xmlns   string      `xml:"xmlns,attr"`
	Records []xmlRecord `xml:"senml"`
}

type xmlRecord struct {
	XMLName *bool `json:"_,omitempty" xml:"senml"`
	Record  `xml:",innerxml"`
}
