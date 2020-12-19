package context_error

import (
	"fmt"
	"github.com/magiconair/properties"
	"github.com/o-z/featuretoggle/base/base_model"
	"path/filepath"
)

var Properties *properties.Properties

func GetContextError(errorName string) base_model.ContextError {
	prefix := Properties.FilterPrefix(errorName)
	errorCode, errorCodeOk := prefix.Get(errorName + ".code")
	if errorCodeOk != true {
		fmt.Println("ErrorCode not fiend in property file ")
	}
	errorMessage, errorMessageOk := prefix.Get(errorName + ".message")
	if errorMessageOk != true {
		fmt.Println("ErrorMessage not fiend in property file ")
	}
	var contextError base_model.ContextError
	contextError = base_model.ContextError{
		ErrorCode: errorCode,
		ErrorDesc: errorMessage,
	}
	return contextError
}

func SetAllProperties(propertiesDirectory string) {
	propertiesFileList, _ := filepath.Glob(propertiesDirectory)
	Properties, _ = properties.LoadFiles(propertiesFileList, properties.UTF8, true)

}
