package sentence

import (
	"fmt"
	"testing"

	"gopkg.in/mab-go/nmea.v0/sentence/testhelp"
)

// --- Test Functions ----------------------------------------------------------

func TestSegmentParser_Parse(t *testing.T) {
	// Test with good data
	for _, d := range testhelp.ReadTestData("good/sentences") {
		t.Run(fmt.Sprintf("Good Data/%s", d.Name), func(t *testing.T) {
			parser := &SegmentParser{}
			err := parser.Parse(d.Sentence)
			if err != nil {
				t.Errorf("segment parsing failed: %v", err)
			}
		})
	}

	// Test with invalid checksums
	for _, d := range testhelp.ReadTestData("bad/invalid-checksums") {
		t.Run(fmt.Sprintf("Bad Data/Invalid Checksums/%s", d.Name), func(t *testing.T) {
			parser := &SegmentParser{}
			err := parser.Parse(d.Sentence)
			if err == nil {
				t.Error("segment parsing succeeded (but should not have)")
			}

			expectedMsg := fmt.Sprintf(
				"calculated checksum value \"%s\" does not match sentence-specified value of \"%s\"",
				d.ActualChecksum,
				d.AdvertisedChecksum)
			if err.Error() != expectedMsg {
				t.Errorf("error message should have been '%v' but was '%v'", expectedMsg, err.Error())
			}
		})
	}
}

//func TestSegmentParser_Parse_goodData(t *testing.T) {
//	var testData []testSentence
//	readTestData("good-with-checksums.json", &testData)
//
//	for i, s := range testData {
//		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
//			parser := &SegmentParser{}
//			err := parser.Parse(s.Sentence)
//			if err != nil {
//				t.Errorf("segment parsing failed for test sentence [%v]: %v", i, err)
//			}
//		})
//	}
//}
//
//func TestSegmentParser_Parse_badData(t *testing.T) {
//	var testData []testSentence
//	readTestData("bad-invalid-checksums.json", &testData)
//
//	for i, s := range testData {
//		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
//			parser := &SegmentParser{}
//			err := parser.Parse(s.Sentence)
//			if err == nil {
//				t.Errorf("segment parsing succeeded (but should not have) for test sentence [%v]", i)
//			}
//
//			if err.Error() != s.ErrMsg {
//				t.Errorf("error message should have been '%v' but was '%v' for test sentence [%v]", s.ErrMsg, err.Error(), i)
//			}
//		})
//	}
//}

func TestSegmentParser_Err(t *testing.T) {
	t.Skip()
}

//func TestSegmentParser_AsFloat32_goodData(t *testing.T) {
//	sentence := "$GPGGA,002454.123,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*61"
//	segments := &SegmentParser{}
//	err := segments.Parse(sentence)
//	if err != nil {
//		t.Errorf("segment parsing failed for test sentence: %v", err)
//	}
//
//	expected := float32(2454.123)
//	actual := segments.AsFloat32(1)
//
//	if segments.Err() != nil {
//		t.Errorf("encountered error after calling AsFloat32: %v", segments.Err())
//	}
//
//	if actual != expected {
//		t.Errorf("return value should have been %v but was %v", expected, actual)
//	}
//}
//
//func TestSegmentParser_AsFloat32_badData(t *testing.T) {
//	sentence := "$GPGGA,bad_FixTime,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*22"
//	segments := &SegmentParser{}
//	err := segments.Parse(sentence)
//	if err != nil {
//		t.Errorf("segment parsing failed for test sentence: %v", err)
//	}
//
//	segments.AsFloat32(1)
//
//	expectedErr := "sentence segment [1] must be parsable as a float32 but was \"bad_FixTime\""
//	if segments.Err() == nil {
//		t.Error("should have encountered error after calling AsFloat32 with bad data but did not")
//	} else if err := segments.Err().Error(); err != expectedErr {
//		t.Errorf("error message should have been '%v' but was '%v'", expectedErr, err)
//	}
//}

func TestSegmentParser_AsFloat64(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_AsInt8(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_AsInt8InRange(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_AsInt16(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_AsInt32(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_RequireString(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_RequireStrings(t *testing.T) {
	t.Skip()
}
