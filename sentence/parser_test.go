package sentence

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mab-go/nmea/sentence/testhelp"
)

// referenceSentence is a valid GPGGA sentence used as the base for most
// SegmentParser tests. Segment indices:
//
//		[0]=GPGGA  [1]=183730  [2]=3907.356  [3]=N    [4]=12102.482
//		[5]=W      [6]=1       [7]=05        [8]=1.6  [9]=646.4
//	 [10]=M     [11]=-24.1  [12]=M        [13]=(empty)
const referenceSentence = "$GPGGA,183730,3907.356,N,12102.482,W,1,05,1.6,646.4,M,-24.1,M,,*75"

// mustParse returns a SegmentParser that has successfully parsed referenceSentence.
func mustParse(t *testing.T) *SegmentParser {
	t.Helper()
	p := &SegmentParser{}
	if err := p.Parse(referenceSentence); err != nil {
		t.Fatalf("failed to parse reference sentence: %v", err)
	}

	return p
}

type parseTestData struct {
	Title, Sentence, ActualChecksum, AdvertisedChecksum, ErrMsg string
}

func mapParseTestData(title string, input map[string]interface{}) interface{} {
	return parseTestData{
		Title:              title,
		Sentence:           testhelp.EnsureString(input["Sentence"]),
		ActualChecksum:     testhelp.EnsureString(input["ActualChecksum"]),
		AdvertisedChecksum: testhelp.OptString(input["AdvertisedChecksum"]),
		ErrMsg:             testhelp.OptString(input["ErrMsg"]),
	}
}

func sortParseTestData(result []interface{}, i, j int) bool {
	return result[i].(parseTestData).Title < result[j].(parseTestData).Title
}

func TestSegmentParser_Parse_goodData(t *testing.T) {
	for _, data := range testhelp.ReadTestData("good-data", mapParseTestData, sortParseTestData) {
		d := data.(parseTestData)

		t.Run(d.Title, func(t *testing.T) {
			parser := &SegmentParser{}
			err := parser.Parse(d.Sentence) // Unit under test
			if err != nil {
				t.Errorf("segment parsing failed: %v", err)
			}
		})
	}
}

