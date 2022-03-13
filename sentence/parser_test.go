package sentence

import (
	"fmt"
	"testing"

	"gopkg.in/mab-go/nmea.v0/sentence/testhelp"
)

type testParseInput struct {
	Name, Sentence, ActualChecksum, AdvertisedChecksum, ErrMsg string
}

func mapParseInput(title string, input map[string]string) interface{} {
	return testParseInput{
		Name:               title,
		Sentence:           input["Sentence"],
		ActualChecksum:     input["ActualChecksum"],
		AdvertisedChecksum: input["AdvertisedChecksum"],
		ErrMsg:             input["ErrMsg"],
	}
}

func sortParseInput(result []interface{}, i, j int) bool {
	return result[i].(testParseInput).Name < result[j].(testParseInput).Name
}

func TestSegmentParser_Parse(t *testing.T) {
	// Test with good data
	goodData := testhelp.ReadTestData("good/sentences", mapParseInput, sortParseInput)
	for _, data := range goodData {
		d := data.(testParseInput)

		t.Run(fmt.Sprintf("Good Data/%s", d.Name), func(t *testing.T) {
			parser := &SegmentParser{}
			err := parser.Parse(d.Sentence) // Unit under test
			if err != nil {
				t.Errorf("segment parsing failed: %v", err)
			}
		})
	}

	// Test with invalid checksums
	badData := testhelp.ReadTestData("bad/invalid-checksums", mapParseInput, sortParseInput)
	for _, data := range badData {
		d := data.(testParseInput)

		t.Run(fmt.Sprintf("Bad Data/Invalid Checksums/%s", d.Name), func(t *testing.T) {
			parser := &SegmentParser{}
			err := parser.Parse(d.Sentence) // Unit under test
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

func TestSegmentParser_Err(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_AsFloat32(t *testing.T) {
	sentence := "$GPGGA,183730,3907.356,N,12102.482,W,1,05,1.6,646.4,M,-24.1,M,,*75"
	parser := &SegmentParser{}
	err := parser.Parse(sentence)
	if err != nil {
		t.Errorf("segment parsing failed: %v", err)
	}

	// Test with a float32
	t.Run("Good Data", func(t *testing.T) {
		expected := float32(646.4)
		actual := parser.AsFloat32(9) // Unit under test
		if actual != expected {
			t.Errorf("expected result to be %v but was %v", expected, actual)
		}
	})

	// Test with out-of-range index
	// t.Run("Index Out of Range", func(t *testing.T) {
	//	v := parser.AsFloat32(99)
	//	if parser.Err() == nil {
	//		t.Errorf("!!! %v", err)
	//	}
	// })
}

func TestSegmentParser_AsFloat64(t *testing.T) {
	/*
		"$GPGGA,183730,3907.356,N,12102.482,W,1,05,1.6,646.4,M,-24.1,M,,*75"
	*/

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

func TestSegmentParser_AsInt64(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_RequireString(t *testing.T) {
	t.Skip()
}

func TestSegmentParser_RequireStrings(t *testing.T) {
	t.Skip()
}
