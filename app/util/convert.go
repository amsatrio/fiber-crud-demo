package util

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
)

func CamelCaseToSnakeCase(input string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(input, "${1}_${2}")

	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func ResponseToByte(response response.Response) ([]byte, error) {

	resBytes := new(bytes.Buffer)
	err := json.NewEncoder(resBytes).Encode(response)
	if err != nil {
		return nil, err
	}

	return resBytes.Bytes(), nil
}
