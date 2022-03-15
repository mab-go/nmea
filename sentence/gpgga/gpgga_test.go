package gpgga

import (
	"fmt"
	"testing"
)

type testVec struct {
	input    string
	expected GPGGA
	errMsg   string
}

var goodTestData = map[string]testVec{
	"NMEA Generator [1/1]": {
		input: "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*70",
		expected: GPGGA{
			FixTime:        174800.864,
			Latitude:       4002.741,
			NorthSouth:     North,
			Longitude:      7618.55,
			EastWest:       West,
			FixQuality:     GPSFixQuality,
			SatCount:       12,
			HDOP:           1.0,
			Altitude:       0.0,
			AltitudeUOM:    "M",
			GeoidHeight:    0.0,
			GeoidHeightUOM: "M",
		},
	},
	"Garmin G12 (v 4.57)": {
		input: "$GPGGA,183730,3907.356,N,12102.482,W,1,05,1.6,646.4,M,-24.1,M,300,123*76",
		expected: GPGGA{
			FixTime:        183730.0,
			Latitude:       3907.356,
			NorthSouth:     North,
			Longitude:      12102.482,
			EastWest:       West,
			FixQuality:     GPSFixQuality,
			SatCount:       5,
			HDOP:           1.6,
			Altitude:       646.4,
			AltitudeUOM:    "M",
			GeoidHeight:    -24.1,
			GeoidHeightUOM: "M",
			DGPSUpdateAge:  300.0,
			DGPSStationID:  123,
		},
	},
	"Garmin eTrex Summit": {
		input: "$GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*7F",
		expected: GPGGA{
			FixTime:        2454.0,
			Latitude:       3553.5295,
			NorthSouth:     North,
			Longitude:      13938.657,
			EastWest:       East,
			FixQuality:     GPSFixQuality,
			SatCount:       5,
			HDOP:           2.2,
			Altitude:       18.3,
			AltitudeUOM:    "M",
			GeoidHeight:    39.0,
			GeoidHeightUOM: "M",
		},
	},
}

var badTestData = map[string]testVec{
	"Bad SentenceType": {
		input:  "$GAAAA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*61",
		errMsg: "sentence segment [0] must be \"GPGGA\" (case insensitive) but was \"GAAAA\"",
	},
	"Bad FixTime": {
		input:  "$GPGGA,bad_FixTime,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*34",
		errMsg: "sentence segment [1] must be parsable as a float32 but was \"bad_FixTime\"",
	},
	"Bad Latitude": {
		input:  "$GPGGA,174800.864,bad_Latitude,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*62",
		errMsg: "sentence segment [2] must be parsable as a float64 but was \"bad_Latitude\"",
	},
	"Bad NorthSouth": {
		input:  "$GPGGA,174800.864,4002.741,bad_NorthSouth,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*1C",
		errMsg: "sentence segment [3] must be parsable as a NorthSouth but was \"bad_NorthSouth\"",
	},
	"Bad Longitude": {
		input:  "$GPGGA,174800.864,4002.741,N,bad_Longitude,W,1,12,1.0,0.0,M,0.0,M,,*2D",
		errMsg: "sentence segment [4] must be parsable as a float64 but was \"bad_Longitude\"",
	},
	"Bad EastWest": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,bad_EastWest,1,12,1.0,0.0,M,0.0,M,,*09",
		errMsg: "sentence segment [5] must be parsable as an EastWest but was \"bad_EastWest\"",
	},
	"Bad FixQuality (Wrong Type)": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,bad_FixQuality,12,1.0,0.0,M,0.0,M,,*63",
		errMsg: "sentence segment [6] must be parsable as a FixQuality but was \"bad_FixQuality\"",
	},
	"Bad FixQuality (Out of Range)": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,9,12,1.0,0.0,M,0.0,M,,*78",
		errMsg: "sentence segment [6] must be parsable as a FixQuality but was \"9\"",
	},
	"Bad FixQuality (Negative Number)": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,-2,12,1.0,0.0,M,0.0,M,,*5E",
		errMsg: "sentence segment [6] must be parsable as a FixQuality but was \"-2\"",
	},
	"Bad SatCount": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,bad_SatCount,1.0,0.0,M,0.0,M,,*4E",
		errMsg: "sentence segment [7] must be parsable as an int8 but was \"bad_SatCount\"",
	},
	"Bad HDOP": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,bad_HDOP,0.0,M,0.0,M,,*74",
		errMsg: "sentence segment [8] must be parsable as a float32 but was \"bad_HDOP\"",
	},
	"Bad Altitude": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,bad_Altitude,M,0.0,M,,*56",
		errMsg: "sentence segment [9] must be parsable as a float32 but was \"bad_Altitude\"",
	},
	"Bad AltitudeUOM": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,bad_AltitudeUOM,0.0,M,,*62",
		errMsg: "sentence segment [10] must be \"M\" (case insensitive) but was \"bad_AltitudeUOM\"",
	},
	"Bad GeoidHeight": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,bad_GeoidHeight,M,,*19",
		errMsg: "sentence segment [11] must be parsable as a float32 but was \"bad_GeoidHeight\"",
	},
	"Bad GeoidHeightUOM": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,bad_GeoidHeightUOM,,*2D",
		errMsg: "sentence segment [12] must be \"M\" (case insensitive) but was \"bad_GeoidHeightUOM\"",
	},
	"Bad DGPSUpdateAge": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,bad_DGPSUpdateAge,*3A",
		errMsg: "sentence segment [13] must be parsable as a float32 but was \"bad_DGPSUpdateAge\"",
	},
	"Bad DGPSStationID": {
		input:  "$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,bad_DGPSStationID*1F",
		errMsg: "sentence segment [14] must be parsable as an int16 but was \"bad_DGPSStationID\"",
	},
}

