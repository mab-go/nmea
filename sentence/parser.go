package sentence

import (
	"fmt"
	"strconv"
	"strings"
)

// --- Public ------------------------------------------------------------------

// SegmentParser provides functionality for parsing individual segments of an
// NMEA sentence.
type SegmentParser struct {
	sentence string
	segments []string
	err      *ParsingError
}

// Parse parses the specified NMEA sentence into a series of sentence segments.
func (p *SegmentParser) Parse(s string) (error) {
	if err := VerifyChecksum(s); err != nil {
		return err
	}

	// Strip the first character ("$") and the last three characters (the checksum),
	// and then split the remaining string on a comma (",").
	segments := strings.Split(s[1:len(s)-3], ",")

	p.sentence = s
	p.segments = segments

	return nil
}

// Err returns a SegmentParser's *ParsingError.
func (p *SegmentParser) Err() *ParsingError {
	return p.err
}

// AsFloat32 parses the sentence segment at the specified index as a float32 value.
// If p.Err() is not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsFloat32(i int8) float32 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	} else if val, err := strconv.ParseFloat(p.segments[i], 32); err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a float32 but was \"%s\"", p.segments[i]),
		}
		return 0
	} else {
		return float32(val)
	}
}

// AsFloat64 parses the sentence segment at the specified index as a float64 value.
// If p.Err() is not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsFloat64(i int8) float64 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	} else if val, err := strconv.ParseFloat(p.segments[i], 64); err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a float64 but was \"%s\"", p.segments[i]),
		}
		return 0
	} else {
		return val
	}
}

// AsInt8 parses the sentence segment at the specified index as an int8 value.
// If p.Err() is not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt8(i int8) int8 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	} else if val, err := strconv.ParseInt(p.segments[i], 10, 8); err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an int8 but was \"%s\"", p.segments[i]),
		}
		return 0
	} else {
		return int8(val)
	}
}

// AsInt8InRange parses the sentence segment at the specified index as an int8
// value and ensures that it matches one of the required values in the range from
// l to u (lower and upper bound inclusive). If p.Err() is not nil, this function
// returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt8InRange(i int8, l int8, u int8) int8 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	val := p.AsInt8(i)
	if p.err != nil {
		return 0
	}

	if val < l || val > u {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be within range [%d, %d] but was %s", l, u, p.segments[i]),
		}
		return 0
	}

	return val
}

// AsInt16 parses the sentence segment at the specified index as an int32 value.
// If p.Err() is not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt16(i int8) int16 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	} else if val, err := strconv.ParseInt(p.segments[i], 10, 16); err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an int16 but was \"%s\"", p.segments[i]),
		}
		return 0
	} else {
		return int16(val)
	}
}

// AsInt32 parses the sentence segment at the specified index as an int32 value.
// If p.Err() is not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt32(i int8) int32 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	} else if val, err := strconv.ParseInt(p.segments[i], 10, 32); err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an int32 but was \"%s\"", p.segments[i]),
		}
		return 0
	} else {
		return int32(val)
	}
}

// RequireString parses the sentence segment at the specified index as a string
// value and ensures that it matches the required value s (case insensitive).
// If p.Err() is not nil, this function returns an empty string and leaves the
// error unchanged.
func (p *SegmentParser) RequireString(i int8, s string) string {
	if p.checkInRange(i); p.err != nil {
		return ""
	}

	if strings.ToUpper(s) != strings.ToUpper(p.segments[i]) {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be \"%s\" (case insensitive) but was \"%s\"", s, p.segments[i]),
		}
		return ""
	}

	return p.segments[i]
}

// RequireStrings parses the sentence segment at the specified index as a string
// value and ensures that it matches one of the required values in s (case
// insensitive). If p.Err() is not nil, this function returns an empty string and
// leaves the error unchanged.
func (p *SegmentParser) RequireStrings(i int8, s []string) string {
	if p.checkInRange(i); p.err != nil {
		return ""
	}

	for _, st := range s {
		if strings.ToUpper(st) == strings.ToUpper(p.segments[i]) {
			return p.segments[i] // The value matches
		}
	}

	// We didn't find a match
	p.err = &ParsingError{
		Segment: i,
		Message: fmt.Sprintf("must be one of %v (case insensitive) but was \"%s\"", s, p.segments[i]),
	}

	return ""
}

// --- Private -----------------------------------------------------------------

func (p *SegmentParser) checkInRange(i int8) {
	if p.err != nil {
		return // There's already an error; exit early.
	}

	if int8(len(p.segments)-1) < i {
		p.err = &ParsingError{Segment: i, Message: "is out of range"}
	}
}
