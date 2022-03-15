package gpgll

import (
	"fmt"

	"gopkg.in/mab-go/nmea.v0/sentence"
)

// SegmentParser extends sentence.SegmentParser to provide GPGLL-specific segment parsing methods.
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

// AsNorthSouth parses the input segment at the specified index as a NorthSouth value. If p.Err()
// is not nil, this function returns NorthSouth(0) and leaves the error unchanged.
func (p *SegmentParser) AsNorthSouth(i int8) NorthSouth {
	ns, err := NorthSouthString(p.AsString(i))
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a NorthSouth but was \"%s\"", p.AsString(i)),
		}

		return NorthSouth(0)
	}

	return ns
}

// AsEastWest parses the input segment at the specified index as an EastWest value. If p.Err() is
// not nil, this function returns EastWest(0) and leaves the error unchanged.
func (p *SegmentParser) AsEastWest(i int8) EastWest {
	ew, err := EastWestString(p.AsString(i))
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an EastWest but was \"%s\"", p.AsString(i)),
		}

		return EastWest(0)
	}

	return ew
}

// AsDataStatus parses the input segment at the specified index as a DataStatus value. If p.Err()
// is not nil, this function returns DataStatus(0) and leaves the error unchanged.
func (p *SegmentParser) AsDataStatus(i int8) DataStatus {
	ds, err := DataStatusString(p.AsString(i))
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a DataStatus but was \"%s\"", p.AsString(i)),
		}

		return DataStatus(0)
	}

	return ds
}

// AsMode parses the input segment at the specified index as a Mode value. If p.Err() is not
// nil, this function returns Mode(0) and leaves the error unchanged.
func (p *SegmentParser) AsMode(i int8) Mode {
	m, err := ModeString(p.AsString(i))
	if err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a Mode but was \"%s\"", p.AsString(i)),
		}

		return Mode(0)
	}

	return m
}
