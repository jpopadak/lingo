// Code generated by Lingo for table information_schema.SESSION_VARIABLES - DO NOT EDIT

package qsessionvariables

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QSessionVariables {
	return instance
}

func VariableName() path.StringPath {
	return instance.variableName
}

func VariableValue() path.StringPath {
	return instance.variableValue
}
