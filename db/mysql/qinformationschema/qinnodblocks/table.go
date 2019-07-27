// Code generated by Lingo for table information_schema.INNODB_LOCKS - DO NOT EDIT

package qinnodblocks

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbLocks {
	return newQInnodbLocks(alias)
}

func New() QInnodbLocks {
	return newQInnodbLocks("")
}

func newQInnodbLocks(alias string) QInnodbLocks {
	q := QInnodbLocks{_alias: alias}
	q.lockId = path.NewStringPath(q, "lock_id")
	q.lockTrxId = path.NewStringPath(q, "lock_trx_id")
	q.lockMode = path.NewStringPath(q, "lock_mode")
	q.lockType = path.NewStringPath(q, "lock_type")
	q.lockTable = path.NewStringPath(q, "lock_table")
	q.lockIndex = path.NewStringPath(q, "lock_index")
	q.lockSpace = path.NewInt64Path(q, "lock_space")
	q.lockPage = path.NewInt64Path(q, "lock_page")
	q.lockRec = path.NewInt64Path(q, "lock_rec")
	q.lockData = path.NewStringPath(q, "lock_data")
	return q
}

type QInnodbLocks struct {
	_alias    string
	lockId    path.StringPath
	lockTrxId path.StringPath
	lockMode  path.StringPath
	lockType  path.StringPath
	lockTable path.StringPath
	lockIndex path.StringPath
	lockSpace path.Int64Path
	lockPage  path.Int64Path
	lockRec   path.Int64Path
	lockData  path.StringPath
}

// core.Table Functions

func (q QInnodbLocks) GetColumns() []core.Column {
	return []core.Column{
		q.lockId,
		q.lockTrxId,
		q.lockMode,
		q.lockType,
		q.lockTable,
		q.lockIndex,
		q.lockSpace,
		q.lockPage,
		q.lockRec,
		q.lockData,
	}
}

func (q QInnodbLocks) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbLocks) GetAlias() string {
	return q._alias
}

func (q QInnodbLocks) GetName() string {
	return "INNODB_LOCKS"
}

func (q QInnodbLocks) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbLocks) LockId() path.StringPath {
	return q.lockId
}

func (q QInnodbLocks) LockTrxId() path.StringPath {
	return q.lockTrxId
}

func (q QInnodbLocks) LockMode() path.StringPath {
	return q.lockMode
}

func (q QInnodbLocks) LockType() path.StringPath {
	return q.lockType
}

func (q QInnodbLocks) LockTable() path.StringPath {
	return q.lockTable
}

func (q QInnodbLocks) LockIndex() path.StringPath {
	return q.lockIndex
}

func (q QInnodbLocks) LockSpace() path.Int64Path {
	return q.lockSpace
}

func (q QInnodbLocks) LockPage() path.Int64Path {
	return q.lockPage
}

func (q QInnodbLocks) LockRec() path.Int64Path {
	return q.lockRec
}

func (q QInnodbLocks) LockData() path.StringPath {
	return q.lockData
}