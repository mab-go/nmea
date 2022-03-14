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
				t.Errorf("error creating GPGLL from NMEA input \"%v\": %v", title, err)
				return
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
