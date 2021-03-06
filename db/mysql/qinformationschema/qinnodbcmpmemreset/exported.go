// Code generated by Lingo for table information_schema.INNODB_CMPMEM_RESET - DO NOT EDIT

package qinnodbcmpmemreset

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbCmpmemReset {
	return instance
}

func PageSize() path.IntPath {
	return instance.pageSize
}

func BufferPoolInstance() path.IntPath {
	return instance.bufferPoolInstance
}

func PagesUsed() path.IntPath {
	return instance.pagesUsed
}

func PagesFree() path.IntPath {
	return instance.pagesFree
}

func RelocationOps() path.Int64Path {
	return instance.relocationOps
}

func RelocationTime() path.IntPath {
	return instance.relocationTime
}
