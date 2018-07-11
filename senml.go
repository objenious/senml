package senml

import (
	"encoding/xml"
	"sort"
	"time"
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

// Normalize resolves the SenML Records, as explained in https://tools.ietf.org/html/draft-ietf-core-senml-16#section-4.6.
// All base items are removed, and records are sorted in chronological order.
func (p Pack) Normalize() Pack {
	var bname string
	var bunit Unit
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

// NormalizeAt resolves the SenML Records, and replaces all relative times
// by absolute times, based on the t reference time.
func (p Pack) NormalizeAt(t time.Time) Pack {
	n := p.Normalize()
	rt := Time(t)
	for i := range n {
		n[i].Time = absoluteTime(n[i].Time, rt)
	}
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
	err := d.DecodeElement(&n, &start)
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
