// Code generated by "enumer -type=NorthSouth,EastWest,DataStatus,Mode -text -linecomment -transform=first-upper -output=enum_gen.go"; DO NOT EDIT.

package gpgll

import (
	"fmt"
	"strings"
)

const _NorthSouthName = "NS"

var _NorthSouthIndex = [...]uint8{0, 1, 2}

const _NorthSouthLowerName = "ns"

func (i NorthSouth) String() string {
	if i < 0 || i >= NorthSouth(len(_NorthSouthIndex)-1) {
		return fmt.Sprintf("NorthSouth(%d)", i)
	}
	return _NorthSouthName[_NorthSouthIndex[i]:_NorthSouthIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NorthSouthNoOp() {
	var x [1]struct{}
	_ = x[North-(0)]
	_ = x[South-(1)]
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

// NorthSouthValues returns all values of the enum
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
	if i < 0 || i >= EastWest(len(_EastWestIndex)-1) {
		return fmt.Sprintf("EastWest(%d)", i)
	}
	return _EastWestName[_EastWestIndex[i]:_EastWestIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EastWestNoOp() {
	var x [1]struct{}
	_ = x[East-(0)]
	_ = x[West-(1)]
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

const _DataStatusName = "AV"

var _DataStatusIndex = [...]uint8{0, 1, 2}

const _DataStatusLowerName = "av"

func (i DataStatus) String() string {
	if i < 0 || i >= DataStatus(len(_DataStatusIndex)-1) {
		return fmt.Sprintf("DataStatus(%d)", i)
	}
	return _DataStatusName[_DataStatusIndex[i]:_DataStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _DataStatusNoOp() {
	var x [1]struct{}
	_ = x[ValidDataStatus-(0)]
	_ = x[InvalidDataStatus-(1)]
}

var _DataStatusValues = []DataStatus{ValidDataStatus, InvalidDataStatus}

var _DataStatusNameToValueMap = map[string]DataStatus{
	_DataStatusName[0:1]:      ValidDataStatus,
	_DataStatusLowerName[0:1]: ValidDataStatus,
	_DataStatusName[1:2]:      InvalidDataStatus,
	_DataStatusLowerName[1:2]: InvalidDataStatus,
}

var _DataStatusNames = []string{
	_DataStatusName[0:1],
	_DataStatusName[1:2],
}

// DataStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DataStatusString(s string) (DataStatus, error) {
	if val, ok := _DataStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _DataStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DataStatus values", s)
}

// DataStatusValues returns all values of the enum
func DataStatusValues() []DataStatus {
	return _DataStatusValues
}

// DataStatusStrings returns a slice of all String values of the enum
func DataStatusStrings() []string {
	strs := make([]string, len(_DataStatusNames))
	copy(strs, _DataStatusNames)
	return strs
}

// IsADataStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DataStatus) IsADataStatus() bool {
	for _, v := range _DataStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for DataStatus
func (i DataStatus) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for DataStatus
func (i *DataStatus) UnmarshalText(text []byte) error {
	var err error
	*i, err = DataStatusString(string(text))
	return err
}

const _ModeName = "ADEM\""

var _ModeIndex = [...]uint8{0, 1, 2, 3, 4, 5}

const _ModeLowerName = "adem\""

func (i Mode) String() string {
	if i < 0 || i >= Mode(len(_ModeIndex)-1) {
		return fmt.Sprintf("Mode(%d)", i)
	}
	return _ModeName[_ModeIndex[i]:_ModeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ModeNoOp() {
	var x [1]struct{}
	_ = x[AutonomousMode-(0)]
	_ = x[DifferentialMode-(1)]
	_ = x[EstimatedMode-(2)]
	_ = x[ManualInputMode-(3)]
	_ = x[InvalidMode-(4)]
}

var _ModeValues = []Mode{AutonomousMode, DifferentialMode, EstimatedMode, ManualInputMode, InvalidMode}

var _ModeNameToValueMap = map[string]Mode{
	_ModeName[0:1]:      AutonomousMode,
	_ModeLowerName[0:1]: AutonomousMode,
	_ModeName[1:2]:      DifferentialMode,
	_ModeLowerName[1:2]: DifferentialMode,
	_ModeName[2:3]:      EstimatedMode,
	_ModeLowerName[2:3]: EstimatedMode,
	_ModeName[3:4]:      ManualInputMode,
	_ModeLowerName[3:4]: ManualInputMode,
	_ModeName[4:5]:      InvalidMode,
	_ModeLowerName[4:5]: InvalidMode,
}

var _ModeNames = []string{
	_ModeName[0:1],
	_ModeName[1:2],
	_ModeName[2:3],
	_ModeName[3:4],
	_ModeName[4:5],
}

// ModeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ModeString(s string) (Mode, error) {
	if val, ok := _ModeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ModeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Mode values", s)
}

// ModeValues returns all values of the enum
func ModeValues() []Mode {
	return _ModeValues
}

// ModeStrings returns a slice of all String values of the enum
func ModeStrings() []string {
	strs := make([]string, len(_ModeNames))
	copy(strs, _ModeNames)
	return strs
}

// IsAMode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Mode) IsAMode() bool {
	for _, v := range _ModeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Mode
func (i Mode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Mode
func (i *Mode) UnmarshalText(text []byte) error {
	var err error
	*i, err = ModeString(string(text))
	return err
}