func assertMatches(t *testing.T, title, field string, expected, actual interface{}) {
	if actual != expected {
		t.Errorf("%s should have been %v but was %v for NMEA input \"%v\"", field, expected, actual, title)
	}
}

func TestParse_goodData(t *testing.T) {
	for title, vec := range goodTestData {
		t.Run(title, func(t *testing.T) {
			actual, err := Parse(vec.input)
			if err != nil {
				t.Fatalf("error creating GPGLL from NMEA input \"%v\": %v", title, err)
			}

			expected := vec.expected
			assertMatches(t, title, "FixTime", expected.FixTime, actual.FixTime)
			assertMatches(t, title, "Latitude", expected.Latitude, actual.Latitude)
			assertMatches(t, title, "NorthSouth", expected.NorthSouth, actual.NorthSouth)
			assertMatches(t, title, "Longitude", expected.Longitude, actual.Longitude)
			assertMatches(t, title, "EastWest", expected.EastWest, actual.EastWest)
			assertMatches(t, title, "FixQuality", expected.FixQuality, actual.FixQuality)
			assertMatches(t, title, "SatCount", expected.SatCount, actual.SatCount)
			assertMatches(t, title, "HDOP", expected.HDOP, actual.HDOP)
			assertMatches(t, title, "Altitude", expected.Altitude, actual.Altitude)
			assertMatches(t, title, "AltitudeUOM", expected.AltitudeUOM, actual.AltitudeUOM)
			assertMatches(t, title, "GeoidHeight", expected.GeoidHeight, actual.GeoidHeight)
			assertMatches(t, title, "GeoidHeightUOM", expected.GeoidHeightUOM, actual.GeoidHeightUOM)
			assertMatches(t, title, "DGPSUpdateAge", expected.DGPSUpdateAge, actual.DGPSUpdateAge)
			assertMatches(t, title, "DGPSStationID", expected.DGPSStationID, actual.DGPSStationID)
		})
	}
}

func TestParse_invalidChecksum(t *testing.T) {
	gpgga, err := Parse("$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*42")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	if gpgga != nil {
		t.Errorf("result should have been <nil> but was %v", gpgga)
	}

	expected := "calculated checksum value \"70\" does not match sentence-specified value of \"42\""
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestParse_badSegments(t *testing.T) {
	for title, vec := range badTestData {
		t.Run(title, func(t *testing.T) {
			gpgga, err := Parse(vec.input)
			if err == nil {
				t.Fatalf("parsing succeeded (but should not have) for test sentence %q", title)
			}

			if gpgga != nil {
				t.Fatalf("result should have been <nil> but was %v for test sentence %q", gpgga, title)
			}

			if err.Error() != vec.errMsg {
				t.Fatalf("error message should have been '%v' but was '%v' for test sentence %q", vec.errMsg, err.Error(), title)
			}
		})
	}
}

func TestGPGGA_GetSentenceType(t *testing.T) {
	gpgga := &GPGGA{}
	if st := gpgga.GetSentenceType(); st != "GPGGA" {
		t.Errorf("GetSentenceType() should have returned \"GPGGA\" but returned \"%v\"", st)
	}
}

func ExampleParse() {
	sentence := "$GPGGA,023042,3907.3837,N,12102.4684,W,1,04,2.3,507.3,M,-24.1,M,,*75"
	gpgga, err := Parse(sentence)
	if err != nil {
		// Handle error
	}

	fmt.Printf("%+v", gpgga)
	// Output:
	// &{FixTime:23042 Latitude:3907.3837 NorthSouth:N Longitude:12102.4684 EastWest:W FixQuality:1 SatCount:4 HDOP:2.3 Altitude:507.3 AltitudeUOM:M GeoidHeight:-24.1 GeoidHeightUOM:M DGPSUpdateAge:0 DGPSStationID:0}
}
