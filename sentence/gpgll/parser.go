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

// AsNorthSouth parses the input segment at the specified index as a NorthSouth value. If p.Err()
// is not nil, this function returns NorthSouth(-1) and leaves the error unchanged.
func (p *SegmentParser) AsNorthSouth(i int8) NorthSouth {
	if ns, err := NorthSouthString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a NorthSouth but was \"%s\"", p.AsString(i)),
		}
		return NorthSouth(-1)
	} else {
		return ns
	}
}

// AsEastWest parses the input segment at the specified index as an EastWest value. If p.Err() is
// not nil, this function returns EastWest(-1) and leaves the error unchanged.
func (p *SegmentParser) AsEastWest(i int8) EastWest {
	if ew, err := EastWestString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an EastWest but was \"%s\"", p.AsString(i)),
		}
		return EastWest(-1)
	} else {
		return ew
	}
}

// AsDataStatus parses the input segment at the specified index as a DataStatus value. If p.Err()
// is not nil, this function returns DataStatus(-1) and leaves the error unchanged.
func (p *SegmentParser) AsDataStatus(i int8) DataStatus {
	if ds, err := DataStatusString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a DataStatus but was \"%s\"", p.AsString(i)),
		}
		return DataStatus(-1)
	} else {
		return ds
	}
}

// AsMode parses the input segment at the specified index as a Mode value. If p.Err() is not
// nil, this function returns Mode(-1) and leaves the error unchanged.
func (p *SegmentParser) AsMode(i int8) Mode {
	if m, err := ModeString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a Mode but was \"%s\"", p.AsString(i)),
		}
		return Mode(-1)
	} else {
		return m
	}
}
