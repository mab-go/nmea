// Package testhelp contains helper functions and data structures used by unit
// tests within the gopkg.in/mab-go/nmea.v0/sentence package.
package testhelp // import "gopkg.in/mab-go/nmea.v0/sentence/testhelp"

import (
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
)

// TestDatum represents a single datum from a set of test data.
type TestDatum struct {
	Name, Sentence, ActualChecksum, AdvertisedChecksum, ErrMsg string
}

// ReadTestData reads the contents of the specified test data set (a YAML file) into a slice of
// TestDatum structs. The file must have a ".yaml" extension and must be located in a directory
// named "testdata".
//
// Example:
//
//     // In 'foo/test_bar.go', read file 'foo/testdata/good/sentences.yaml':
//     testData := testhelp.ReadTestData("good/sentences")
//     for datum := range testData {
//         // Use test data
//     }
func ReadTestData(name string) []TestDatum {
	contents, err := ioutil.ReadFile("testdata/" + name + ".yaml")
	if err != nil {
		panic(err)
	}

	var target map[string]map[string]string
	err = yaml.Unmarshal(contents, &target)
	if err != nil {
		panic(err)
	}

	var result []TestDatum
	for k, v := range target {
		result = append(result, TestDatum{
			Name:               k,
			Sentence:           v["Sentence"],
			ActualChecksum:     v["ActualChecksum"],
			AdvertisedChecksum: v["AdvertisedChecksum"],
			ErrMsg:             v["ErrMsg"],
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}