func TestSegmentParser_Parse_invalidChecksums(t *testing.T) {
	for _, data := range testhelp.ReadTestData("bad-invalid-checksums", mapParseTestData, sortParseTestData) {
		d := data.(parseTestData)

		t.Run(d.Title, func(t *testing.T) {
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
	t.Run("No Error", func(t *testing.T) {
		p := mustParse(t)
		if p.Err() != nil {
			t.Errorf("expected Err() to be nil but was %v", p.Err())
		}
	})

	t.Run("After Failed Accessor", func(t *testing.T) {
		p := mustParse(t)
		p.AsFloat32(99) // out-of-range index sets an error
		if p.Err() == nil {
			t.Error("expected Err() to be non-nil after out-of-range access")
		}
	})
}

func TestSegmentParser_AsFloat32(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		expected := float32(646.4)
		actual := p.AsFloat32(9)
		if actual != expected {
			t.Errorf("expected %v but was %v", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsFloat32(13) // segment [13] is empty
		if actual != 0 {
			t.Errorf("expected 0 for empty segment but was %v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})

	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[2] = "not_a_float"
		actual := p.AsFloat32(2)
		if actual != 0 {
			t.Errorf("expected 0 on parse failure but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsFloat32(99)
		if actual != 0 {
			t.Errorf("expected 0 on out-of-range index but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsFloat32(99) // sets error
		firstErr := p.Err()
		p.AsFloat32(9) // should exit early, leaving error unchanged
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_AsFloat64(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		expected := 3907.356
		actual := p.AsFloat64(2)
		if actual != expected {
			t.Errorf("expected %v but was %v", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsFloat64(13)
		if actual != 0 {
			t.Errorf("expected 0 for empty segment but was %v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})

	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[2] = "not_a_float"
		actual := p.AsFloat64(2)
		if actual != 0 {
			t.Errorf("expected 0 on parse failure but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsFloat64(99)
		if actual != 0 {
			t.Errorf("expected 0 on out-of-range index but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsFloat64(99)
		firstErr := p.Err()
		p.AsFloat64(2)
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_AsInt8(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		expected := int8(1)
		actual := p.AsInt8(6) // segment [6] = "1"
		if actual != expected {
			t.Errorf("expected %v but was %v", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt8(13)
		if actual != 0 {
			t.Errorf("expected 0 for empty segment but was %v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})

	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[6] = "not_an_int"
		actual := p.AsInt8(6)
		if actual != 0 {
			t.Errorf("expected 0 on parse failure but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt8(99)
		if actual != 0 {
			t.Errorf("expected 0 on out-of-range index but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsInt8(99)
		firstErr := p.Err()
		p.AsInt8(6)
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_AsInt8InRange(t *testing.T) {
	t.Run("Good Data In Range", func(t *testing.T) {
		p := mustParse(t)
		expected := int8(1)
		actual := p.AsInt8InRange(6, 0, 9) // segment [6] = "1", range [0,9]
		if actual != expected {
			t.Errorf("expected %v but was %v", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Value Below Range", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt8InRange(6, 5, 9) // "1" is below lower bound 5
		if actual != 0 {
			t.Errorf("expected 0 when below range but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error when value is below range but got nil")
		}
	})

	t.Run("Value Above Range", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt8InRange(6, 0, 0) // "1" is above upper bound 0
		if actual != 0 {
			t.Errorf("expected 0 when above range but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error when value is above range but got nil")
		}
	})
}

func TestSegmentParser_AsInt8InRange_errors(t *testing.T) {
	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt8InRange(99, 0, 9)
		if actual != 0 {
			t.Errorf("expected 0 on out-of-range index but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[6] = "not_an_int"
		actual := p.AsInt8InRange(6, 0, 9)
		if actual != 0 {
			t.Errorf("expected 0 on unparsable value but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsInt8InRange(99, 0, 9)
		firstErr := p.Err()
		p.AsInt8InRange(6, 0, 9)
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_AsInt16(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt16(7) // segment [7] = "05"
		if actual != 5 {
			t.Errorf("expected 5 but was %v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt16(13)
		if actual != 0 {
			t.Errorf("expected 0 for empty segment but was %v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})

	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[7] = "not_an_int"
		actual := p.AsInt16(7)
		if actual != 0 {
			t.Errorf("expected 0 on parse failure but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt16(99)
		if actual != 0 {
			t.Errorf("expected 0 on out-of-range index but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsInt16(99)
		firstErr := p.Err()
		p.AsInt16(7)
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_AsInt32(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		expected := int32(183730)
		actual := p.AsInt32(1) // segment [1] = "183730"
		if actual != expected {
			t.Errorf("expected %v but was %v", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt32(13)
		if actual != 0 {
			t.Errorf("expected 0 for empty segment but was %v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})

	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[1] = "not_an_int"
		actual := p.AsInt32(1)
		if actual != 0 {
			t.Errorf("expected 0 on parse failure but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsInt32(99)
		if actual != 0 {
			t.Errorf("expected 0 on out-of-range index but was %v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsInt32(99)
		firstErr := p.Err()
		p.AsInt32(1)
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_AsString(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		expected := "GPGGA"
		actual := p.AsString(0)
		if actual != expected {
			t.Errorf("expected %q but was %q", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsString(13)
		if actual != "" {
			t.Errorf("expected empty string for empty segment but was %q", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsString(99)
		if actual != "" {
			t.Errorf("expected empty string on out-of-range index but was %q", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsString(99)
		firstErr := p.Err()
		p.AsString(0)
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_RequireString(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireString(0, "GPGGA")
		if actual != "GPGGA" {
			t.Errorf("expected %q but was %q", "GPGGA", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Case-Insensitive Match", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireString(0, "gpgga")
		if actual != "GPGGA" {
			t.Errorf("expected %q but was %q", "GPGGA", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error on case-insensitive match but got %v", p.Err())
		}
	})

	t.Run("Mismatch", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireString(0, "GPGLL")
		if actual != "" {
			t.Errorf("expected empty string on mismatch but was %q", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error on mismatch but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireString(99, "GPGGA")
		if actual != "" {
			t.Errorf("expected empty string on out-of-range index but was %q", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.RequireString(99, "GPGGA")
		firstErr := p.Err()
		p.RequireString(0, "GPGGA")
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}

func TestSegmentParser_RequireStrings(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireStrings(0, []string{"GPGLL", "GPGGA", "GPRMC"})
		if actual != "GPGGA" {
			t.Errorf("expected %q but was %q", "GPGGA", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Case-Insensitive Match", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireStrings(0, []string{"gpgga"})
		if actual != "GPGGA" {
			t.Errorf("expected %q but was %q", "GPGGA", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error on case-insensitive match but got %v", p.Err())
		}
	})

	t.Run("No Match", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireStrings(0, []string{"GPGLL", "GPRMC"})
		if actual != "" {
			t.Errorf("expected empty string when no match but was %q", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error when no match but got nil")
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.RequireStrings(99, []string{"GPGGA"})
		if actual != "" {
			t.Errorf("expected empty string on out-of-range index but was %q", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.RequireStrings(99, []string{"GPGGA"})
		firstErr := p.Err()
		p.RequireStrings(0, []string{"GPGGA"})
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}
