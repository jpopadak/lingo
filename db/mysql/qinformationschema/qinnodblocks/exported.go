// Code generated by Lingo for table information_schema.INNODB_LOCKS - DO NOT EDIT

package qinnodblocks

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbLocks {
	return instance
}

func LockId() path.StringPath {
	return instance.lockId
}

func LockTrxId() path.StringPath {
	return instance.lockTrxId
}

func LockMode() path.StringPath {
	return instance.lockMode
}

func LockType() path.StringPath {
	return instance.lockType
}

func LockTable() path.StringPath {
	return instance.lockTable
}

func LockIndex() path.StringPath {
	return instance.lockIndex
}

func LockSpace() path.Int64Path {
	return instance.lockSpace
}

func LockPage() path.Int64Path {
	return instance.lockPage
}

func LockRec() path.Int64Path {
	return instance.lockRec
}

func LockData() path.StringPath {
	return instance.lockData
}
