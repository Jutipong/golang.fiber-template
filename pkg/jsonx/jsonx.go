package jsonx

import "github.com/goccy/go-json"

func Marshal[T any](v T) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal[T any](data []byte, v T) error {
	return json.Unmarshal(data, v)
}

func MarshalIndent[T any](v T, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
