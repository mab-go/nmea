package gpvtg

import (
	"fmt"

	"github.com/mab-go/nmea/sentence"
)

// SegmentParser extends sentence.SegmentParser to provide GPVTG-specific segment parsing methods.
type SegmentParser struct {
	sentence.SegmentParser

	err error
}

// Err returns a SegmentParser's error value.
func (p *SegmentParser) Err() error {
	err := p.SegmentParser.Err()
	if err == nil {
		err = p.err
	}

	return err
}

// AsMode parses the input segment at the specified index as a Mode value. If p.Err() is not nil,
// this function returns Mode(0) and leaves the error unchanged. An empty segment returns Mode(0)
// with no error, since the Mode field is optional (NMEA 2.3+).
func (p *SegmentParser) AsMode(i int8) Mode {
	if p.Err() != nil {
		return Mode(0)
	}

	// Mode is optional (NMEA 2.3+); skip if the segment is not present.
	if int(i) >= p.SegmentCount() {
		return Mode(0)
	}

	s := p.AsString(i)
	if p.err != nil {
		return Mode(0)
	}

	if s == "" {
		return Mode(0)
	}

	mode, err := ModeString(s)
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a Mode but was %q", s),
		}

		return Mode(0)
	}

	return mode
}
