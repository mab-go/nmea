// Package gpgll contains data structures and functions related to NMEA sentences
// of type "GPGLL".
package gpgll // import "gopkg.in/mab-go/nmea.v0/sentence/gpgll"

import (
	"gopkg.in/mab-go/nmea.v0/sentence"
)

// GPGLL represents an NMEA sentence of type "$GPGLL". It contains a position fix location (latitude
// and longitude), the time of the position fix, and the fix status.
type GPGLL struct {
	// Latitude is the "latitude" component of a GPS fix. The format is (d)ddmm.mmmmmmm. For
	// example, the value 5106.7198674 represents a latitude value of 51° 6.7198674'. It is element
	// [1] of a GPGLL sentence.
	Latitude float64

	// NorthSouth indicates the hemisphere in which the latitude value resides. It is element [2] of
	// a GPGLL sentence.
	NorthSouth NorthSouth

	// Longitude is the "longitude" component of a GPS fix. The format is (d)ddmm.mmmm. For example,
	// the value 11402.3587526 represents a longitude value of 114° 2.3587526'. It is element [3] of
	// a GPGLL sentence.
	Longitude float64

	// EastWest indicates the hemisphere in which the longitude value resides. It is element [4] of
	// a GPGLL sentence.
	EastWest EastWest

	// FixTime is the time at which the GPS fix was acquired. The format is (h)hmmss.sss. For
	// example, the value 174831.864 represents the time 17:48:31.864. It is element [5] of a
	// GPGLL sentence.
	FixTime float32

	// DataStatus represents the status of the GPS fix. It can be either "A" (valid) or "V"
	// (invalid). It is element [6] of a GPGLL sentence.
	DataStatus DataStatus

	// Mode indicates the operating mode of a positioning system. It is element [7] of a GPGLL
	// sentence.
	Mode Mode
}

// GetSentenceType returns the type of NMEA sentence represented by the struct GPGLL. It always
// returns "GPGLL". It represents element [0] of a GPGLL sentence.
func (g GPGLL) GetSentenceType() string {
	return "GPGLL"
}

// Ensure that GPGLL properly implements the NMEASentence interface
var _ sentence.NMEASentence = GPGLL{}

// ParseGPGLL parses a GPGLL sentence string and returns a pointer to a GPGLL struct (or an error if
// the sentence is invalid).
func ParseGPGLL(s string) (*GPGLL, error) {
	segments := &SegmentParser{}
	if err := segments.Parse(s); err != nil {
		return nil, err
	}

	_ = segments.RequireString(0, "GPGLL") // Verify sentence type
	gpgll := &GPGLL{
		Latitude:   segments.AsFloat64(1),
		NorthSouth: segments.AsNorthSouth(2),
		Longitude:  segments.AsFloat64(3),
		EastWest:   segments.AsEastWest(4),
		FixTime:    segments.AsFloat32(5),
		DataStatus: segments.AsDataStatus(6),
		Mode:       segments.AsMode(7),
	}

	if err := segments.Err(); err != nil {
		return nil, err
	}

	return gpgll, nil
}
