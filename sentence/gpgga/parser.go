package gpgga

import (
	"fmt"

	"gopkg.in/mab-go/nmea.v0/sentence"
)

// SegmentParser extends sentence.SegmentParser to provide GPGGA-specific segment parsing methods.
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
	if ns, err := NorthSouthString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a NorthSouth but was \"%s\"", p.AsString(i)),
		}
		return NorthSouth(0)
	} else {
		return ns
	}
}

// AsEastWest parses the input segment at the specified index as an EastWest value. If p.Err() is
// not nil, this function returns EastWest(0) and leaves the error unchanged.
func (p *SegmentParser) AsEastWest(i int8) EastWest {
	if ew, err := EastWestString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as an EastWest but was \"%s\"", p.AsString(i)),
		}
		return EastWest(0)
	} else {
		return ew
	}
}

// AsFixQuality parses the input segment at the specified index as a FixQuality value. If p.Err()
// is not nil, this function returns FixQuality(0) and leaves the error unchanged.
func (p *SegmentParser) AsFixQuality(i int8) FixQuality {
	if ds, err := FixQualityString(p.AsString(i)); err != nil {
		p.err = &sentence.ParsingError{
			Segment: i,
			Message: fmt.Sprintf("must be parsable as a FixQuality but was \"%s\"", p.AsString(i)),
		}
		return FixQuality(0)
	} else {
		return ds
	}
}
