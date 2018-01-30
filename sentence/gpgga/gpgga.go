// Package gpgga contains data structures and functions related to NMEA sentences
// of type "GPGGA".
package gpgga // import "gopkg.in/mab-go/nmea.v0/sentence/gpgga"

import (
	"fmt"
	"strconv"
	"strings"

	"bitbucket.org/mabgo/nmea/sentence"
)

// --- Public ------------------------------------------------------------------

// GPGGA represents an NMEA sentence of type "GPGGA".
type GPGGA struct {
	// FixTime is the time at which the GPS fix was acquired. The format is
	// (h)hmmss.sss. For example, the value 174831.864 represents the time
	// 17:48:31.864. It is element [1] of a GPGGA sentence.
	FixTime float32

	// Latitude is the "latitude" component of a GPS fix. The format is
	// (d)ddmm.mmmm. For example, the value 4807.038 represents a latitude
	// value of 48° 7.038'. It is element [2] of a GPGGA sentence.
	Latitude float64

	// NorthSouth indicates the hemisphere in which the latitude value resides.
	// It can be either "N" or "S". It is element [3] of a GPGGA sentence.
	NorthSouth string

	// Longitude is the "longitude" component of a GPS fix. The format is
	// (d)ddmm.mmmm. For example, the value 01131.215 represents a longitude
	// value of 11° 31.215'. It is element [4] of a GPGGA sentence.
	Longitude float64

	// EastWest indicates the hemisphere in which the longitude value resides.
	// It can be either "E" or "W". It is element [5] of a GPGGA sentence.
	EastWest string

	// FixQuality indicates the type/quality of the GPS fix. It is element [6]
	// of a GPGGA sentence.
	//
	// The following values are valid:
	//     0 = Invalid
	//     1 = GPS fix (SPS)
	//     2 = DGPS fix
	//     3 = PPS fix
	//     4 = Real Time Kinematic
	//     5 = Float RTK
	//     6 = Estimated (dead reckoning) (NMEA 2.3 feature)
	//     7 = Manual input mode
	//     8 = Simulation mode
	FixQuality int8

	// SatCount is the number of satellites used to obtain the GPS fix. It is
	// element [7] of a GPGGA sentence.
	SatCount int8

	// HDOP is the horizontal dilution of precision (HDOP) of the GPS fix. It
	// indicates a relative "confidence" level in the precision reported. It is
	// element [8] of a GPGGA sentence.
	//
	// Refer to https://en.wikipedia.org/wiki/Dilution_of_precision_(navigation)#Meaning_of_DOP_Values
	// for a better understanding of the meaning of HDOP values.
	HDOP float32

	// Altitude is the above or below mean sea level for the GPS fix. Its unit of
	// measure is specified by the AltitudeUOM field. It is element [9] of a GPGGA
	// sentence.
	Altitude float32

	// AltitudeUOM is the unit of measure in which Altitude is expressed. It should
	// always be "M" (meters). It is element [10] of a GPGGA sentence.
	AltitudeUOM string

	// GeoidHeight is the height of the geoid above or below the WGS84 ellipsoid.
	// Its unit of measure is specified by the GeoidHeightUOM field. It is element
	// [11] of a GPGGA sentence.
	GeoidHeight float32

	// GeoidHeightUOM is the unit of measure in which GeoidHeight is expressed. It
	// should always be "M" (meters). It is element [12] of a GPGGA sentence.
	GeoidHeightUOM string

	// DGPSUpdateAge is the age (in seconds) since the last update from a differential
	// GPS reference station. It is element [13] of a GPGGA sentence. If differential
	// GPS was not used to obtain the fix, (i.e., if FixQuality is not 2), then both
	// DGPSUpdateAge and DGPSStationID should be 0.
	DGPSUpdateAge int32

	// DGPSStationID is the unique identifier for the differential GPS reference station
	// that was used to obtain the GPS fix (if DGPS was used). It is element [14] of a
	// GPGGA sentence. If differential GPS was not used to obtain the fix, (i.e., if
	// FixQuality is not 2), then both DGPSUpdateAge and DGPSStationID should be 0.
	DGPSStationID int16
}

