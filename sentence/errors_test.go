package sentence

import "testing"

func TestParsingError_Error(t *testing.T) {
	err := ParsingError{Segment: 3, Message: "is out of range"}
	expected := "sentence segment [3] is out of range"
	if err.Error() != expected {
		t.Errorf("expected %q but was %q", expected, err.Error())
	}
}
