package kongfig

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// GetKongfigFromFile loads kong configuration generated with kongfig
func GetKongfigFromFile(filename string) (*Config, error) {
	if filename == "" {
		return nil, errors.New("filename cannot be empty")
	}
	fileContent, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}

func readFile(kongStateFile string) (*Config, error) {

	var s Config
	var err error
	var b []byte
	if kongStateFile == "-" {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		b, err = ioutil.ReadFile(kongStateFile)
	}
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
