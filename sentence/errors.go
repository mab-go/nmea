package sentence

import (
	"fmt"
)

// ParsingError represents an error that occurs when attempting to parse an NMEA
// sentence.
type ParsingError struct {
	Segment int8
	Message string
}

// Error returns the ParsingError's message.
func (e ParsingError) Error() string {
	return fmt.Sprintf("sentence segment [%d] %s", e.Segment, e.Message)
}
