// Package testhelp contains helper functions and data structures used by unit
// tests within the gopkg.in/mab-go/nmea.v0/sentence package.
package testhelp // import "gopkg.in/mab-go/nmea.v0/sentence/testhelp"

import (
	"io/ioutil"
	"path"
	"sort"

	"gopkg.in/yaml.v2"
)

// ReadTestData reads the contents of the specified test data set (a YAML file) into a slice of
// some type defined by the return value of mapFn. The file must have a ".yaml" extension and must
// be located in a directory named "_testdata" relative to the caller.
//
// Example:
//
//     // In 'foo/test_bar.go', read file 'foo/_testdata/good/sentences.yaml':
//     goodData := testhelp.ReadTestData("good/sentences", mapParseInput, sortParseInput)
//	   for _, data := range goodData {
//	       d := data.(myInputType)
//         // Use test data...
//     }
func ReadTestData(
	name string,
	mapFn func(title string, input map[string]string) interface{},
	sortFn func(result []interface{}, i, j int) bool,
) []interface{} {
	contents, err := ioutil.ReadFile(path.Join("_testdata/", name+".yaml"))
	if err != nil {
		panic(err)
	}

	// input is a map of test case titles to test case input objects
	var input map[string]map[string]string
	err = yaml.Unmarshal(contents, &input)
	if err != nil {
		panic(err)
	}

	var result []interface{}
	for k, v := range input {
		result = append(result, mapFn(k, v))
	}

	sort.Slice(result, func(i, j int) bool {
		return sortFn(result, i, j)
	})

	return result
}
