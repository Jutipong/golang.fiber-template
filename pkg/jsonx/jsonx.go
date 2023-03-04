package jsonx

import "github.com/bytedance/sonic"

// import "github.com/goccy/go-json"

func Marshal[T any](v T) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal[T any](data []byte, v T) error {
	return sonic.Unmarshal(data, v)
}

// func MarshalIndent[T any](v T, prefix string, indent string) ([]byte, error) {
// 	return sonic.mash(v, prefix, indent)
// }
