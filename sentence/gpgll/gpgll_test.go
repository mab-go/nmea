package gpgll

import (
	"testing"

	"gopkg.in/mab-go/nmea.v0/sentence/testhelp"
)

type testData struct {
	GPGLL
	Title, Sentence, ErrMsg string
}

func mapGoodTestData(title string, input map[string]interface{}) interface{} {
	northSouth, err := NorthSouthString(testhelp.EnsureString(input["NorthSouth"]))
	if err != nil {
		panic(err)
	}

	// eastWest, err := NorthSouthString(testhelp.EnsureString(input["EastWest"]))
	// if err != nil {
	// 	panic(err)
	// }

	return testData{
		Title:    title,
		Sentence: testhelp.EnsureString(input["Sentence"]),
		GPGLL: GPGLL{
			Latitude:   testhelp.EnsureFloat64(input["Latitude"]),
			NorthSouth: northSouth,
			Longitude:  testhelp.EnsureFloat64(input["Longitude"]),
			EastWest:   "",
			FixTime:    testhelp.EnsureFloat32(input["FixTime"]),
		},
	}
}

func mapBadTestData(title string, input map[string]interface{}) interface{} {
	return testData{
		Title:    title,
		Sentence: testhelp.EnsureString(input["Sentence"]),
		ErrMsg:   testhelp.EnsureString(input["ErrMsg"]),
	}
}

func sortTestData(result []interface{}, i, j int) bool {
	return result[i].(testData).Title < result[j].(testData).Title
}

func TestParseGPGLL_goodData(t *testing.T) {
	for _, d := range testhelp.ReadTestData("good-data", mapGoodTestData, sortTestData) {
		expected := d.(testData)

		t.Run(expected.Title, func(t *testing.T) {
			actual, err := ParseGPGLL(expected.Sentence)

			if err != nil {
				t.Errorf("error creating GPGLL from NMEA sentence \"%v\": %v", expected.Title, err)
				return
			}

			if actual.Latitude != expected.Latitude {
				t.Errorf("Latitude should have been %v but was %v for NMEA sentence \"%v\"", expected.Latitude, actual.Latitude, expected.Title)
			}
		})
	}
}
