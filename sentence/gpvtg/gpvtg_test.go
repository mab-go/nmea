package gpvtg

import (
	"fmt"
	"testing"
)

type testVec struct {
	input    string
	expected GPVTG
	errMsg   string
}

var goodTestData = map[string]testVec{
	"Without Mode": {
		input: "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K*48",
		expected: GPVTG{
			TrueTrack:  54.7,
			MagTrack:   34.4,
			SpeedKnots: 5.5,
			SpeedKmh:   10.2,
		},
	},
	"With Mode A": {
		input: "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K,A*25",
		expected: GPVTG{
			TrueTrack:  54.7,
			MagTrack:   34.4,
			SpeedKnots: 5.5,
			SpeedKmh:   10.2,
			Mode:       Autonomous,
		},
	},
	"With Mode D": {
		input: "$GPVTG,309.62,T,309.65,M,0.000,N,0.000,K,D*21",
		expected: GPVTG{
			TrueTrack:  309.62,
			MagTrack:   309.65,
			SpeedKnots: 0.0,
			SpeedKmh:   0.0,
			Mode:       Differential,
		},
	},
	"Empty values": {
		input: "$GPVTG,,T,,M,,N,,K*4E",
		expected: GPVTG{
			TrueTrack:  0,
			MagTrack:   0,
			SpeedKnots: 0,
			SpeedKmh:   0,
		},
	},
}

var badTestData = map[string]testVec{
	"Bad SentenceType": {
		input:  "$GAAAA,054.7,T,034.4,M,005.5,N,010.2,K*5D",
		errMsg: "sentence segment [0] must be \"GPVTG\" (case insensitive) but was \"GAAAA\"",
	},
	"Bad TrueTrack": {
		input:  "$GPVTG,bad,T,034.4,M,005.5,N,010.2,K*07",
		errMsg: "sentence segment [1] must be parsable as a float64 but was \"bad\"",
	},
	"Bad T indicator": {
		input:  "$GPVTG,054.7,X,034.4,M,005.5,N,010.2,K*44",
		errMsg: "sentence segment [2] must be \"T\" (case insensitive) but was \"X\"",
	},
	"Bad MagTrack": {
		input:  "$GPVTG,054.7,T,bad,M,005.5,N,010.2,K*02",
		errMsg: "sentence segment [3] must be parsable as a float64 but was \"bad\"",
	},
	"Bad M indicator": {
		input:  "$GPVTG,054.7,T,034.4,X,005.5,N,010.2,K*5D",
		errMsg: "sentence segment [4] must be \"M\" (case insensitive) but was \"X\"",
	},
	"Bad SpeedKnots": {
		input:  "$GPVTG,054.7,T,034.4,M,bad,N,010.2,K*01",
		errMsg: "sentence segment [5] must be parsable as a float64 but was \"bad\"",
	},
	"Bad N indicator": {
		input:  "$GPVTG,054.7,T,034.4,M,005.5,X,010.2,K*5E",
		errMsg: "sentence segment [6] must be \"N\" (case insensitive) but was \"X\"",
	},
	"Bad SpeedKmh": {
		input:  "$GPVTG,054.7,T,034.4,M,005.5,N,bad,K*02",
		errMsg: "sentence segment [7] must be parsable as a float64 but was \"bad\"",
	},
	"Bad K indicator": {
		input:  "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,X*5B",
		errMsg: "sentence segment [8] must be \"K\" (case insensitive) but was \"X\"",
	},
	"Bad Mode": {
		input:  "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K,X*3C",
		errMsg: "sentence segment [9] must be parsable as a Mode but was \"X\"",
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
				t.Fatalf("error creating GPVTG from NMEA input \"%v\": %v", title, err)
			}

			expected := vec.expected
			assertMatches(t, title, "TrueTrack", expected.TrueTrack, actual.TrueTrack)
			assertMatches(t, title, "MagTrack", expected.MagTrack, actual.MagTrack)
			assertMatches(t, title, "SpeedKnots", expected.SpeedKnots, actual.SpeedKnots)
			assertMatches(t, title, "SpeedKmh", expected.SpeedKmh, actual.SpeedKmh)
			assertMatches(t, title, "Mode", expected.Mode, actual.Mode)
		})
	}
}

func TestParse_invalidChecksum(t *testing.T) {
	gpvtg, err := Parse("$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K*42")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	if gpvtg != nil {
		t.Errorf("result should have been <nil> but was %v", gpvtg)
	}

	expected := "calculated checksum value \"48\" does not match sentence-specified value of \"42\""
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestParse_badSegments(t *testing.T) {
	for title, vec := range badTestData {
		t.Run(title, func(t *testing.T) {
			gpvtg, err := Parse(vec.input)
			if err == nil {
				t.Fatalf("parsing succeeded (but should not have) for test sentence %q", title)
			}

			if gpvtg != nil {
				t.Fatalf("result should have been <nil> but was %v for test sentence %q", gpvtg, title)
			}

			if err.Error() != vec.errMsg {
				t.Fatalf("error message should have been '%v' but was '%v' for test sentence %q", vec.errMsg, err.Error(), title)
			}
		})
	}
}

func TestGPVTG_GetSentenceType(t *testing.T) {
	gpvtg := &GPVTG{}
	if st := gpvtg.GetSentenceType(); st != "GPVTG" {
		t.Errorf("GetSentenceType() should have returned \"GPVTG\" but returned \"%v\"", st)
	}
}

func ExampleParse() {
	s := "$GPVTG,054.7,T,034.4,M,005.5,N,010.2,K,A*25"
	gpvtg, err := Parse(s)
	_ = err

	fmt.Printf("%+v", gpvtg)
	// Output:
	// &{TrueTrack:54.7 MagTrack:34.4 SpeedKnots:5.5 SpeedKmh:10.2 Mode:A}
}
