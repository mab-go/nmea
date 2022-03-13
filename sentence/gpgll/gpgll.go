// Package gpgll contains data structures and functions related to NMEA sentences
// of type "GPGLL".
package gpgll // import "gopkg.in/mab-go/nmea.v0/sentence/gpgll"
import (
	"gopkg.in/mab-go/nmea.v0/sentence"
)

// NorthSouth indicates the hemisphere in which a latitude value resides. It can be either
// "N" or "S".
type NorthSouth string

//goland:noinspection GoUnusedConst
const (
	// North represents the northern hemisphere.
	North NorthSouth = "N"

	// South represents the southern hemisphere.
	South NorthSouth = "S"
)

// EastWest indicates the hemisphere in which a longitude value resides. It can be either
// "E" or "W".
type EastWest string

//goland:noinspection GoUnusedConst
const (
	// East represents the northern hemisphere.
	East EastWest = "E"

	// West represents the southern hemisphere.
	West EastWest = "W"
)

// DataStatus represents the status of a GPS fix. It can be either "A" (valid) or
// "V" (invalid).
type DataStatus string

//goland:noinspection GoUnusedConst
const (
	// ValidDataStatus represents a valid GPS fix.
	ValidDataStatus DataStatus = "A"

	// InvalidDataStatus represents an invalid GPS fix.
	InvalidDataStatus DataStatus = "V"
)

// Mode indicates the operating mode of a positioning system. It can be one of "A", "D",
// "E", "M", or "N".
type Mode string

//goland:noinspection GoUnusedConst
const (
	// AutonomousMode represents an autonomous operating mode.
	AutonomousMode Mode = "A"

	// DifferentialMode represents a differential operating mode.
	DifferentialMode Mode = "D"

	// EstimatedMode represents an estimated (dead reckoning) operating mode.
	EstimatedMode Mode = "E"

	// ManualInputMode represents a "manual input" operating mode.
	ManualInputMode Mode = "M"

	// InvalidMode represents an invalid operating mode.
	InvalidMode Mode = "N"
)

// @formatter:off

// GPGLL represents an NMEA sentence of type "GPGLL".
type GPGLL struct { // nolint: maligned
	// @formatter:on

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
