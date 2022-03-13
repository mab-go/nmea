// Package gpgll contains data structures and functions related to NMEA sentences
// of type "GPGLL".
package gpgll // import "gopkg.in/mab-go/nmea.v0/sentence/gpgll"

// @formatter:off

// GPGLL represents an NMEA sentence of type "GPGLL".
type GPGLL struct { // nolint: maligned
	// @formatter:on

	// Latitude is the "latitude" component of a GPS fix. The format is (d)ddmm.mmmmmmm. For
	// example, the value 5106.7198674 represents a latitude value of 51° 6.7198674'. It is element
	// [1] of a GPGLL sentence.
	Latitude float64

	// NorthSouth indicates the hemisphere in which the latitude value resides. It can be either
	// "N" or "S". It is element [2] of a GPGLL sentence.
	NorthSouth string

	// Longitude is the "longitude" component of a GPS fix. The format is (d)ddmm.mmmm. For example,
	// the value 11402.3587526 represents a longitude value of 114° 2.3587526'. It is element [3] of
	// a GPGLL sentence.
	Longitude float64

	// EastWest indicates the hemisphere in which the longitude value resides. It can be either "E"
	// or "W". It is element [4] of a GPGLL sentence.
	EastWest string

	// FixTime is the time at which the GPS fix was acquired. The format is (h)hmmss.sss. For
	// example, the value 174831.864 represents the time 17:48:31.864. It is element [5] of a
	// GPGLL sentence.
	FixTime float32
}