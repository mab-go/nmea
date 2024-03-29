// Package gpgga contains data structures and functions related to NMEA sentences
// of type "GPGGA".
package gpgga // import "gopkg.in/mab-go/nmea.v0/sentence/gpgga"

import (
	"gopkg.in/mab-go/nmea.v0/sentence"
)

// GPGGA represents an NMEA sentence of type "GPGGA".
type GPGGA struct {
	// FixTime is the time at which the GPS fix was acquired. The format is (h)hmmss.sss. For
	// example, the value 174831.864 represents the time 17:48:31.864. It is element [1] of a
	// GPGGA sentence.
	FixTime float32

	// Latitude is the "latitude" component of a GPS fix. The format is (d)ddmm.mmmm. For example,
	// the value 4807.038 represents a latitude value of 48° 7.038'. It is element [2] of a GPGGA
	// sentence.
	Latitude float64

	// NorthSouth indicates the hemisphere in which the latitude value resides. It is element [3] of
	// a GPGGA sentence.
	NorthSouth NorthSouth

	// Longitude is the "longitude" component of a GPS fix. The format is (d)ddmm.mmmm. For example,
	// the value 01131.215 represents a longitude value of 11° 31.215'. It is element [4] of a GPGGA
	// sentence.
	Longitude float64

	// EastWest indicates the hemisphere in which the longitude value resides. It is element [5] of
	// a GPGGA sentence.
	EastWest EastWest

	// FixQuality indicates the type/quality of the GPS fix. It is element [6] of a GPGGA sentence.
	FixQuality FixQuality

	// SatCount is the number of satellites used to obtain the GPS fix. It is element [7] of a GPGGA
	// sentence.
	SatCount int8

	// HDOP is the horizontal dilution of precision (HDOP) of the GPS fix. It indicates a relative
	// "confidence" level in the precision reported. Generally an HDOP of 1.0 is the best possible
	// value. It is element [8] of a GPGGA sentence.
	//
	// Refer to https://en.wikipedia.org/wiki/Dilution_of_precision_(navigation)#Meaning_of_DOP_Values
	// for a better understanding of the meaning of HDOP values.
	HDOP float32

	// Altitude is the above or below mean sea level for the GPS fix. Its unit of measure is
	// specified by the AltitudeUOM field. It is element [9] of a GPGGA sentence.
	Altitude float32

	// AltitudeUOM is the unit of measure in which Altitude is expressed. It should always be "M"
	// (meters). It is element [10] of a GPGGA sentence.
	AltitudeUOM string

	// GeoidHeight is the height of the geoid above or below the WGS84 ellipsoid. Its unit of
	// measure is specified by the GeoidHeightUOM field. It is element [11] of a GPGGA sentence.
	GeoidHeight float32

	// GeoidHeightUOM is the unit of measure in which GeoidHeight is expressed. It should always be
	// "M" (meters). It is element [12] of a GPGGA sentence.
	GeoidHeightUOM string

	// DGPSUpdateAge is the age (in seconds) since the last update from a differential GPS reference
	// station. It is element [13] of a GPGGA sentence. If differential GPS was not used to obtain
	// the fix, (i.e., if FixQuality is not 2), then both DGPSUpdateAge and DGPSStationID should be
	// 0.
	DGPSUpdateAge float32

	// DGPSStationID is the unique identifier for the differential GPS reference station that was
	// used to obtain the GPS fix (if DGPS was used). It is element [14] of a GPGGA sentence. If
	// differential GPS was not used to obtain the fix, (i.e., if FixQuality is not 2), then both
	// DGPSUpdateAge and DGPSStationID should be 0.
	DGPSStationID int16
}

// GetSentenceType returns the type of NMEA sentence represented by the struct GPGGA. It always
// returns "GPGGA". It represents element [0] of a GPGGA sentence.
func (g GPGGA) GetSentenceType() string {
	return "GPGGA"
}

// Ensure that GPGGA properly implements the NMEASentence interface
var _ sentence.NMEASentence = GPGGA{}

// Parse parses a GPGGA sentence string and returns a pointer to a GPGGA struct (or an error if
// the sentence is invalid).
func Parse(s string) (*GPGGA, error) {
	segments := &SegmentParser{}
	if err := segments.Parse(s); err != nil {
		return nil, err
	}

	_ = segments.RequireString(0, "GPGGA") // Verify sentence type
	gpgga := &GPGGA{
		FixTime:        segments.AsFloat32(1),
		Latitude:       segments.AsFloat64(2),
		NorthSouth:     segments.AsNorthSouth(3),
		Longitude:      segments.AsFloat64(4),
		EastWest:       segments.AsEastWest(5),
		FixQuality:     segments.AsFixQuality(6),
		SatCount:       segments.AsInt8(7),
		HDOP:           segments.AsFloat32(8),
		Altitude:       segments.AsFloat32(9),
		AltitudeUOM:    segments.RequireString(10, "M"),
		GeoidHeight:    segments.AsFloat32(11),
		GeoidHeightUOM: segments.RequireString(12, "M"),
		DGPSUpdateAge:  segments.AsFloat32(13),
		DGPSStationID:  segments.AsInt16(14),
	}

	if err := segments.Err(); err != nil {
		return nil, err
	}

	return gpgga, nil
}
