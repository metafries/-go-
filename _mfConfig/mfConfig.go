package _mfConfig

import (
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
	JSON
	XML
)

var wrongTypeError error = errors.New("[error] Type must be a pointer to a struct")

func GetConfiguration(configType uint8, obj interface{}, fileName string) (err error) {
	// Check if this is a type pointer
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return wrongTypeError
	}
	// Confirm the struct value
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}
	switch configType {
	case CUSTOM:
		err = MarshalCustomConfig(mysRValue, fileName)
	case JSON:
		err = decodeJSONConfig(obj, fileName)
	case XML:
		err = decodeXMLConfig(obj, fileName)
	}
	return err
}
