// Code generated by Lingo for table information_schema.INNODB_LOCK_WAITS - DO NOT EDIT

package qinnodblockwaits

import "github.com/weworksandbox/lingo/pkg/core/path"

var instance = New()

func Q() QInnodbLockWaits {
	return instance
}

func RequestingTrxId() path.StringPath {
	return instance.requestingTrxId
}

func RequestedLockId() path.StringPath {
	return instance.requestedLockId
}

func BlockingTrxId() path.StringPath {
	return instance.blockingTrxId
}

func BlockingLockId() path.StringPath {
	return instance.blockingLockId
}
