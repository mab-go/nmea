// Code generated by "enumer -type=NorthSouth -text -sql -json -yaml -transform=first-upper -output=enum_northsouth_gen.go"; DO NOT EDIT.

package gpgll

import (
	"database/sql/driver"
	"encoding/json"
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

// MarshalJSON implements the json.Marshaler interface for NorthSouth
func (i NorthSouth) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for NorthSouth
func (i *NorthSouth) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NorthSouth should be a string, got %s", data)
	}

	var err error
	*i, err = NorthSouthString(s)
	return err
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

// MarshalYAML implements a YAML Marshaler for NorthSouth
func (i NorthSouth) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for NorthSouth
func (i *NorthSouth) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = NorthSouthString(s)
	return err
}

func (i NorthSouth) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *NorthSouth) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of NorthSouth: %[1]T(%[1]v)", value)
	}

	val, err := NorthSouthString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}