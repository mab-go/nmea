package sentence

import (
	"errors"
	"fmt"
	"strings"
)

// --- Public ------------------------------------------------------------------

// VerifyChecksum verifies the checksum of the given NMEA sentence. It returns
// an error if the sentence's checksum is invalid.
func VerifyChecksum(sentence string) error {
	calculated := 0

	// Iterate over each character in the sentence
	for i, c := range sentence {
		ch := string(c)

		// If this is the first character (char [0])...
		if i == 0 {
			// The first character MUST be "$"
			if ch != "$" {
				return fmt.Errorf("character [0] must be \"$\" but was \"%v\"", ch)
			}

			// '$' is not used as part of the checksum calculation
			continue
		}

		// If this is the beginning of the checksum value ("*")...
		if ch == "*" {
			// There MUST be exactly two characters remaining
			if (i + 2) != (len(sentence) - 1) {
				return fmt.Errorf("there must be exactly 2 characters remaining after \"*\" but there was/were %v", len(sentence)-i-1)
			}

			expectedHex := strings.ToUpper(fmt.Sprintf("%c%c", sentence[i+1], sentence[i+2]))
			calculatedHex := fmt.Sprintf("%02X", calculated)
			if calculatedHex != expectedHex {
				return fmt.Errorf("calculated checksum value \"%v\" does not match expected value of \"%v\"", calculatedHex, expectedHex)
			}

			return nil // No errors
		}

		calculated = calculated ^ int(ch[0])
	}

	return errors.New("sentence does not contain a checksum")
}