// GetSentenceType returns the type of NMEA sentence represented by the struct
// GPGGA. It always returns "GPGGA". It represents element [0] of a GPGGA sentence.
func (g GPGGA) GetSentenceType() string {
	return "GPGGA"
}

// Ensure that GPGGA properly implements the NMEASentence interface
var _ sentence.NMEASentence = GPGGA{}

// ParseGPGGA parses a GPGGA sentence string and returns a pointer to a GPGGA struct
// (or an error if the sentence is invalid).
func ParseGPGGA(s string) (*GPGGA, error) {
	gpgga := &GPGGA{}

	if err := sentence.VerifyChecksum(s); err != nil {
		return nil, err
	}

	// Strip the first character ("$") and the last three characters (the checksum),
	// and then split the remaining string on a comma (",").
	parts := strings.Split(strings.ToUpper(s[1:len(s)-3]), ",")

	// Part [0]: SentenceType
	if parts[0] != "GPGGA" {
		return nil, fmt.Errorf("sentence segment [0] must be \"GPGGA\" but was \"%v\"", parts[0])
	}

	// Part [1]: FixTime
	fixTime, err := parseFixTime(parts[1])
	if err != nil {
		return nil, err
	}
	gpgga.FixTime = fixTime

	// Part [2]: Latitude
	latitude, err := parseLatitude(parts[2])
	if err != nil {
		return nil, err
	}
	gpgga.Latitude = latitude

	// Part [3]: NorthSouth
	northSouth, err := parseNorthSouth(parts[3])
	if err != nil {
		return nil, err
	}
	gpgga.NorthSouth = northSouth

	// Part [4]: Longitude
	longitude, err := parseLongitude(parts[4])
	if err != nil {
		return nil, err
	}
	gpgga.Longitude = longitude

	// Part [5]: EastWest
	eastWest, err := parseEastWest(parts[5])
	if err != nil {
		return nil, err
	}
	gpgga.EastWest = eastWest

	// Part [6]: FixQuality
	fixQuality, err := parseFixQuality(parts[6])
	if err != nil {
		return nil, err
	}
	gpgga.FixQuality = fixQuality

	// Part [7]: SatCount
	satCount, err := parseSatCount(parts[7])
	if err != nil {
		return nil, err
	}
	gpgga.SatCount = satCount

	// Part [8]: HDOP
	hdop, err := parseHdop(parts[8])
	if err != nil {
		return nil, err
	}
	gpgga.HDOP = hdop

	// Part [9]: Altitude
	altitude, err := parseAltitude(parts[9])
	if err != nil {
		return nil, err
	}
	gpgga.Altitude = altitude

	// Part [10]: AltitudeUOM
	altitudeUom, err := parseAltitudeUom(parts[10])
	if err != nil {
		return nil, err
	}
	gpgga.AltitudeUOM = altitudeUom

	// Part [11]: GeoidHeight
	geoidHeight, err := parseGeoidHeight(parts[11])
	if err != nil {
		return nil, err
	}
	gpgga.GeoidHeight = geoidHeight

	// Part [12]: GeoidHeightUOM
	geoidHeightUom, err := parseGeoidHeightUom(parts[12])
	if err != nil {
		return nil, err
	}
	gpgga.GeoidHeightUOM = geoidHeightUom

	// Part [13]: DGPSUpdateAge
	dgpsUpdateAge, err := parseDgpsUpdateAge(parts[13])
	if err != nil {
		return nil, err
	}
	gpgga.DGPSUpdateAge = dgpsUpdateAge

	// Part [14]: DGPS station ID
	dgpsStationID, err := parseDgpsStationID(parts[14])
	if err != nil {
		return nil, err
	}
	gpgga.DGPSStationID = dgpsStationID

	return gpgga, nil
}

// --- Private -----------------------------------------------------------------

func parseFixTime(s string) (float32, error) {
	//var fixTime float32
	fixTime, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("sentence segment [1] must be parsable as a float32 but was \"%v\"", s)
	}

	return float32(fixTime), nil
}

