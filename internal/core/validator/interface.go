package validator

import (
	"reflect"
)

type Validatable interface {
	Validate() error
}

func Validate(data Validatable) error {
	validatable := reflect.TypeOf((*Validatable)(nil)).Elem()
	e := reflect.ValueOf(data).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Type.Implements(validatable) {
			result := e.Field(i).MethodByName("Validate").Call(nil)
			if result != nil {
				return result[0].Interface().(error)
			}
		}
	}
	return nil
}
