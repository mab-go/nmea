// Package gpvtg contains data structures and functions related to NMEA sentences
// of type "GPVTG".
package gpvtg // import "github.com/mab-go/nmea/sentence/gpvtg"

import (
	"github.com/mab-go/nmea/sentence"
)

// GPVTG represents an NMEA sentence of type "GPVTG". It contains course over ground
// and speed data measured by a GPS receiver.
type GPVTG struct {
	// TrueTrack is the course over ground relative to true north, in degrees. It is
	// element [1] of a GPVTG sentence.
	TrueTrack float64

	// MagTrack is the course over ground relative to magnetic north, in degrees. It is
	// element [3] of a GPVTG sentence.
	MagTrack float64

	// SpeedKnots is the speed over ground in knots. It is element [5] of a GPVTG
	// sentence.
	SpeedKnots float64

	// SpeedKmh is the speed over ground in kilometres per hour. It is element [7] of
	// a GPVTG sentence.
	SpeedKmh float64

	// Mode is the FAA mode indicator (NMEA 2.3+). It is element [9] of a GPVTG
	// sentence. An empty Mode field yields a zero [Mode] without error.
	Mode Mode
}

// GetSentenceType returns the type of NMEA sentence represented by the struct GPVTG. It always
// returns "GPVTG". It represents element [0] of a GPVTG sentence.
func (g GPVTG) GetSentenceType() string {
	return "GPVTG"
}

// Ensure that GPVTG properly implements the NMEASentence interface
var _ sentence.NMEASentence = GPVTG{}

// Parse parses a GPVTG sentence string and returns a pointer to a GPVTG struct (or an error if
// the sentence is invalid).
func Parse(s string) (*GPVTG, error) {
	segments := &SegmentParser{}
	if err := segments.Parse(s); err != nil {
		return nil, err
	}

	_ = segments.RequireString(0, "GPVTG") // Verify sentence type
	_ = segments.RequireString(2, "T") // True track indicator
	_ = segments.RequireString(4, "M") // Magnetic track indicator
	_ = segments.RequireString(6, "N") // Speed in knots indicator
	_ = segments.RequireString(8, "K") // Speed in km/h indicator
	gpvtg := &GPVTG{
		TrueTrack:  segments.AsFloat64(1),
		MagTrack:   segments.AsFloat64(3),
		SpeedKnots: segments.AsFloat64(5),
		SpeedKmh:   segments.AsFloat64(7),
	}

	// Mode is optional (NMEA 2.3+).
	gpvtg.Mode = segments.AsMode(9)

	if err := segments.Err(); err != nil {
		return nil, err
	}

	return gpvtg, nil
}
