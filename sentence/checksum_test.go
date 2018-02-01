package sentence

import (
	"fmt"
	"testing"

	"gopkg.in/mab-go/nmea.v0/sentence/testhelp"
)

// --- Test Functions ----------------------------------------------------------

func TestVerifyChecksum(t *testing.T) {
	// Test with good data
	for _, d := range testhelp.ReadTestData("good/sentences") {
		t.Run(fmt.Sprintf("Good Data/%s", d.Name), func(t *testing.T) {
			err := VerifyChecksum(d.Sentence)
			if err != nil {
				t.Errorf("checksum verification failed: %v", err)
			}
		})
	}

	// Test with invalid checksums
	for _, d := range testhelp.ReadTestData("bad/invalid-checksums") {
		t.Run(fmt.Sprintf("Bad Data/Invalid Checksums/%s", d.Name), func(t *testing.T) {
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

	// Test with malformed sentences
	for _, d := range testhelp.ReadTestData("bad/malformed") {
		t.Run(fmt.Sprintf("Bad Data/Malformed Sentences/%s", d.Name), func(t *testing.T) {
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

// --- Example Functions -------------------------------------------------------

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
