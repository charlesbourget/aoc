package utils

import "errors"

type Language string

const (
	Go   Language = "go"
	Rust Language = "rust"
)

func (l *Language) String() string {
	return string(*l)
}

func (e *Language) Set(v string) error {
	switch v {
	case "go", "rust":
		*e = Language(v)
		return nil
	default:
		return errors.New(`must be one of "go", "rust"`)
	}
}

func (e *Language) Type() string {
	return "Language"
}
