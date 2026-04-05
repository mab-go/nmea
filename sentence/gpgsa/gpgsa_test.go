package gpgsa

import (
	"fmt"
	"testing"
)

type testVec struct {
	input    string
	expected GPGSA
	errMsg   string
}

var goodTestData = map[string]testVec{
	// Source: AMOD_AGL3080_20121104_134730.txt, line 4
	// Scenario: 3D fix, 10 active satellites, DOP 1.8/0.8/1.6
	"AMOD AGL3080 [3D/full]": {
		input: "$GPGSA,A,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*3F",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       Fix3D,
			PRNs:          [12]int8{3, 22, 6, 19, 11, 14, 32, 1, 28, 18, 0, 0},
			PDOP:          1.8,
			HDOP:          0.8,
			VDOP:          1.6,
		},
	},
	// Source: AMOD_AGL3080_20121104_134730.txt, line 50
	// Scenario: 3D fix, 8 active satellites, some slots empty
	"AMOD AGL3080 [3D/sparse]": {
		input: "$GPGSA,A,3,03,22,06,19,14,32,28,18,,,,,2.1,1.0,1.8*33",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       Fix3D,
			PRNs:          [12]int8{3, 22, 6, 19, 14, 32, 28, 18, 0, 0, 0, 0},
			PDOP:          2.1,
			HDOP:          1.0,
			VDOP:          1.8,
		},
	},
	// Source: AMOD_AGL3080_20121104_134730.txt, line 186
	// Scenario: 2D fix, 3 active satellites, noticeably degraded DOP
	"AMOD AGL3080 [2D/degraded]": {
		input: "$GPGSA,A,2,03,32,18,,,,,,,,,,3.1,2.9,1.0*30",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       Fix2D,
			PRNs:          [12]int8{3, 32, 18, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			PDOP:          3.1,
			HDOP:          2.9,
			VDOP:          1.0,
		},
	},
	// Source: AMOD_AGL3080_20121104_134730.txt, line 190
	// Scenario: Lost signal — all 12 PRN slots empty, all DOP fields at sentinel value 50.0
	// Note: hardware reports FixMode 2 (Fix2D) during this state, not FixMode 1 (NoFix)
	"AMOD AGL3080 [no-fix/sentinel]": {
		input: "$GPGSA,A,2,,,,,,,,,,,,,50.0,50.0,50.0*06",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       Fix2D,
			PRNs:          [12]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			PDOP:          50.0,
			HDOP:          50.0,
			VDOP:          50.0,
		},
	},
	// Source: AMOD_AGL3080_20121104_134730.txt, line 255
	// Scenario: Partial re-acquisition — 3 satellites, 2D fix, partial DOP recovery
	"AMOD AGL3080 [partial-reacq]": {
		input: "$GPGSA,A,2,03,06,32,,,,,,,,,,50.0,50.0,1.0*36",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       Fix2D,
			PRNs:          [12]int8{3, 6, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			PDOP:          50.0,
			HDOP:          50.0,
			VDOP:          1.0,
		},
	},
	// Scenario: Manual satellite selection, 3D fix
	"Manual selection [3D]": {
		input: "$GPGSA,M,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*33",
		expected: GPGSA{
			SelectionMode: ManualSelectionMode,
			FixMode:       Fix3D,
			PRNs:          [12]int8{3, 22, 6, 19, 11, 14, 32, 1, 28, 18, 0, 0},
			PDOP:          1.8,
			HDOP:          0.8,
			VDOP:          1.6,
		},
	},
	// Scenario: NoFix (NMEA wire value 1), no satellites tracked, sentinel DOP values
	"No fix [NoFix/sentinel]": {
		input: "$GPGSA,A,1,,,,,,,,,,,,,99.9,99.9,99.9*09",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       NoFix,
			PRNs:          [12]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			PDOP:          99.9,
			HDOP:          99.9,
			VDOP:          99.9,
		},
	},
	// Scenario: All 12 PRN slots populated, 3D fix
	"All 12 PRNs [3D/full]": {
		input: "$GPGSA,A,3,01,02,03,04,05,06,07,08,09,10,11,12,2.5,1.5,2.0*30",
		expected: GPGSA{
			SelectionMode: AutomaticSelectionMode,
			FixMode:       Fix3D,
			PRNs:          [12]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			PDOP:          2.5,
			HDOP:          1.5,
			VDOP:          2.0,
		},
	},
}

