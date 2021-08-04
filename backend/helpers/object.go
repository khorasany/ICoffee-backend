package helpers

import (
	"encoding/json"
	"strings"
)

func ConvertObjectToString(items struct{}) (string, error) {

	b, err := json.Marshal(items)
	if err != nil {
		return "", err
	}

	s := string(b)
	s = strings.TrimSpace(s)
	// trim leading and trailing spaces
	s = s[1 : len(s)-1]
	return s, nil
}
