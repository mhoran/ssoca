package config

import (
	"errors"
	"fmt"
)

type EnvironmentOption interface {
	Key() string
	SetValue(interface{}) error
}

type StringEnvironmentOption struct {
	key   string
	value string
}

var _ EnvironmentOption = &StringEnvironmentOption{}

func NewStringEnvironmentOption(key string) StringEnvironmentOption {
	return StringEnvironmentOption{key: key}
}

func (eo *StringEnvironmentOption) Key() string {
	return eo.key
}

func (eo *StringEnvironmentOption) SetValue(value interface{}) error {
	stringValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("Cannot convert option value to string: %#+v", value)
	}

	eo.value = stringValue

	return nil
}

func (eo *StringEnvironmentOption) GetValue() string {
	return eo.value
}

type StringSliceEnvironmentOption struct {
	key   string
	value []string
}

var _ EnvironmentOption = &StringSliceEnvironmentOption{}

func NewStringSliceEnvironmentOption(key string) StringSliceEnvironmentOption {
	return StringSliceEnvironmentOption{key: key}
}

func (eo *StringSliceEnvironmentOption) Key() string {
	return eo.key
}

func (eo *StringSliceEnvironmentOption) SetValue(value interface{}) error {
	if stringSliceValue, ok := value.([]string); ok {
		eo.value = stringSliceValue

		return nil
	}

	sliceValue, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("Cannot convert option value to slice: %#+v", value)
	}

	eo.value = []string{}

	for _, itemValue := range sliceValue {
		stringItemValue, ok := itemValue.(string)
		if !ok {
			return errors.New("Cannot convert option slice value item to string")
		}

		eo.value = append(eo.value, stringItemValue)
	}

	return nil
}

func (eo *StringSliceEnvironmentOption) GetValue() []string {
	return eo.value
}
