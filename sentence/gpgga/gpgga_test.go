package gpgga

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// --- Test Helpers ------------------------------------------------------------

func readTestData(path string, data interface{}) {
	contents, err := ioutil.ReadFile("_testdata/" + path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(contents, &data)
	if err != nil {
		panic(err)
	}
}

// --- Test Functions ----------------------------------------------------------

// nolint: gocyclo
func TestParseGPGGA_goodData(t *testing.T) {
	type expectedGpgga struct {
		GPGGA
		Title string `json:"_Title"`
	}
	testData := make(map[string]expectedGpgga)
	readTestData("good-with-expected-vals.json", &testData)

	for sentence, expected := range testData {
		t.Run(expected.Title, func(t *testing.T) {
			actual, err := ParseGPGGA(sentence)

			if err != nil {
				t.Errorf("error creating GPGGA from NMEA sentence \"%v\": %v", expected.Title, err)
				return
			}

			if actual.FixTime != expected.FixTime {
				t.Errorf("FixTime should have been %v but was %v for NMEA sentence \"%v\"", expected.FixTime, actual.FixTime, expected.Title)
			}

			if actual.Latitude != expected.Latitude {
				t.Errorf("Latitude should have been %v but was %v for NMEA sentence \"%v\"", expected.Latitude, actual.Latitude, expected.Title)
			}

			if actual.NorthSouth != expected.NorthSouth {
				t.Errorf("NorthSouth should have been %v but was %v for NMEA sentence \"%v\"", expected.NorthSouth, actual.NorthSouth, expected.Title)
			}

			if actual.Longitude != expected.Longitude {
				t.Errorf("Longitude should have been %v but was %v for NMEA sentence \"%v\"", expected.Longitude, actual.Longitude, expected.Title)
			}

			if actual.EastWest != expected.EastWest {
				t.Errorf("EastWest should have been %v but was %v for NMEA sentence \"%v\"", expected.EastWest, actual.EastWest, expected.Title)
			}

			if actual.FixQuality != expected.FixQuality {
				t.Errorf("FixQuality should have been %v but was %v for NMEA sentence \"%v\"", expected.FixQuality, actual.FixQuality, expected.Title)
			}

			if actual.SatCount != expected.SatCount {
				t.Errorf("SatCount should have been %v but was %v for NMEA sentence \"%v\"", expected.SatCount, actual.SatCount, expected.Title)
			}

			if actual.HDOP != expected.HDOP {
				t.Errorf("HDOP should have been %v but was %v for NMEA sentence \"%v\"", expected.HDOP, actual.HDOP, expected.Title)
			}

			if actual.Altitude != expected.Altitude {
				t.Errorf("Altitude should have been %v but was %v for NMEA sentence \"%v\"", expected.Altitude, actual.Altitude, expected.Title)
			}

			if actual.AltitudeUOM != expected.AltitudeUOM {
				t.Errorf("AltitudeUOM should have been %v but was %v for NMEA sentence \"%v\"", expected.AltitudeUOM, actual.AltitudeUOM, expected.Title)
			}

			if actual.GeoidHeight != expected.GeoidHeight {
				t.Errorf("GeoidHeight should have been %v but was %v for NMEA sentence \"%v\"", expected.GeoidHeight, actual.GeoidHeight, expected.Title)
			}

			if actual.GeoidHeightUOM != expected.GeoidHeightUOM {
				t.Errorf("GeoidHeightUOM should have been %v but was %v for NMEA sentence \"%v\"", expected.GeoidHeightUOM, actual.GeoidHeightUOM, expected.Title)
			}

			if actual.DGPSUpdateAge != expected.DGPSUpdateAge {
				t.Errorf("DGPSUpdateAge should have been %v but was %v for NMEA sentence \"%v\"", expected.DGPSUpdateAge, actual.DGPSUpdateAge, expected.Title)
			}

			if actual.DGPSStationID != expected.DGPSStationID {
				t.Errorf("DGPSStationID should have been %v but was %v for NMEA sentence \"%v\"", expected.DGPSStationID, actual.DGPSStationID, expected.Title)
			}
		})
	}
}

func TestParseGPGGA_invalidChecksum(t *testing.T) {
	gpgga, err := ParseGPGGA("$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*42")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	if gpgga != nil {
		t.Errorf("result should have been <nil> but was %v", gpgga)
	}

	expected := "calculated checksum value \"70\" does not match expected value of \"42\""
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestParseGPGGA_badSegments(t *testing.T) {
	type badSentence struct {
		Sentence string
		ErrMsg   string
		Title    string `json:"_Title"`
	}
	var testData []badSentence
	readTestData("bad-with-err-msgs.json", &testData)

	for i, data := range testData {
		t.Run(data.Title, func(t *testing.T) {
			gpgga, err := ParseGPGGA(data.Sentence)

			if err == nil {
				t.Errorf("parsing succeeded (but should not have) for test sentence [%v]", i)
				return
			}

			if gpgga != nil {
				t.Errorf("result should have been <nil> but was %v for test sentence [%v]", gpgga, i)
				return
			}

			if err.Error() != data.ErrMsg {
				t.Errorf("error message should have been '%v' but was '%v' for test sentence [%v]", data.ErrMsg, err.Error(), i)
				return
			}
		})
	}
}

func TestGPGGA_GetSentenceType(t *testing.T) {
	gpgga := &GPGGA{}
	if st := gpgga.GetSentenceType(); st != "GPGGA" {
		t.Errorf("GetSentenceType() should have returned \"GPGGA\" but returned \"%v\"", st)
	}
}

// --- Example Functions -------------------------------------------------------

func ExampleParseGPGGA() {
	sentence := "$GPGGA,023042,3907.3837,N,12102.4684,W,1,04,2.3,507.3,M,-24.1,M,,*75"
	gpgga, err := ParseGPGGA(sentence)
	if err != nil {
		// Handle error
	}

	fmt.Printf("%+v", gpgga)
	// Output:
	// &{FixTime:23042 Latitude:3907.3837 NorthSouth:N Longitude:12102.4684 EastWest:W FixQuality:1 SatCount:4 HDOP:2.3 Altitude:507.3 AltitudeUOM:M GeoidHeight:-24.1 GeoidHeightUOM:M DGPSUpdateAge:0 DGPSStationID:0}
}
