package senml

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

func fptr(f float64) *float64 {
	return &f
}

func TestNormalize(t *testing.T) {
	tcs := []struct {
		src  Pack
		norm Pack
	}{
		{
			src: Pack{
				{BaseName: "urn:dev:ow:10e2073a01080063", BaseTime: 1.320067464e+09, BaseUnit: "%RH", Value: fptr(20)},
				{Unit: "lon", Value: fptr(24.30621)},
				{Unit: "lat", Value: fptr(60.07965)},
				{Time: 60, Value: fptr(20.3)},
				{Unit: "lon", Time: 60, Value: fptr(24.30622)},
				{Unit: "lat", Time: 60, Value: fptr(60.07965)},
				{Time: 120, Value: fptr(20.7)},
				{Unit: "lon", Time: 120, Value: fptr(24.30623)},
				{Unit: "lat", Time: 120, Value: fptr(60.07966)},
				{Unit: "%EL", Time: 150, Value: fptr(98)},
			},
			norm: Pack{
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067464e+09, Unit: "%RH", Value: fptr(20), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067464e+09, Unit: "lon", Value: fptr(24.30621), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067464e+09, Unit: "lat", Value: fptr(60.07965), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067524e+09, Unit: "%RH", Value: fptr(20.3), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067524e+09, Unit: "lon", Value: fptr(24.30622), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067524e+09, Unit: "lat", Value: fptr(60.07965), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067584e+09, Unit: "%RH", Value: fptr(20.7), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067584e+09, Unit: "lon", Value: fptr(24.30623), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067584e+09, Unit: "lat", Value: fptr(60.07966), Version: 5},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067614e+09, Unit: "%EL", Value: fptr(98), Version: 5},
			},
		},
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Value: fptr(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Value: fptr(1), Version: 5},
			},
		},
		{
			src: Pack{
				{BaseValue: fptr(1)},
				{Name: "foo", Value: fptr(1)},
			},
			norm: Pack{
				{Name: "foo", Value: fptr(2), Version: 5},
			},
		},
		{
			src: Pack{
				{BaseSum: fptr(1)},
				{Name: "foo", Sum: fptr(1)},
			},
			norm: Pack{
				{Name: "foo", Sum: fptr(2), Version: 5},
			},
		},
	}
	for _, tc := range tcs {
		norm := tc.norm.Normalize()
		if !norm.Equals(tc.norm) {
			t.Errorf("Normalized version of %+v should be %+v not %+v", tc.src, tc.norm, norm)
		}
	}
}

func TestJSON(t *testing.T) {
	tcs := []struct {
		src  Pack
		json string
	}{
		{
			src: Pack{
				{Name: "urn:dev:ow:10e2073a01080063", Unit: "Cel", Value: fptr(23.1)},
			},
			json: `[{"n":"urn:dev:ow:10e2073a01080063","u":"Cel","v":23.1}]`,
		},
	}
	for _, tc := range tcs {
		enc, err := json.Marshal(tc.src)
		if err != nil {
			t.Errorf("JSON encoding of %+v returned an error : %s", tc.src, err)
		}
		if string(enc) != tc.json {
			t.Errorf("JSON encoding of %+v should be %s not %s", tc.src, tc.json, enc)
		}
	}
}
func TestXML(t *testing.T) {
	tcs := []struct {
		src Pack
		xml string
	}{
		{
			src: Pack{
				{Name: "urn:dev:ow:10e2073a01080063", Unit: "Cel", Value: fptr(23.1)},
			},
			xml: `<sensml xmlns="urn:ietf:params:xml:ns:senml"><senml n="urn:dev:ow:10e2073a01080063" u="Cel" v="23.1"></senml></sensml>`,
		},
		{
			src: Pack{
				{BaseName: "urn:dev:ow:10e2073a01080063", BaseTime: 1.276020076001e+09, BaseUnit: "A", Version: 5, Name: "voltage", Unit: "V", Value: fptr(120.1)},
				{Name: "current", Time: -5, Value: fptr(1.2)},
				{Name: "current", Time: -4, Value: fptr(1.3)},
				{Name: "current", Time: -3, Value: fptr(1.4)},
				{Name: "current", Time: -2, Value: fptr(1.5)},
				{Name: "current", Time: -1, Value: fptr(1.6)},
				{Name: "current", Value: fptr(1.7)},
			},
			xml: `<sensml xmlns="urn:ietf:params:xml:ns:senml"><senml bn="urn:dev:ow:10e2073a01080063" bt="1.276020076001e+09" bu="A" bver="5" n="voltage" u="V" v="120.1"></senml><senml n="current" t="-5" v="1.2"></senml><senml n="current" t="-4" v="1.3"></senml><senml n="current" t="-3" v="1.4"></senml><senml n="current" t="-2" v="1.5"></senml><senml n="current" t="-1" v="1.6"></senml><senml n="current" v="1.7"></senml></sensml>`,
		},
	}
	for _, tc := range tcs {
		enc, err := xml.Marshal(tc.src)
		if err != nil {
			t.Errorf("XML encoding of %+v returned an error : %s", tc.src, err)
		}
		if string(enc) != tc.xml {
			t.Errorf("XML encoding of %+v should be %s not %s", tc.src, tc.xml, enc)
		}
	}
}
