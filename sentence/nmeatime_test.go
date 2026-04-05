package sentence

import (
	"errors"
	"testing"
)

type nmeatimeVec struct {
	input    string
	expected NMEATime
	canon    string // expected String() output
}

var validNMEATimeInputs = []nmeatimeVec{
	{
		input:    "174800.864",
		expected: NMEATime{Hour: 17, Minute: 48, Second: 0, Millisecond: 864},
		canon:    "174800.864",
	},
	{
		input:    "161229.487",
		expected: NMEATime{Hour: 16, Minute: 12, Second: 29, Millisecond: 487},
		canon:    "161229.487",
	},
	{
		input:    "215052.603",
		expected: NMEATime{Hour: 21, Minute: 50, Second: 52, Millisecond: 603},
		canon:    "215052.603",
	},
	{
		input:    "000000.000",
		expected: NMEATime{Hour: 0, Minute: 0, Second: 0, Millisecond: 0},
		canon:    "000000.000",
	},
	{
		input:    "002454",
		expected: NMEATime{Hour: 0, Minute: 24, Second: 54, Millisecond: 0},
		canon:    "002454.000",
	},
	{
		input:    "2454",
		expected: NMEATime{Hour: 0, Minute: 24, Second: 54, Millisecond: 0},
		canon:    "002454.000",
	},
	{
		input:    "183730",
		expected: NMEATime{Hour: 18, Minute: 37, Second: 30, Millisecond: 0},
		canon:    "183730.000",
	},
	{
		input:    "30000.5",
		expected: NMEATime{Hour: 3, Minute: 0, Second: 0, Millisecond: 500},
		canon:    "030000.500",
	},
	{
		input:    "30000.05",
		expected: NMEATime{Hour: 3, Minute: 0, Second: 0, Millisecond: 50},
		canon:    "030000.050",
	},
	{
		input:    "235960.000",
		expected: NMEATime{Hour: 23, Minute: 59, Second: 60, Millisecond: 0},
		canon:    "235960.000",
	},
	{
		input:    "235959.000",
		expected: NMEATime{Hour: 23, Minute: 59, Second: 59, Millisecond: 0},
		canon:    "235959.000",
	},
	{
		input:    "174800.",
		expected: NMEATime{Hour: 17, Minute: 48, Second: 0, Millisecond: 0},
		canon:    "174800.000",
	},
	{
		input:    "174800.86499",
		expected: NMEATime{Hour: 17, Minute: 48, Second: 0, Millisecond: 864},
		canon:    "174800.864",
	},
	{
		input:    "000000.999",
		expected: NMEATime{Hour: 0, Minute: 0, Second: 0, Millisecond: 999},
		canon:    "000000.999",
	},
}

func TestParseNMEATime_validInputs(t *testing.T) {
	for _, vec := range validNMEATimeInputs {
		t.Run(vec.input, func(t *testing.T) {
			got, err := parseNMEATime(vec.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != vec.expected {
				t.Errorf("expected %+v but got %+v", vec.expected, got)
			}
		})
	}
}

func TestNMEATime_String_roundTrips(t *testing.T) {
	for _, vec := range validNMEATimeInputs {
		t.Run(vec.input, func(t *testing.T) {
			got, err := parseNMEATime(vec.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.String() != vec.canon {
				t.Errorf("String() expected %q but got %q", vec.canon, got.String())
			}
		})
	}
}

var invalidNMEATimeInputs = []string{
	"",
	"bad_FixTime",
	"174800.864.0",
	"45",
	"1748XX.864",
	"174800.XY",
	"994800.000",
	"174899.000",
	"174861.000",
	"1234567",
	"125960.000",
	"225960.000",
	"235961.000",
	" 174800.864",
	"ab4812",
	"12ab34",
	"176099.000",
	"174800.-50",
}

func TestParseNMEATime_invalidInputs(t *testing.T) {
	for _, input := range invalidNMEATimeInputs {
		t.Run(input, func(t *testing.T) {
			_, err := parseNMEATime(input)
			if err == nil {
				t.Errorf("expected an error for input %q but got nil", input)
			}
		})
	}
}

func TestParseNMEATime_invalidInputs_errorMessages(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "ab4812", want: "invalid hour"},
		{input: "12ab34", want: "invalid minute"},
		{input: "176099.000", want: "minute 60 out of range [0, 59]"},
		{input: "174800.-50", want: "millisecond -50 out of range [0, 999]"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := parseNMEATime(tt.input)
			if err == nil {
				t.Fatal("expected error")
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("error %q, want %q", got, tt.want)
			}
		})
	}
}

func TestValidateNMEATimeFields_millisecondOver999(t *testing.T) {
	err := validateNMEATimeFields(0, 0, 0, 1000)
	if err == nil {
		t.Fatal("expected error")
	}
	want := "millisecond 1000 out of range [0, 999]"
	if err.Error() != want {
		t.Errorf("error %q, want %q", err.Error(), want)
	}
}

func TestSegmentParser_AsNMEATime(t *testing.T) {
	t.Run("Good Data", func(t *testing.T) {
		p := mustParse(t)
		expected := NMEATime{Hour: 18, Minute: 37, Second: 30, Millisecond: 0}
		actual := p.AsNMEATime(1) // segment [1] = "183730"
		if actual != expected {
			t.Errorf("expected %+v but got %+v", expected, actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error but got %v", p.Err())
		}
	})

	t.Run("Empty Segment", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsNMEATime(13) // segment [13] is empty
		if actual != (NMEATime{}) {
			t.Errorf("expected zero NMEATime for empty segment but got %+v", actual)
		}
		if p.Err() != nil {
			t.Errorf("expected no error for empty segment but got %v", p.Err())
		}
	})
}

func TestSegmentParser_AsNMEATime_errors(t *testing.T) {
	t.Run("Unparsable Value", func(t *testing.T) {
		p := mustParse(t)
		p.segments[1] = "bad_time"
		actual := p.AsNMEATime(1)
		if actual != (NMEATime{}) {
			t.Errorf("expected zero NMEATime on parse failure but got %+v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for unparsable value but got nil")
		}
		expectedMsg := "sentence segment [1] must be parsable as an NMEATime but was \"bad_time\""
		if p.Err().Error() != expectedMsg {
			t.Errorf("expected error %q but got %q", expectedMsg, p.Err().Error())
		}
	})

	t.Run("Out-of-Range Index", func(t *testing.T) {
		p := mustParse(t)
		actual := p.AsNMEATime(99)
		if actual != (NMEATime{}) {
			t.Errorf("expected zero NMEATime on out-of-range index but got %+v", actual)
		}
		if p.Err() == nil {
			t.Error("expected an error for out-of-range index but got nil")
		}
	})

	t.Run("Pre-existing Error", func(t *testing.T) {
		p := mustParse(t)
		p.AsNMEATime(99) // sets error
		firstErr := p.Err()
		p.AsNMEATime(1) // should exit early, leaving error unchanged
		if !errors.Is(p.Err(), firstErr) {
			t.Errorf("expected error to remain unchanged but it changed to %v", p.Err())
		}
	})
}
