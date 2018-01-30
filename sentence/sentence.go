// Package sentence provides a set of interfaces and functions used to describe
// and work with NMEA sentences.
package sentence // import "gopkg.in/mab-go/nmea.v0/sentence"

// NMEASentence describes the (minimum) functionality of a struct that represents
// an NMEA sentence.
type NMEASentence interface {
	// GetSentenceType returns the type of NMEA sentence represented by a struct
	// that implements this interface. It always represents element [0] of any
	// NMEA sentence.
	GetSentenceType() string
}
