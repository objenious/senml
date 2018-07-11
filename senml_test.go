package senml

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"
)

func TestEquals(t *testing.T) {
	tcs := []struct {
		a   Pack
		b   Pack
		res bool
	}{
		{
			a: Pack{
				{Name: "foo", Value: Float(1)},
				{Name: "bar", BoolValue: Bool(true)},
			},
			b: Pack{
				{Name: "foo", Value: Float(1)},
				{Name: "bar", BoolValue: Bool(true)},
			},
			res: true,
		},
		{
			a: Pack{
				{Name: "foo", Value: Float(1)},
			},
			b:   nil,
			res: false,
		},
		{
			a: nil,
			b: Pack{
				{Name: "foo", Value: Float(1)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", Value: Float(1)},
			},
			b: Pack{
				{Name: "foo", Value: Float(1)},
				{Name: "foo", Value: Float(1)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", Value: Float(1)},
				{Name: "foo", Value: Float(1)},
			},
			b: Pack{
				{Name: "foo", Value: Float(1)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", Value: Float(1)},
			},
			b: Pack{
				{Name: "foo", Value: Float(2)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", Value: Float(1)},
			},
			b: Pack{
				{Name: "foo"},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BoolValue: Bool(true)},
			},
			b: Pack{
				{Name: "foo", BoolValue: Bool(false)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BoolValue: Bool(true)},
			},
			b: Pack{
				{Name: "foo"},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BoolValue: Bool(true)},
			},
			b: Pack{
				{Name: "bar", BoolValue: Bool(true)},
			},
			res: false,
		},
		{
			a: Pack{
				{BaseName: "foo", BoolValue: Bool(true)},
			},
			b: Pack{
				{BaseName: "bar", BoolValue: Bool(true)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BaseValue: Float(1)},
			},
			b: Pack{
				{Name: "foo", BaseValue: Float(2)},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BaseValue: Float(1)},
			},
			b: Pack{
				{Name: "foo"},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BaseTime: 1},
			},
			b: Pack{
				{Name: "foo", BaseTime: 2},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", Time: 1},
			},
			b: Pack{
				{Name: "foo", Time: 2},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", UpdateTime: 1},
			},
			b: Pack{
				{Name: "foo", UpdateTime: 2},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", Unit: Ampere},
			},
			b: Pack{
				{Name: "foo", Unit: Volt},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", BaseUnit: Ampere},
			},
			b: Pack{
				{Name: "foo", BaseUnit: Volt},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", StringValue: "foo"},
			},
			b: Pack{
				{Name: "foo", StringValue: "bar"},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2}},
			},
			b: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2}},
			},
			res: true,
		},
		{
			a: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2}},
			},
			b: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x3}},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2}},
			},
			b: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2, 0x3}},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2, 0x3}},
			},
			b: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2}},
			},
			res: false,
		},
		{
			a: Pack{
				{Name: "foo", DataValue: []byte{0x1, 0x2}},
			},
			b: Pack{
				{Name: "foo", DataValue: nil},
			},
			res: false,
		},
	}
	for _, tc := range tcs {
		if tc.a.Equals(tc.b) != tc.res {
			t.Errorf("Equals with %+v and %+v should return %v", tc.a, tc.b, tc.res)
		}
	}
}

