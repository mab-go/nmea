package gpgsa

// SelectionMode indicates whether satellite selection is automatic or manual.
// It can be either "A" or "M".
type SelectionMode int

const (
	// AutomaticSelectionMode represents automatic satellite selection.
	AutomaticSelectionMode SelectionMode = iota + 1 // A

	// ManualSelectionMode represents manual satellite selection.
	ManualSelectionMode // M
)

// FixMode indicates the type of GPS fix: no fix, 2D, or 3D.
type FixMode int

const (
	// NoFix represents no GPS fix. Its NMEA wire value is 1.
	NoFix FixMode = iota + 1 // 1

	// Fix2D represents a 2D GPS fix. Its NMEA wire value is 2.
	Fix2D // 2

	// Fix3D represents a 3D GPS fix. Its NMEA wire value is 3.
	Fix3D // 3
)

//go:generate go run github.com/dmarkham/enumer@v1.6.3 -type=SelectionMode,FixMode -text -linecomment -transform=first-upper -output=enum_gen.go
