package util

import (
	"errors"
	"reflect"
)

const (
	JSON = "json"
	BSON = "bson"
)

func GetReflectTagField(field interface{}, format string) (string, error) {

	switch format {
	case JSON:
		return reflect.TypeOf(field).Field(0).Tag.Get(JSON), nil
	case BSON:
		return reflect.TypeOf(field).Field(0).Tag.Get(BSON), nil
	default:
		return "", errors.New("input formant not found")
	}
}
