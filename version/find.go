package version

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func FindInVersionFile(filename string) (Version, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(string(contents))
}

type packageJson struct {
	Version *string `json:"version,omitempty"`
}

func FindInPackageJSON(filename string) (Version, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	theJson := packageJson{}
	err = json.Unmarshal(contents, &theJson)
	if err != nil {
		return nil, err
	}

	if theJson.Version == nil {
		return nil, errors.New("Missing 'version'")
	}

	return Parse(*theJson.Version)
}
