package gpvtg

// Mode indicates the FAA mode indicator for a GPS fix (NMEA 2.3+).
type Mode int

const (
	// Autonomous represents an autonomous GPS fix mode.
	Autonomous Mode = iota + 1 // A

	// Differential represents a differential GPS fix mode.
	Differential // D

	// Estimated represents an estimated (dead reckoning) fix mode.
	Estimated // E

	// NotValid represents a GPS fix that is not valid.
	NotValid // N

	// Simulator represents a simulated GPS fix mode.
	Simulator // S
)

//go:generate go run github.com/dmarkham/enumer@v1.6.3 -type=Mode -text -linecomment -transform=first-upper -output=enum_gen.go
