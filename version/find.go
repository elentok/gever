package version

import "io/ioutil"

func FindInVersionFile(filename string) (Version, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	v, err := Parse(string(contents))
	if err != nil {
		return nil, err
	}

	return v, nil
}
