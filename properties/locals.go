package properties

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// LocalProperties are read from a file
type PropertiesMap map[string]string

// LoadLocalProperties reads a file and just loads it
func LoadLocalProperties(path string) (PropertiesMap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// result, returned anyway
	result := make(PropertiesMap)
	// global error, may be null
	var globalError error

	for {
		// read current line
		line, err := reader.ReadString('\n')
		// find equals and then split
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				// assign the config map
				result[key] = value
			}
		}

		if err == io.EOF {
			break
		} else if err != nil {
			globalError = errors.Join(globalError, err)
		}
	}

	return result, globalError
}
