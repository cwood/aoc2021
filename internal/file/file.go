package file

import (
	"io/ioutil"
	"strings"
)

func LoadAsStrings(filepath string) ([]string, error) {
	rawIn, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var o []string

	for _, s := range strings.Split(string(rawIn), "\n") {
		trimmed := strings.TrimSpace(s)

		if trimmed == "" {
			continue
		}
		o = append(o, trimmed)
	}

	return o, nil
}

func MustLoad(s []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return s
}