func TestNormalize(t *testing.T) {
	tcs := []struct {
		src  Pack
		norm Pack
	}{
		{
			src: Pack{
				{BaseName: "urn:dev:ow:10e2073a01080063", BaseTime: 1.320067464e+09, BaseUnit: RelativeHumidity, Value: Float(20)},
				{Unit: DegreesLongitude, Value: Float(24.30621)},
				{Unit: DegreesLatitude, Value: Float(60.07965)},
				{Time: 60, Value: Float(20.3)},
				{Unit: DegreesLongitude, Time: 60, Value: Float(24.30622)},
				{Unit: DegreesLatitude, Time: 60, Value: Float(60.07965)},
				{Time: 120, Value: Float(20.7)},
				{Unit: DegreesLongitude, Time: 120, Value: Float(24.30623)},
				{Unit: DegreesLatitude, Time: 120, Value: Float(60.07966)},
				{Unit: EnergyLevel, Time: 150, Value: Float(98)},
			},
			norm: Pack{
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067464e+09, Unit: RelativeHumidity, Value: Float(20)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067464e+09, Unit: DegreesLongitude, Value: Float(24.30621)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067464e+09, Unit: DegreesLatitude, Value: Float(60.07965)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067524e+09, Unit: RelativeHumidity, Value: Float(20.3)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067524e+09, Unit: DegreesLongitude, Value: Float(24.30622)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067524e+09, Unit: DegreesLatitude, Value: Float(60.07965)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067584e+09, Unit: RelativeHumidity, Value: Float(20.7)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067584e+09, Unit: DegreesLongitude, Value: Float(24.30623)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067584e+09, Unit: DegreesLatitude, Value: Float(60.07966)},
				{Name: "urn:dev:ow:10e2073a01080063", Time: 1.320067614e+09, Unit: EnergyLevel, Value: Float(98)},
			},
		},
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Value: Float(1)},
			},
		},
		{
			src: Pack{
				{BaseValue: Float(1)},
				{Name: "foo", Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo", Value: Float(2)},
			},
		},
		{
			src: Pack{
				{BaseSum: Float(1)},
				{Name: "foo", Sum: Float(1)},
			},
			norm: Pack{
				{Name: "foo", Sum: Float(2)},
			},
		},
		{
			src: Pack{
				{BaseTime: 1},
				{Name: "foo", Time: 1, Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo", Time: 2, Value: Float(1)},
			},
		},
		{
			src: Pack{
				{Name: "foo", Time: 1, Value: Float(1)},
				{Name: "foo", Time: 1, BoolValue: Bool(true)},
				{Name: "foo", Time: 1, StringValue: "foo"},
				{Name: "foo", Time: 1, DataValue: []byte{0x01, 0x02}},
				{Name: "foo", Time: 1, Sum: Float(1)},
			},
			norm: Pack{
				{Name: "foo", Time: 1, Value: Float(1)},
				{Name: "foo", Time: 1, BoolValue: Bool(true)},
				{Name: "foo", Time: 1, StringValue: "foo"},
				{Name: "foo", Time: 1, DataValue: []byte{0x01, 0x02}},
				{Name: "foo", Time: 1, Sum: Float(1)},
			},
		},
		{
			src: Pack{
				{Name: "foo", Time: 2, Value: Float(1)},
				{Name: "foo", Time: 1, Value: Float(2)},
			},
			norm: Pack{
				{Name: "foo", Time: 1, Value: Float(2)},
				{Name: "foo", Time: 2, Value: Float(1)},
			},
		},
	}
	for _, tc := range tcs {
		norm := tc.src.Normalize()
		if !norm.Equals(tc.norm) {
			t.Errorf("Normalized version of %+v should be %+v not %+v", tc.src, tc.norm, norm)
		}
	}
}

func TestNormalizeAt(t *testing.T) {
	t0 := time.Now()
	tcs := []struct {
		src  Pack
		norm Pack
	}{
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Time: Time(t0), Value: Float(1)},
			},
		},
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Time: 1, Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Time: Time(t0) + 1, Value: Float(1)},
			},
		},
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Time: -268435457, Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Time: Time(t0) - 268435457, Value: Float(1)},
			},
		},
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Time: 268435456, Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Time: Time(t0) + 268435456, Value: Float(1)},
			},
		},
		{
			src: Pack{
				{BaseName: "foo."},
				{Name: "bar", Time: 268435457, Value: Float(1)},
			},
			norm: Pack{
				{Name: "foo.bar", Time: 268435457, Value: Float(1)},
			},
		},
	}
	for _, tc := range tcs {
		norm := tc.src.NormalizeAt(t0)
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
				{Name: "urn:dev:ow:10e2073a01080063", Unit: Celsius, Value: Float(23.1)},
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

		dec := Pack{}
		err = json.Unmarshal([]byte(tc.json), &dec)
		if err != nil {
			t.Errorf("JSON decoding of %s returned an error : %s", tc.json, err)
		}
		if !tc.src.Equals(dec) {
			t.Errorf("JSON decoding of %s should be %+v not %+v", tc.json, tc.src, dec)
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
				{Name: "urn:dev:ow:10e2073a01080063", Unit: Celsius, Value: Float(23.1)},
			},
			xml: `<sensml xmlns="urn:ietf:params:xml:ns:senml"><senml n="urn:dev:ow:10e2073a01080063" u="Cel" v="23.1"></senml></sensml>`,
		},
		{
			src: Pack{
				{BaseName: "urn:dev:ow:10e2073a01080063", BaseTime: 1.276020076001e+09, BaseUnit: Ampere, Version: 5, Name: "voltage", Unit: Volt, Value: Float(120.1)},
				{Name: "current", Time: -5, Value: Float(1.2)},
				{Name: "current", Time: -4, Value: Float(1.3)},
				{Name: "current", Time: -3, Value: Float(1.4)},
				{Name: "current", Time: -2, Value: Float(1.5)},
				{Name: "current", Time: -1, Value: Float(1.6)},
				{Name: "current", Value: Float(1.7)},
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

		dec := Pack{}
		err = xml.Unmarshal([]byte(tc.xml), &dec)
		if err != nil {
			t.Errorf("XML decoding of %s returned an error : %s", tc.xml, err)
		}
		if !tc.src.Equals(dec) {
			t.Errorf("XML decoding of %s should be %+v not %+v", tc.xml, tc.src, dec)
		}
	}
}
