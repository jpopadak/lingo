// Code generated by Lingo for table information_schema.INNODB_SYS_FOREIGN_COLS - DO NOT EDIT

package qinnodbsysforeigncols

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbSysForeignCols {
	return instance
}

func Id() path.StringPath {
	return instance.id
}

func ForColName() path.StringPath {
	return instance.forColName
}

func RefColName() path.StringPath {
	return instance.refColName
}

func Pos() path.IntPath {
	return instance.pos
}