var badTestData = map[string]testVec{
	"Bad SentenceType": {
		input:  "$GAAAA,A,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*3A",
		errMsg: "sentence segment [0] must be \"GPGSA\" (case insensitive) but was \"GAAAA\"",
	},
	"Bad SelectionMode": {
		input:  "$GPGSA,bad_SelectionMode,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*25",
		errMsg: "sentence segment [1] must be parsable as a SelectionMode but was \"bad_SelectionMode\"",
	},
	"Bad FixMode (Wrong Type)": {
		input:  "$GPGSA,A,bad_FixMode,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*40",
		errMsg: "sentence segment [2] must be parsable as a FixMode but was \"bad_FixMode\"",
	},
	"Bad FixMode (Out of Range)": {
		input:  "$GPGSA,A,9,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*35",
		errMsg: "sentence segment [2] must be parsable as a FixMode but was \"9\"",
	},
	"Bad FixMode (Negative Number)": {
		input:  "$GPGSA,A,-1,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*10",
		errMsg: "sentence segment [2] must be parsable as a FixMode but was \"-1\"",
	},
	"Bad PRN": {
		input:  "$GPGSA,A,3,bad_PRN,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*48",
		errMsg: "sentence segment [3] must be parsable as an int8 but was \"bad_PRN\"",
	},
	"Bad PDOP": {
		input:  "$GPGSA,A,3,03,22,06,19,11,14,32,01,28,18,,,bad_PDOP,0.8,1.6*2B",
		errMsg: "sentence segment [15] must be parsable as a float32 but was \"bad_PDOP\"",
	},
	"Bad HDOP": {
		input:  "$GPGSA,A,3,03,22,06,19,11,14,32,01,28,18,,,1.8,bad_HDOP,1.6*32",
		errMsg: "sentence segment [16] must be parsable as a float32 but was \"bad_HDOP\"",
	},
	"Bad VDOP": {
		input:  "$GPGSA,A,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,bad_VDOP*23",
		errMsg: "sentence segment [17] must be parsable as a float32 but was \"bad_VDOP\"",
	},
}

func assertMatches(t *testing.T, title, field string, expected, actual any) {
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
				t.Fatalf("error creating GPGSA from NMEA input \"%v\": %v", title, err)
			}

			expected := vec.expected
			assertMatches(t, title, "SelectionMode", expected.SelectionMode, actual.SelectionMode)
			assertMatches(t, title, "FixMode", expected.FixMode, actual.FixMode)
			assertMatches(t, title, "PRNs", expected.PRNs, actual.PRNs)
			assertMatches(t, title, "PDOP", expected.PDOP, actual.PDOP)
			assertMatches(t, title, "HDOP", expected.HDOP, actual.HDOP)
			assertMatches(t, title, "VDOP", expected.VDOP, actual.VDOP)
		})
	}
}

func TestParse_invalidChecksum(t *testing.T) {
	gpgsa, err := Parse("$GPGSA,A,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*00")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	if gpgsa != nil {
		t.Errorf("result should have been <nil> but was %v", gpgsa)
	}

	expected := "calculated checksum value \"3F\" does not match sentence-specified value of \"00\""
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestParse_badSegments(t *testing.T) {
	for title, vec := range badTestData {
		t.Run(title, func(t *testing.T) {
			gpgsa, err := Parse(vec.input)
			if err == nil {
				t.Fatalf("parsing succeeded (but should not have) for test sentence %q", title)
			}

			if gpgsa != nil {
				t.Fatalf("result should have been <nil> but was %v for test sentence %q", gpgsa, title)
			}

			if err.Error() != vec.errMsg {
				t.Fatalf("error message should have been '%v' but was '%v' for test sentence %q", vec.errMsg, err.Error(), title)
			}
		})
	}
}

func TestGPGSA_GetSentenceType(t *testing.T) {
	gpgsa := &GPGSA{}
	if st := gpgsa.GetSentenceType(); st != "GPGSA" {
		t.Errorf("GetSentenceType() should have returned \"GPGSA\" but returned \"%v\"", st)
	}
}

func ExampleParse() {
	sentence := "$GPGSA,A,3,03,22,06,19,11,14,32,01,28,18,,,1.8,0.8,1.6*3F"
	gpgsa, err := Parse(sentence)
	_ = err

	fmt.Printf("%+v", gpgsa)
	// Output:
	// &{SelectionMode:A FixMode:3 PRNs:[3 22 6 19 11 14 32 1 28 18 0 0] PDOP:1.8 HDOP:0.8 VDOP:1.6}
}
