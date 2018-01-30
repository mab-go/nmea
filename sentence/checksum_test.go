package sentence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

// --- Test Helpers ------------------------------------------------------------

type testSentence struct {
	Sentence, Checksum, ErrMsg string
}

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

func TestVerifyChecksum_goodData(t *testing.T) {
	var testData []testSentence
	readTestData("good-with-checksums.json", &testData)

	for i, s := range testData {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			err := VerifyChecksum(s.Sentence)
			if err != nil {
				t.Errorf("checksum verification failed for test sentence [%v]: %v", i, err)
			}
		})
	}
}

func TestVerifyChecksum_badData(t *testing.T) {
	var testData []testSentence
	readTestData("bad-invalid-checksums.json", &testData)
	for i, s := range testData {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			err := VerifyChecksum(s.Sentence)
			if err == nil {
				t.Errorf("checksum verification passed (but should not have) for test sentence [%v]", i)
			}

			if err.Error() != s.ErrMsg {
				t.Errorf("error message should have been '%v' but was '%v' for test sentence [%v]", s.ErrMsg, err.Error(), i)
			}
		})
	}
}

func TestVerifyChecksum_doesNotStartWithDollar(t *testing.T) {
	err := VerifyChecksum("GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*7F")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	expected := "character [0] must be \"$\" but was \"G\""
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestVerifyChecksum_oneDigitChecksum(t *testing.T) {
	err := VerifyChecksum("$GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*7")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	expected := "there must be exactly 2 characters remaining after \"*\" but there was/were 1"
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestVerifyChecksum_zeroDigitChecksum(t *testing.T) {
	err := VerifyChecksum("$GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,*")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	expected := "there must be exactly 2 characters remaining after \"*\" but there was/were 0"
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
	}
}

func TestVerifyChecksum_noChecksumIndicator(t *testing.T) {
	err := VerifyChecksum("$GPGGA,002454,3553.5295,N,13938.6570,E,1,05,2.2,18.3,M,39.0,M,,")
	if err == nil {
		t.Error("checksum verification passed (but should not have)")
	}

	expected := "sentence does not contain a checksum"
	if err.Error() != expected {
		t.Errorf("error message should have been '%v' but was '%v'", expected, err.Error())
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
	// err == calculated checksum value "70" does not match expected value of "24"
}
