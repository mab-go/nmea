package sentence

import (
	"fmt"
	"strconv"
	"strings"
)

// --- Public ------------------------------------------------------------------

// SegmentParser provides functionality for parsing individual segments of an NMEA sentence.
type SegmentParser struct {
	sentence string
	segments []string
	err      error
}

// Parse parses the specified NMEA sentence into a series of sentence segments.
func (p *SegmentParser) Parse(s string) error {
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

// Err returns a SegmentParser's error value.
func (p *SegmentParser) Err() error {
	return p.err
}

// AsFloat32 parses the sentence segment at the specified index as a float32 value. If p.Err() is
// not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsFloat32(i int8) float32 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	}

	val, err := strconv.ParseFloat(p.segments[i], 32)
	if err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a float32 but was \"%s\"", p.segments[i]),
		}

		return 0
	}

	return float32(val)
}

// AsFloat64 parses the sentence segment at the specified index as a float64 value. If p.Err() is
// not nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsFloat64(i int8) float64 {
	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	}

	val, err := strconv.ParseFloat(p.segments[i], 64)
	if err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a float64 but was \"%s\"", p.segments[i]),
		}

		return 0
	}

	return val
}

// AsInt8 parses the sentence segment at the specified index as an int8 value. If p.Err() is not
// nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt8(i int8) int8 {
	if val := p.asInt(i, 8); val == 0 {
		return int8(val.(int))
	}

	return p.asInt(i, 8).(int8)
}

// AsInt8InRange parses the sentence segment at the specified index as an int8 value and ensures
// that it matches one of the required values in the range from l to u (lower and upper bound
// inclusive). If p.Err() is not nil, this function returns 0 and leaves the error unchanged.
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

// AsInt16 parses the sentence segment at the specified index as an int32 value. If p.Err() is not
// nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt16(i int8) int16 {
	if val := p.asInt(i, 16); val == 0 {
		return int16(val.(int))
	}

	return p.asInt(i, 16).(int16)
}

// AsInt32 parses the sentence segment at the specified index as an int32 value. If p.Err() is not
// nil, this function returns 0 and leaves the error unchanged.
func (p *SegmentParser) AsInt32(i int8) int32 {
	if val := p.asInt(i, 32); val == 0 {
		return int32(val.(int))
	}

	return p.asInt(i, 32).(int32)
}

// AsNMEATime parses the sentence segment at the specified index as an NMEATime value. If p.Err()
// is not nil, this function returns NMEATime{} and leaves the error unchanged. An empty segment
// returns NMEATime{} with no error.
func (p *SegmentParser) AsNMEATime(i int8) NMEATime {
	if p.checkInRange(i); p.err != nil {
		return NMEATime{}
	}

	if p.segments[i] == "" {
		return NMEATime{}
	}

	t, err := parseNMEATime(p.segments[i])
	if err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an NMEATime but was \"%s\"", p.segments[i]),
		}

		return NMEATime{}
	}

	return t
}

// AsString parses the sentence segment at the specified index as a string value. If p.Err() is not
// nil, this function returns "" and leaves the error unchanged.
func (p *SegmentParser) AsString(i int8) string {
	if p.err != nil {
		return "" // There's already an error; exit early.
	}

	if p.checkInRange(i); p.err != nil {
		return ""
	}

	if p.segments[i] == "" {
		return ""
	}

	return p.segments[i]
}

// RequireString parses the sentence segment at the specified index as a string value and ensures
// that it matches the required value s (case-insensitive). If p.Err() is not nil, this function
// returns an empty string and leaves the error unchanged.
func (p *SegmentParser) RequireString(i int8, s string) string {
	if p.checkInRange(i); p.err != nil {
		return ""
	}

	if !strings.EqualFold(s, p.segments[i]) {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be \"%s\" (case insensitive) but was \"%s\"", s, p.segments[i]),
		}

		return ""
	}

	return p.segments[i]
}

// RequireStrings parses the sentence segment at the specified index as a string value and ensures
// that it matches one of the required values in s (case-insensitive). If p.Err() is not nil, this
// function returns an empty string and leaves the error unchanged.
func (p *SegmentParser) RequireStrings(i int8, s []string) string {
	if p.checkInRange(i); p.err != nil {
		return ""
	}

	for _, st := range s {
		if strings.EqualFold(st, p.segments[i]) {
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

// SegmentCount returns the number of segments in the parsed sentence.
func (p *SegmentParser) SegmentCount() int {
	return len(p.segments)
}

// --- Private -----------------------------------------------------------------

func (p *SegmentParser) asInt(i int8, bitSize int) interface{} {
	if p.err != nil {
		return 0 // There's already an error; exit early.
	}

	if p.checkInRange(i); p.err != nil {
		return 0
	}

	if p.segments[i] == "" {
		return 0
	}

	val, err := strconv.ParseInt(p.segments[i], 10, bitSize)
	if err != nil {
		p.err = &ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an int%d but was \"%s\"", bitSize, p.segments[i]),
		}

		return 0
	}

	switch bitSize {
	case 8:
		return int8(val)
	case 16:
		return int16(val)
	case 32:
		return int32(val)
	case 64:
		return val // Type is already int64
	}

	panic("should have returned (but did not) because bitSize was " + fmt.Sprintf("%d", bitSize))
}

func (p *SegmentParser) checkInRange(i int8) {
	if p.err != nil {
		return // There's already an error; exit early.
	}

	if int8(len(p.segments)-1) < i {
		p.err = &ParsingError{Segment: i, Message: "is out of range"}
	}
}
