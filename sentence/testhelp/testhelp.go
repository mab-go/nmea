// Package testhelp contains helper functions and data structures used by unit
// tests within the gopkg.in/mab-go/nmea.v0/sentence package.
package testhelp // import "gopkg.in/mab-go/nmea.v0/sentence/testhelp"

import (
	"io/ioutil"
	"path"
	"sort"

	"gopkg.in/yaml.v2"
)

// EnsureFloat32 attempts to cast v as a float64, then converts it to a float32. If the cast fails,
// a runtime panic occurs.
func EnsureFloat32(v interface{}) float32 {
	f := EnsureFloat64(v)

	return float32(f)
}

// OptFloat32 attempts to cast v as a float64, then converts it to a float32. If the cast fails,
// it returns the type's zero-value (0).
func OptFloat32(v interface{}) float32 {
	f := OptFloat64(v)

	return float32(f)
}

// EnsureFloat64 attempts to cast v as a float64. If the cast fails, a runtime panic occurs.
func EnsureFloat64(v interface{}) float64 {
	return v.(float64)
}

// OptFloat64 attempts to cast v as a float64. If the cast fails, it returns the type's zero-value (0).
func OptFloat64(v interface{}) float64 {
	if f, ok := v.(float64); ok {
		return f
	}

	return 0
}

// EnsureInt attempts to cast v as an int. If the cast fails, a runtime panic occurs.
func EnsureInt(v interface{}) int {
	return v.(int)
}

// OptInt attempts to cast v as an int. If the cast fails, it returns the type's zero-value (0).
func OptInt(v interface{}) int {
	if i, ok := v.(int); ok {
		return i
	}

	return 0
}

// EnsureInt8 attempts to cast v as an int, then converts it to an int8. If the cast fails,
// a runtime panic occurs.
func EnsureInt8(v interface{}) int8 {
	i := EnsureInt(v)

	return int8(i)
}

// OptInt8 attempts to cast v as an int, then converts it to an int8. If the cast fails, it returns
// the type's zero-value (0).
func OptInt8(v interface{}) int8 {
	i := OptInt(v)

	return int8(i)
}

// EnsureInt16 attempts to cast v as an int, then converts it to an int16. If the cast fails,
// a runtime panic occurs.
func EnsureInt16(v interface{}) int16 {
	i := EnsureInt(v)

	return int16(i)
}

// OptInt16 attempts to cast v as an int, then converts it to an int16. If the cast fails, it returns
// the type's zero-value (0).
func OptInt16(v interface{}) int16 {
	i := OptInt(v)

	return int16(i)
}

// EnsureString attempts to cast v as a string. If the cast fails, a runtime panic occurs.
func EnsureString(v interface{}) string {
	return v.(string)
}

// OptString attempts to cast v as a string. If the cast fails, it returns
// the type's zero-value ("").
func OptString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}

	return ""
}

// ReadTestData reads the contents of the specified test data set (a YAML file) into a slice of
// some type defined by the return value of mapFn. The file must have a ".yaml" extension and must
// be located in a directory named "_testdata" relative to the caller.
//
// Example:
//
//     // In 'foo/test_bar.go', read file 'foo/_testdata/good/sentences.yaml':
//     type myInputType struct { /* ... */ }
//     goodData := testhelp.ReadTestData("good/sentences", mapInput, sortInput)
//	   for _, d := range goodData {
//	       expected := d.(myInputType)
//         // Use expected value in test...
//     }
func ReadTestData(
	name string,
	mapFn func(title string, input map[string]interface{}) interface{},
	sortFn func(result []interface{}, i, j int) bool,
) []interface{} {
	contents, err := ioutil.ReadFile(path.Join("_testdata/", name+".yaml"))
	if err != nil {
		panic(err)
	}

	// input is a map of test case titles to test case input objects
	var input map[string]map[string]interface{}
	err = yaml.Unmarshal(contents, &input)
	if err != nil {
		panic(err)
	}

	result := make([]interface{}, len(input))
	for k, v := range input {
		result = append(result, mapFn(k, v))
	}

	sort.Slice(result, func(i, j int) bool {
		return sortFn(result, i, j)
	})

	return result
}
