package senml

import "time"

// Record is a SenML Record.
type Record struct {
	BaseName  string   `json:"bn,omitempty"  xml:"bn,attr,omitempty"`
	BaseTime  float64  `json:"bt,omitempty"  xml:"bt,attr,omitempty"`
	BaseUnit  Unit     `json:"bu,omitempty"  xml:"bu,attr,omitempty"`
	BaseValue *float64 `json:"bv,omitempty"  xml:"bv,attr,omitempty"`
	BaseSum   *float64 `json:"bs,omitempty"  xml:"bs,attr,omitempty"`

	BaseVersion int `json:"bver,omitempty"  xml:"bver,attr,omitempty"`

	Name string `json:"n,omitempty"  xml:"n,attr,omitempty"`
	Unit Unit   `json:"u,omitempty"  xml:"u,attr,omitempty"`

	Time       float64 `json:"t,omitempty"  xml:"t,attr,omitempty"`
	UpdateTime float64 `json:"ut,omitempty"  xml:"ut,attr,omitempty"`

	Value       *float64 `json:"v,omitempty"  xml:"v,attr,omitempty"`
	StringValue string   `json:"vs,omitempty"  xml:"vs,attr,omitempty"`
	DataValue   []byte   `json:"vd,omitempty"  xml:"vd,attr,omitempty"`
	BoolValue   *bool    `json:"vb,omitempty"  xml:"vb,attr,omitempty"`
	Sum         *float64 `json:"s,omitempty"  xml:"s,attr,omitempty"`
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
	if r.BaseVersion != r2.BaseVersion {
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

// GoTime returns the Time of the Record as a Go time.Time.
func (r *Record) GoTime() time.Time {
	return GoTime(r.Time)
}