func parseLatitude(s string) (float64, error) {
	latitude, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("sentence segment [2] must be parsable as a float64 but was \"%v\"", s)
	}

	return latitude, nil
}

func parseNorthSouth(s string) (string, error) {
	var northSouth string
	if s == "N" || s == "S" {
		northSouth = s
	} else {
		return "", fmt.Errorf("sentence segment [3] must be \"N\" or \"S\" (case insensitive) but was \"%v\"", s)
	}

	return northSouth, nil
}

func parseLongitude(s string) (float64, error) {
	longitude, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("sentence segment [4] must be parsable as a float64 but was \"%v\"", s)
	}

	return longitude, nil
}

func parseEastWest(s string) (string, error) {
	var eastWest string
	if s == "E" || s == "W" {
		eastWest = s
	} else {
		return "", fmt.Errorf("sentence segment [5] must be \"E\" or \"W\" (case insensitive) but was \"%v\"", s)
	}

	return eastWest, nil
}

func parseFixQuality(s string) (int8, error) {
	var fixQuality int8
	if v, err := strconv.ParseInt(s, 10, 8); err != nil {
		return 0, fmt.Errorf("sentence segment [6] must be parsable as an int8 but was \"%v\"", s)
	} else if v < 0 || v > 8 {
		return 0, fmt.Errorf("sentence segment [6] must be within 0..8 but was %v", s)
	} else {
		fixQuality = int8(v)
	}

	return fixQuality, nil
}

func parseSatCount(s string) (int8, error) {
	var satCount int8
	if v, err := strconv.ParseInt(s, 10, 8); err != nil {
		return 0, fmt.Errorf("sentence segment [7] must be parsable as an int8 but was \"%v\"", s)
	} else if v < 0 {
		return 0, fmt.Errorf("sentence segment [7] must not be negative but was %v", s)
	} else {
		satCount = int8(v)
	}

	return satCount, nil
}

func parseHdop(s string) (float32, error) {
	hdop, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("sentence segment [8] must be parsable as a float32 but was \"%v\"", s)
	}

	return float32(hdop), nil
}

func parseAltitude(s string) (float32, error) {
	altitude, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("sentence segment [9] must be parsable as a float32 but was \"%v\"", s)
	}

	return float32(altitude), nil
}

func parseAltitudeUom(s string) (string, error) {
	var altitudeUom string
	if s == "M" {
		altitudeUom = s
	} else {
		return "", fmt.Errorf("sentence segment [10] must be \"M\" (case insensitive) but was \"%v\"", s)
	}

	return altitudeUom, nil
}

func parseGeoidHeight(s string) (float32, error) {
	geoidHeight, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("sentence segment [11] must be parsable as a float32 but was \"%v\"", s)
	}

	return float32(geoidHeight), nil
}

func parseGeoidHeightUom(s string) (string, error) {
	var geoidHeightUom string
	if s == "M" {
		geoidHeightUom = s
	} else {
		return "", fmt.Errorf("sentence segment [12] must be \"M\" (case insensitive) but was \"%v\"", s)
	}

	return geoidHeightUom, nil
}

func parseDgpsUpdateAge(s string) (int32, error) {
	var dgpsUpdateAge int32
	if s == "" {
		dgpsUpdateAge = 0
	} else if v, err := strconv.ParseInt(s, 10, 32); err != nil {
		return 0, fmt.Errorf("sentence segment [13] must be parsable as an int32 but was \"%v\"", s)
	} else {
		dgpsUpdateAge = int32(v)
	}

	return dgpsUpdateAge, nil
}

func parseDgpsStationID(s string) (int16, error) {
	var dgpsStationID int16
	if s == "" {
		dgpsStationID = 0
	} else if v, err := strconv.ParseInt(s, 10, 32); err != nil {
		return 0, fmt.Errorf("sentence segment [14] must be parsable as an int16 but was \"%v\"", s)
	} else {
		dgpsStationID = int16(v)
	}

	return dgpsStationID, nil
}
