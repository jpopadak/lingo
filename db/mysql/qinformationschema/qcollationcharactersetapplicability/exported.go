// Code generated by Lingo for table information_schema.COLLATION_CHARACTER_SET_APPLICABILITY - DO NOT EDIT

package qcollationcharactersetapplicability

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QCollationCharacterSetApplicability {
	return instance
}

func CollationName() path.StringPath {
	return instance.collationName
}

func CharacterSetName() path.StringPath {
	return instance.characterSetName
}
