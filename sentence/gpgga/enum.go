package gpgga

// NorthSouth indicates the hemisphere in which a latitude value resides. It can be either
// "N" or "S".
type NorthSouth int

const (
	// North represents the northern hemisphere.
	North NorthSouth = iota + 1 // N

	// South represents the southern hemisphere.
	South // S
)

// EastWest indicates the hemisphere in which a longitude value resides. It can be either
// "E" or "W".
type EastWest int

const (
	// East represents the eastern hemisphere.
	East EastWest = iota + 1 // E

	// West represents the western hemisphere.
	West // W
)

// FixQuality indicates the type/quality of a GPS fix.
type FixQuality int

const (
	// InvalidFixQuality represents an invalid GPS fix quality. Its value is 0.
	InvalidFixQuality FixQuality = iota + 1 // 0

	// GPSFixQuality represents a standard GPS fix quality. Its value is 1.
	GPSFixQuality // 1

	// DGPSFixQuality represents a differential GPS (DGPS) fix quality. Its value is 2.
	DGPSFixQuality // 2

	// PPSFixQuality represents a precise positioning system (PPS) fix quality. Its value is 3.
	PPSFixQuality // 3

	// RTKFixQuality represents a Real Time Kinematic fix quality. Its value is 4.
	RTKFixQuality // 4

	// FloatRTKFixQuality represents a Float Real Time Kinematic fix quality. Its value is 5.
	FloatRTKFixQuality // 5

	// EstimatedFixQuality represents an estimated (dead reckoning) fix quality. Its value is 6.
	EstimatedFixQuality // 6

	// ManualInputFixQuality represents a "manual input mode" fix quality. Its value is 7.
	ManualInputFixQuality // 7

	// SimulationFixQuality represents a "simulation mode" fix quality. Its value is 8.
	SimulationFixQuality // 8
)

//go:generate enumer -type=NorthSouth,EastWest,FixQuality -text -linecomment -transform=first-upper -output=enum_gen.go
