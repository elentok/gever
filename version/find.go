package version

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type finder func(filename string) (Version, error)

func FindInVersionFile(filename string) (Version, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(string(contents))
}

type jsonData struct {
	Version *string `json:"version,omitempty"`
}

func FindInPackageJSON(filename string) (Version, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := jsonData{}
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return nil, err
	}

	if data.Version == nil {
		return nil, errors.New("Missing 'version'")
	}

	return Parse(*data.Version)
}

type yamlData struct {
	Major   int    `yaml:":major"`
	Minor   int    `yaml:":minor"`
	Patch   int    `yaml:":patch"`
	Special string `yaml:":special"`
}

func FindInYaml(filename string) (Version, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := yamlData{}
	err = yaml.Unmarshal(contents, &data)

	if err != nil {
		return nil, err
	}

	v := New(data.Major, data.Minor, data.Patch, data.Special)
	return v, nil
}

func Find(directory string, verbose bool) (Version, error) {

	printIfVerbose("Searching for version in .semver", verbose)
	v, err := FindInYaml(".semver")
	if err == nil {
		return v, nil
	}

	printIfVerbose("Searching for version in 'version'", verbose)
	v, err = FindInVersionFile("version")
	if err == nil {
		return v, nil
	}

	printIfVerbose("Searching for version in 'VERSION'", verbose)
	v, err = FindInVersionFile("VERSION")
	if err == nil {
		return v, nil
	}

	printIfVerbose("Searching for version in package.json", verbose)
	v, err = FindInPackageJSON("package.json")
	if err == nil {
		return v, nil
	}

	return v, err
}

func printIfVerbose(text string, verbose bool) {
	if verbose {
		println(text)
	}
}
