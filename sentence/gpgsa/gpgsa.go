// Package gpgsa contains data structures and functions related to NMEA sentences
// of type "GPGSA".
package gpgsa // import "github.com/mab-go/nmea/sentence/gpgsa"

import (
	"github.com/mab-go/nmea/sentence"
)

// GPGSA represents an NMEA sentence of type "GPGSA".
type GPGSA struct {
	// SelectionMode indicates whether satellite selection is automatic or manual. It is element [1]
	// of a GPGSA sentence.
	SelectionMode SelectionMode

	// FixMode indicates whether the GPS fix is no fix, 2D, or 3D. It is element [2] of a GPGSA
	// sentence.
	FixMode FixMode

	// PRNs contains the PRN IDs of the satellites used in the solution. The 12 slots are
	// fixed-position; unused slots are 0. They are elements [3]–[14] of a GPGSA sentence.
	PRNs [12]int8

	// PDOP is the position dilution of precision. It is element [15] of a GPGSA sentence.
	PDOP float32

	// HDOP is the horizontal dilution of precision. It is element [16] of a GPGSA sentence.
	HDOP float32

	// VDOP is the vertical dilution of precision. It is element [17] of a GPGSA sentence.
	VDOP float32
}

// GetSentenceType returns the type of NMEA sentence represented by the struct GPGSA. It always
// returns "GPGSA". It represents element [0] of a GPGSA sentence.
func (g GPGSA) GetSentenceType() string {
	return "GPGSA"
}

// Ensure that GPGSA properly implements the NMEASentence interface
var _ sentence.NMEASentence = GPGSA{}

// Parse parses a GPGSA sentence string and returns a pointer to a GPGSA struct (or an error if
// the sentence is invalid).
func Parse(s string) (*GPGSA, error) {
	segments := &SegmentParser{}
	if err := segments.Parse(s); err != nil {
		return nil, err
	}

	_ = segments.RequireString(0, "GPGSA") // Verify sentence type
	gpgsa := &GPGSA{
		SelectionMode: segments.AsSelectionMode(1),
		FixMode:       segments.AsFixMode(2),
		PRNs: [12]int8{
			segments.AsInt8(3),
			segments.AsInt8(4),
			segments.AsInt8(5),
			segments.AsInt8(6),
			segments.AsInt8(7),
			segments.AsInt8(8),
			segments.AsInt8(9),
			segments.AsInt8(10),
			segments.AsInt8(11),
			segments.AsInt8(12),
			segments.AsInt8(13),
			segments.AsInt8(14),
		},
		PDOP: segments.AsFloat32(15),
		HDOP: segments.AsFloat32(16),
		VDOP: segments.AsFloat32(17),
	}

	if err := segments.Err(); err != nil {
		return nil, err
	}

	return gpgsa, nil
}
