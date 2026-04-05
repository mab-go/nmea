package gpgsa

import (
	"fmt"

	"github.com/mab-go/nmea/sentence"
)

// SegmentParser extends sentence.SegmentParser to provide GPGSA-specific segment parsing methods.
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

// AsSelectionMode parses the input segment at the specified index as a SelectionMode value. If
// p.Err() is not nil, this function returns SelectionMode(0) and leaves the error unchanged.
func (p *SegmentParser) AsSelectionMode(i int8) SelectionMode {
	if p.Err() != nil {
		return SelectionMode(0)
	}

	sm, err := SelectionModeString(p.AsString(i))
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a SelectionMode but was \"%s\"", p.AsString(i)),
		}

		return SelectionMode(0)
	}

	return sm
}

// AsFixMode parses the input segment at the specified index as a FixMode value. If p.Err() is not
// nil, this function returns FixMode(0) and leaves the error unchanged.
func (p *SegmentParser) AsFixMode(i int8) FixMode {
	if p.Err() != nil {
		return FixMode(0)
	}

	fm, err := FixModeString(p.AsString(i))
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a FixMode but was \"%s\"", p.AsString(i)),
		}

		return FixMode(0)
	}

	return fm
}
