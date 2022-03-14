package gpgll

// NorthSouth indicates the hemisphere in which a latitude value resides. It can be either
// "N" or "S".
type NorthSouth int

const (
	// North represents the northern hemisphere.
	North NorthSouth = iota

	// South represents the southern hemisphere.
	South
)

// EastWest indicates the hemisphere in which a longitude value resides. It can be either
// "E" or "W".
type EastWest int

const (
	// East represents the eastern hemisphere.
	East EastWest = iota

	// West represents the western hemisphere.
	West
)

// DataStatus represents the status of a GPS fix. It can be either "A" (valid) or
// "V" (invalid).
type DataStatus int

const (
	// ValidDataStatus represents a valid GPS fix.
	ValidDataStatus DataStatus = iota // A

	// InvalidDataStatus represents an invalid GPS fix.
	InvalidDataStatus // V
)

// Mode indicates the operating mode of a positioning system. It can be one of "A", "D",
// "E", "M", or "N".
type Mode int

const (
	// AutonomousMode represents an autonomous operating mode.
	AutonomousMode Mode = iota // A

	// DifferentialMode represents a differential operating mode.
	DifferentialMode // D

	// EstimatedMode represents an estimated (dead reckoning) operating mode.
	EstimatedMode // E

	// ManualInputMode represents a "manual input" operating mode.
	ManualInputMode // M

	// InvalidMode represents an invalid operating mode.
	InvalidMode // N
)

//go:generate enumer -type=NorthSouth,EastWest,DataStatus,Mode -text -linecomment -transform=first-upper -output=enum_gen.go
