// Code generated by Lingo for table information_schema.INNODB_VIRTUAL - DO NOT EDIT

package qinnodbvirtual

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QInnodbVirtual {
	return instance
}

func TableId() path.Int64Path {
	return instance.tableId
}

func Pos() path.IntPath {
	return instance.pos
}

func BasePos() path.IntPath {
	return instance.basePos
}