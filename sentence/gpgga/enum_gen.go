// Code generated by "enumer -type=NorthSouth,EastWest,FixQuality -text -linecomment -transform=first-upper -output=enum_gen.go"; DO NOT EDIT.

package gpgga

import (
	"fmt"
	"strings"
)

const _NorthSouthName = "NS"

var _NorthSouthIndex = [...]uint8{0, 1, 2}

const _NorthSouthLowerName = "ns"

func (i NorthSouth) String() string {
	i -= 1
	if i < 0 || i >= NorthSouth(len(_NorthSouthIndex)-1) {
		return fmt.Sprintf("NorthSouth(%d)", i+1)
	}
	return _NorthSouthName[_NorthSouthIndex[i]:_NorthSouthIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NorthSouthNoOp() {
	var x [1]struct{}
	_ = x[North-(1)]
	_ = x[South-(2)]
}

var _NorthSouthValues = []NorthSouth{North, South}

var _NorthSouthNameToValueMap = map[string]NorthSouth{
	_NorthSouthName[0:1]:      North,
	_NorthSouthLowerName[0:1]: North,
	_NorthSouthName[1:2]:      South,
	_NorthSouthLowerName[1:2]: South,
}

var _NorthSouthNames = []string{
	_NorthSouthName[0:1],
	_NorthSouthName[1:2],
}

// NorthSouthString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NorthSouthString(s string) (NorthSouth, error) {
	if val, ok := _NorthSouthNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NorthSouthNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NorthSouth values", s)
}

// NorthSouthValuDataStatus,Modees returns all values of the enum
func NorthSouthValues() []NorthSouth {
	return _NorthSouthValues
}

// NorthSouthStrings returns a slice of all String values of the enum
func NorthSouthStrings() []string {
	strs := make([]string, len(_NorthSouthNames))
	copy(strs, _NorthSouthNames)
	return strs
}

// IsANorthSouth returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NorthSouth) IsANorthSouth() bool {
	for _, v := range _NorthSouthValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for NorthSouth
func (i NorthSouth) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for NorthSouth
func (i *NorthSouth) UnmarshalText(text []byte) error {
	var err error
	*i, err = NorthSouthString(string(text))
	return err
}

const _EastWestName = "EW"

var _EastWestIndex = [...]uint8{0, 1, 2}

const _EastWestLowerName = "ew"

func (i EastWest) String() string {
	i -= 1
	if i < 0 || i >= EastWest(len(_EastWestIndex)-1) {
		return fmt.Sprintf("EastWest(%d)", i+1)
	}
	return _EastWestName[_EastWestIndex[i]:_EastWestIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EastWestNoOp() {
	var x [1]struct{}
	_ = x[East-(1)]
	_ = x[West-(2)]
}

var _EastWestValues = []EastWest{East, West}

var _EastWestNameToValueMap = map[string]EastWest{
	_EastWestName[0:1]:      East,
	_EastWestLowerName[0:1]: East,
	_EastWestName[1:2]:      West,
	_EastWestLowerName[1:2]: West,
}

var _EastWestNames = []string{
	_EastWestName[0:1],
	_EastWestName[1:2],
}

// EastWestString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EastWestString(s string) (EastWest, error) {
	if val, ok := _EastWestNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EastWestNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EastWest values", s)
}

// EastWestValues returns all values of the enum
func EastWestValues() []EastWest {
	return _EastWestValues
}

// EastWestStrings returns a slice of all String values of the enum
func EastWestStrings() []string {
	strs := make([]string, len(_EastWestNames))
	copy(strs, _EastWestNames)
	return strs
}

// IsAEastWest returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EastWest) IsAEastWest() bool {
	for _, v := range _EastWestValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for EastWest
func (i EastWest) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for EastWest
func (i *EastWest) UnmarshalText(text []byte) error {
	var err error
	*i, err = EastWestString(string(text))
	return err
}

const _FixQualityName = "012345678"

var _FixQualityIndex = [...]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

const _FixQualityLowerName = "012345678"

func (i FixQuality) String() string {
	i -= 1
	if i < 0 || i >= FixQuality(len(_FixQualityIndex)-1) {
		return fmt.Sprintf("FixQuality(%d)", i+1)
	}
	return _FixQualityName[_FixQualityIndex[i]:_FixQualityIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FixQualityNoOp() {
	var x [1]struct{}
	_ = x[InvalidFixQuality-(1)]
	_ = x[GPSFixQuality-(2)]
	_ = x[DGPSFixQuality-(3)]
	_ = x[PPSFixQuality-(4)]
	_ = x[RTKFixQuality-(5)]
	_ = x[FloatRTKFixQuality-(6)]
	_ = x[EstimatedFixQuality-(7)]
	_ = x[ManualInputFixQuality-(8)]
	_ = x[SimulationFixQuality-(9)]
}

var _FixQualityValues = []FixQuality{InvalidFixQuality, GPSFixQuality, DGPSFixQuality, PPSFixQuality, RTKFixQuality, FloatRTKFixQuality, EstimatedFixQuality, ManualInputFixQuality, SimulationFixQuality}

var _FixQualityNameToValueMap = map[string]FixQuality{
	_FixQualityName[0:1]:      InvalidFixQuality,
	_FixQualityLowerName[0:1]: InvalidFixQuality,
	_FixQualityName[1:2]:      GPSFixQuality,
	_FixQualityLowerName[1:2]: GPSFixQuality,
	_FixQualityName[2:3]:      DGPSFixQuality,
	_FixQualityLowerName[2:3]: DGPSFixQuality,
	_FixQualityName[3:4]:      PPSFixQuality,
	_FixQualityLowerName[3:4]: PPSFixQuality,
	_FixQualityName[4:5]:      RTKFixQuality,
	_FixQualityLowerName[4:5]: RTKFixQuality,
	_FixQualityName[5:6]:      FloatRTKFixQuality,
	_FixQualityLowerName[5:6]: FloatRTKFixQuality,
	_FixQualityName[6:7]:      EstimatedFixQuality,
	_FixQualityLowerName[6:7]: EstimatedFixQuality,
	_FixQualityName[7:8]:      ManualInputFixQuality,
	_FixQualityLowerName[7:8]: ManualInputFixQuality,
	_FixQualityName[8:9]:      SimulationFixQuality,
	_FixQualityLowerName[8:9]: SimulationFixQuality,
}

var _FixQualityNames = []string{
	_FixQualityName[0:1],
	_FixQualityName[1:2],
	_FixQualityName[2:3],
	_FixQualityName[3:4],
	_FixQualityName[4:5],
	_FixQualityName[5:6],
	_FixQualityName[6:7],
	_FixQualityName[7:8],
	_FixQualityName[8:9],
}

// FixQualityString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FixQualityString(s string) (FixQuality, error) {
	if val, ok := _FixQualityNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FixQualityNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FixQuality values", s)
}

// FixQualityValues returns all values of the enum
func FixQualityValues() []FixQuality {
	return _FixQualityValues
}

// FixQualityStrings returns a slice of all String values of the enum
func FixQualityStrings() []string {
	strs := make([]string, len(_FixQualityNames))
	copy(strs, _FixQualityNames)
	return strs
}

// IsAFixQuality returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FixQuality) IsAFixQuality() bool {
	for _, v := range _FixQualityValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for FixQuality
func (i FixQuality) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for FixQuality
func (i *FixQuality) UnmarshalText(text []byte) error {
	var err error
	*i, err = FixQualityString(string(text))
	return err
}
