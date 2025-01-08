package validator

import (
	"fmt"
	"reflect"
	"strings"
)

type Validatable interface {
	Validate() error
}

type ValidationError struct {
	Path string
	Err  error
}

func (ve ValidationError) Error() string {
	if ve.Path == "" {
		return ve.Err.Error()
	}
	return fmt.Sprintf("%s: %v", ve.Path, ve.Err)
}

type ValidationErrors []error

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return ""
	}
	var sb strings.Builder
	for i, err := range ve {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

func Validate(obj any) error {
	if obj == nil {
		return nil
	}

	errs := validateRecursive(reflect.ValueOf(obj), "")
	if len(errs) > 0 {
		return ValidationErrors(errs)
	}
	return nil
}

func validateRecursive(v reflect.Value, path string) []error {
	var errs []error

	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.CanInterface() {
		validatableType := reflect.TypeOf((*Validatable)(nil)).Elem()
		t := v.Type()

		switch {
		case t.Implements(validatableType):
			if val, ok := v.Interface().(Validatable); ok {
				if err := val.Validate(); err != nil {
					errs = append(errs, ValidationError{Path: path, Err: err})
				}
			}

		case v.CanAddr() && reflect.PointerTo(t).Implements(validatableType):
			if val, ok := v.Addr().Interface().(Validatable); ok {
				if err := val.Validate(); err != nil {
					errs = append(errs, ValidationError{Path: path, Err: err})
				}
			}
		}
	}

	switch v.Kind() {
	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if !v.Field(i).CanInterface() {
				continue
			}
			fieldVal := v.Field(i)
			fieldType := t.Field(i)
			fieldPath := fieldType.Name
			if path != "" {
				fieldPath = path + "." + fieldType.Name
			}
			errs = append(errs, validateRecursive(fieldVal, fieldPath)...)
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			itemVal := v.Index(i)
			itemPath := fmt.Sprintf("%s[%d]", path, i)
			errs = append(errs, validateRecursive(itemVal, itemPath)...)
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			itemPath := fmt.Sprintf("%s[%v]", path, key.Interface())
			errs = append(errs, validateRecursive(val, itemPath)...)
		}
	}

	return errs
}
