package sentence

import (
	"fmt"
	"testing"

	"gopkg.in/mab-go/nmea.v0/sentence/testhelp"
)

type checksumTestData struct {
	Title, Sentence, ActualChecksum, AdvertisedChecksum, ErrMsg string
}

func mapChecksumTestData(title string, input map[string]interface{}) interface{} {
	return checksumTestData{
		Title:              title,
		Sentence:           testhelp.EnsureString(input["Sentence"]),
		ActualChecksum:     testhelp.OptString(input["ActualChecksum"]),
		AdvertisedChecksum: testhelp.OptString(input["AdvertisedChecksum"]),
		ErrMsg:             testhelp.OptString(input["ErrMsg"]),
	}
}

func sortChecksumTestData(result []interface{}, i, j int) bool {
	return result[i].(checksumTestData).Title < result[j].(checksumTestData).Title
}

func TestVerifyChecksum_goodData(t *testing.T) {
	for _, data := range testhelp.ReadTestData("good-data", mapChecksumTestData, sortChecksumTestData) {
		d := data.(checksumTestData)

		t.Run(d.Title, func(t *testing.T) {
			err := VerifyChecksum(d.Sentence)
			if err != nil {
				t.Errorf("checksum verification failed: %v", err)
			}
		})
	}
}

func TestVerifyChecksum_invalidChecksums(t *testing.T) {
	for _, data := range testhelp.ReadTestData("bad-invalid-checksums", mapChecksumTestData, sortChecksumTestData) {
		d := data.(checksumTestData)

		t.Run(d.Title, func(t *testing.T) {
			err := VerifyChecksum(d.Sentence)
			if err == nil {
				t.Error("checksum verification succeeded (but should not have)")
			}

			expectedMsg := fmt.Sprintf(
				"calculated checksum value \"%s\" does not match sentence-specified value of \"%s\"",
				d.ActualChecksum,
				d.AdvertisedChecksum)
			if err.Error() != expectedMsg {
				t.Errorf("error message should have been '%v' but was '%v'", expectedMsg, err.Error())
			}
		})
	}
}

func TestVerifyChecksum_malformedData(t *testing.T) {
	for _, data := range testhelp.ReadTestData("bad-malformed", mapChecksumTestData, sortChecksumTestData) {
		d := data.(checksumTestData)

		t.Run(d.Title, func(t *testing.T) {
			err := VerifyChecksum(d.Sentence)
			if err == nil {
				t.Error("checksum verification succeeded (but should not have)")
			}

			if err.Error() != d.ErrMsg {
				t.Errorf("error message should have been '%v' but was '%v'", d.ErrMsg, err.Error())
			}
		})
	}
}

func ExampleVerifyChecksum_validSentence() {
	err := VerifyChecksum("$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*70")
	fmt.Printf("err == %v", err)
	// Output:
	// err == <nil>
}

func ExampleVerifyChecksum_invalidSentence() {
	err := VerifyChecksum("$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*24")
	fmt.Printf("err == %v", err)
	// Output:
	// err == calculated checksum value "70" does not match sentence-specified value of "24"
}
