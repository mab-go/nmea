package gpgll

import (
	"fmt"
	"testing"
)

type testVec struct {
	input    string
	expected GPGLL
	errMsg   string
}

var goodTestData = map[string]testVec{
	"NMEASimulator (Modified) [1/4]": {
		input: "$GPGLL,3157.905722,S,11551.681852,E,215052.603,A,D*4F",
		expected: GPGLL{
			Latitude:   3157.905722,
			NorthSouth: South,
			Longitude:  11551.681852,
			EastWest:   East,
			FixTime:    215052.603,
			DataStatus: ValidDataStatus,
			Mode:       DifferentialMode,
		},
	},
	"NMEA Simulator (Modified) [2/4]": {
		input: "$GPGLL,3157.905722,S,11551.681852,E,215102.604,A,E*4D",
		expected: GPGLL{
			Latitude:   3157.905722,
			NorthSouth: South,
			Longitude:  11551.681852,
			EastWest:   East,
			FixTime:    215102.604,
			DataStatus: ValidDataStatus,
			Mode:       EstimatedMode,
		},
	},
	"NMEA Simulator (Modified) [3/4]": {
		input: "$GPGLL,3726.489023,N,12212.446039,W,214827.478,A,M*4C",
		expected: GPGLL{
			Latitude:   3726.489023,
			NorthSouth: North,
			Longitude:  12212.446039,
			EastWest:   West,
			FixTime:    214827.478,
			DataStatus: ValidDataStatus,
			Mode:       ManualInputMode,
		},
	},
	"NMEA Simulator (Modified) [4/4]": {
		input: "$GPGLL,3726.489023,N,12212.446039,W,214916.479,A,A*42",
		expected: GPGLL{
			Latitude:   3726.489023,
			NorthSouth: North,
			Longitude:  12212.446039,
			EastWest:   West,
			FixTime:    214916.479,
			DataStatus: ValidDataStatus,
			Mode:       AutonomousMode,
		},
	},
	// Example from https://www.rfwireless-world.com/Terminology/GPS-sentences-or-NMEA-sentences.html
	"RF Wireless World Example": {
		input: "$GPGLL,3723.2475,N,12158.3416,W,161229.487,A,A*41",
		expected: GPGLL{
			Latitude:   3723.2475,
			NorthSouth: North,
			Longitude:  12158.3416,
			EastWest:   West,
			FixTime:    161229.487,
			DataStatus: ValidDataStatus,
			Mode:       AutonomousMode,
		},
	},
}

var badTestData = map[string]testVec{
	"Bad SentenceType": {
		input:  "$GPFOO,3723.2475,N,12158.3416,W,161229.487,A,A*40",
		errMsg: "sentence segment [0] must be \"GPGLL\" (case insensitive) but was \"GPFOO\"",
	},
	"Bad Latitude": {
		input:  "$GPGLL,bad_Latitude,N,12158.3416,W,161229.487,A,A*66",
		errMsg: "sentence segment [1] must be parsable as a float64 but was \"bad_Latitude\"",
	},
	"Bad NorthSouth": {
		input:  "$GPGLL,3723.2475,bad_NorthSouth,12158.3416,W,161229.487,A,A*2D",
		errMsg: "sentence segment [2] must be parsable as a NorthSouth but was \"bad_NorthSouth\"",
	},
	"Bad Longitude": {
		input:  "$GPGLL,3723.2475,N,bad_Longitude,W,161229.487,A,A*2B",
		errMsg: "sentence segment [3] must be parsable as a float64 but was \"bad_Longitude\"",
	},
	"Bad EastWest": {
		input:  "$GPGLL,3723.2475,N,12158.3416,bad_EastWest,161229.487,A,A*38",
		errMsg: "sentence segment [4] must be parsable as an EastWest but was \"bad_EastWest\"",
	},
	"Bad FixTime": {
		input:  "$GPGLL,3723.2475,N,12158.3416,W,bad_FixTime,A,A*01",
		errMsg: "sentence segment [5] must be parsable as a float32 but was \"bad_FixTime\"",
	},
	"Bad DataStatus": {
		input:  "$GPGLL,3723.2475,N,12158.3416,W,161229.487,bad_DataStatus,A*3C",
		errMsg: "sentence segment [6] must be parsable as a DataStatus but was \"bad_DataStatus\"",
	},
	"Bad Mode": {
		input:  "$GPGLL,3723.2475,N,12158.3416,W,161229.487,A,bad_Mode*1B",
		errMsg: "sentence segment [7] must be parsable as a Mode but was \"bad_Mode\"",
	},
}

func assertMatches(t *testing.T, title, field string, expected, actual interface{}) {
	t.Helper()
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
			assertMatches(t, title, "Latitude", expected.Latitude, actual.Latitude)
			assertMatches(t, title, "NorthSouth", expected.NorthSouth, actual.NorthSouth)
			assertMatches(t, title, "Longitude", expected.Longitude, actual.Longitude)
			assertMatches(t, title, "EastWest", expected.EastWest, actual.EastWest)
			assertMatches(t, title, "FixTime", expected.FixTime, actual.FixTime)
			assertMatches(t, title, "DataStatus", expected.DataStatus, actual.DataStatus)
			assertMatches(t, title, "Mode", expected.Mode, actual.Mode)
		})
	}
}

func TestParse_invalidChecksum(t *testing.T) {
	gpgga, err := Parse("$GPGLL,3723.2475,N,12158.3416,W,161229.487,A,A*FE")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	if gpgga != nil {
		t.Errorf("Parse result was incorrect, got: %v, want: %v", gpgga, nil)
	}

	expected := "calculated checksum value \"41\" does not match sentence-specified value of \"FE\""
	if err.Error() != expected {
		t.Errorf("err.Error() is incorrect, got: %v, want: %v", err.Error(), expected)
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

func ExampleParse() {
	sentence := "$GPGLL,3723.2475,N,12158.3416,W,161229.487,A,A*41"
	gpgll, err := Parse(sentence)
	if err != nil {
		// Handle error
	}

	fmt.Printf("%+v", gpgll)
	// Output:
	// &{Latitude:3723.2475 NorthSouth:N Longitude:12158.3416 EastWest:W FixTime:161229.48 DataStatus:A Mode:A}
}
