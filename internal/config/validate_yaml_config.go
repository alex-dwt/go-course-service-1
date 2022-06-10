package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func validateYamlConfig(cfg YamlConfig) error {
	errs := validator.New().Struct(cfg)
	if errs == nil {
		return nil
	}

	var text string

	for i, err := range errs.(validator.ValidationErrors) {
		text += fmt.Sprintf(
			"Error #%d -- Namespace: %v, Field: %v, StructNamespace: %v, StructField: %v, Tag: %v, ActualTag: %v, Kind: %v, Type: %v, Value: %v, Param: %v\n",
			i,
			err.Namespace(),
			err.Field(),
			err.StructNamespace(),
			err.StructField(),
			err.Tag(),
			err.ActualTag(),
			err.Kind(),
			err.Type(),
			err.Value(),
			err.Param(),
		)
	}

	return errors.New(text)
}
