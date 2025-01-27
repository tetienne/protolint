package config

import (
	"fmt"
)

// ImportsSortedOption represents the option for the IMPORTS_SORTED rule.
type ImportsSortedOption struct {
	// Deprecated: not used
	Newline string
}

// UnmarshalYAML implements yaml.v2 Unmarshaler interface.
func (i *ImportsSortedOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var option struct {
		Newline string `yaml:"newline"`
	}
	if err := unmarshal(&option); err != nil {
		return err
	}

	switch option.Newline {
	case "\n", "\r", "\r\n", "":
		i.Newline = option.Newline
	default:
		return fmt.Errorf(`%s is an invalid newline option. valid option is \n, \r or \r\n`, option.Newline)
	}
	return nil
}
