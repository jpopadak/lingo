// Code generated by Lingo for table information_schema.INNODB_TRX - DO NOT EDIT

package qinnodbtrx

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbTrx {
	return newQInnodbTrx(alias)
}

func New() QInnodbTrx {
	return newQInnodbTrx("")
}

func newQInnodbTrx(alias string) QInnodbTrx {
	q := QInnodbTrx{_alias: alias}
	q.trxId = path.NewStringPath(q, "trx_id")
	q.trxState = path.NewStringPath(q, "trx_state")
	q.trxStarted = path.NewTimePath(q, "trx_started")
	q.trxRequestedLockId = path.NewStringPath(q, "trx_requested_lock_id")
	q.trxWaitStarted = path.NewTimePath(q, "trx_wait_started")
	q.trxWeight = path.NewInt64Path(q, "trx_weight")
	q.trxMysqlThreadId = path.NewInt64Path(q, "trx_mysql_thread_id")
	q.trxQuery = path.NewStringPath(q, "trx_query")
	q.trxOperationState = path.NewStringPath(q, "trx_operation_state")
	q.trxTablesInUse = path.NewInt64Path(q, "trx_tables_in_use")
	q.trxTablesLocked = path.NewInt64Path(q, "trx_tables_locked")
	q.trxLockStructs = path.NewInt64Path(q, "trx_lock_structs")
	q.trxLockMemoryBytes = path.NewInt64Path(q, "trx_lock_memory_bytes")
	q.trxRowsLocked = path.NewInt64Path(q, "trx_rows_locked")
	q.trxRowsModified = path.NewInt64Path(q, "trx_rows_modified")
	q.trxConcurrencyTickets = path.NewInt64Path(q, "trx_concurrency_tickets")
	q.trxIsolationLevel = path.NewStringPath(q, "trx_isolation_level")
	q.trxUniqueChecks = path.NewIntPath(q, "trx_unique_checks")
	q.trxForeignKeyChecks = path.NewIntPath(q, "trx_foreign_key_checks")
	q.trxLastForeignKeyError = path.NewStringPath(q, "trx_last_foreign_key_error")
	q.trxAdaptiveHashLatched = path.NewIntPath(q, "trx_adaptive_hash_latched")
	q.trxAdaptiveHashTimeout = path.NewInt64Path(q, "trx_adaptive_hash_timeout")
	q.trxIsReadOnly = path.NewIntPath(q, "trx_is_read_only")
	q.trxAutocommitNonLocking = path.NewIntPath(q, "trx_autocommit_non_locking")
	return q
}

type QInnodbTrx struct {
	_alias                  string
	trxId                   path.StringPath
	trxState                path.StringPath
	trxStarted              path.TimePath
	trxRequestedLockId      path.StringPath
	trxWaitStarted          path.TimePath
	trxWeight               path.Int64Path
	trxMysqlThreadId        path.Int64Path
	trxQuery                path.StringPath
	trxOperationState       path.StringPath
	trxTablesInUse          path.Int64Path
	trxTablesLocked         path.Int64Path
	trxLockStructs          path.Int64Path
	trxLockMemoryBytes      path.Int64Path
	trxRowsLocked           path.Int64Path
	trxRowsModified         path.Int64Path
	trxConcurrencyTickets   path.Int64Path
	trxIsolationLevel       path.StringPath
	trxUniqueChecks         path.IntPath
	trxForeignKeyChecks     path.IntPath
	trxLastForeignKeyError  path.StringPath
	trxAdaptiveHashLatched  path.IntPath
	trxAdaptiveHashTimeout  path.Int64Path
	trxIsReadOnly           path.IntPath
	trxAutocommitNonLocking path.IntPath
}

// core.Table Functions

func (q QInnodbTrx) GetColumns() []core.Column {
	return []core.Column{
		q.trxId,
		q.trxState,
		q.trxStarted,
		q.trxRequestedLockId,
		q.trxWaitStarted,
		q.trxWeight,
		q.trxMysqlThreadId,
		q.trxQuery,
		q.trxOperationState,
		q.trxTablesInUse,
		q.trxTablesLocked,
		q.trxLockStructs,
		q.trxLockMemoryBytes,
		q.trxRowsLocked,
		q.trxRowsModified,
		q.trxConcurrencyTickets,
		q.trxIsolationLevel,
		q.trxUniqueChecks,
		q.trxForeignKeyChecks,
		q.trxLastForeignKeyError,
		q.trxAdaptiveHashLatched,
		q.trxAdaptiveHashTimeout,
		q.trxIsReadOnly,
		q.trxAutocommitNonLocking,
	}
}

func (q QInnodbTrx) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbTrx) GetAlias() string {
	return q._alias
}

func (q QInnodbTrx) GetName() string {
	return "INNODB_TRX"
}

func (q QInnodbTrx) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbTrx) TrxId() path.StringPath {
	return q.trxId
}

func (q QInnodbTrx) TrxState() path.StringPath {
	return q.trxState
}

func (q QInnodbTrx) TrxStarted() path.TimePath {
	return q.trxStarted
}

func (q QInnodbTrx) TrxRequestedLockId() path.StringPath {
	return q.trxRequestedLockId
}

func (q QInnodbTrx) TrxWaitStarted() path.TimePath {
	return q.trxWaitStarted
}

func (q QInnodbTrx) TrxWeight() path.Int64Path {
	return q.trxWeight
}

func (q QInnodbTrx) TrxMysqlThreadId() path.Int64Path {
	return q.trxMysqlThreadId
}

func (q QInnodbTrx) TrxQuery() path.StringPath {
	return q.trxQuery
}

func (q QInnodbTrx) TrxOperationState() path.StringPath {
	return q.trxOperationState
}

func (q QInnodbTrx) TrxTablesInUse() path.Int64Path {
	return q.trxTablesInUse
}

func (q QInnodbTrx) TrxTablesLocked() path.Int64Path {
	return q.trxTablesLocked
}

func (q QInnodbTrx) TrxLockStructs() path.Int64Path {
	return q.trxLockStructs
}

func (q QInnodbTrx) TrxLockMemoryBytes() path.Int64Path {
	return q.trxLockMemoryBytes
}

func (q QInnodbTrx) TrxRowsLocked() path.Int64Path {
	return q.trxRowsLocked
}

func (q QInnodbTrx) TrxRowsModified() path.Int64Path {
	return q.trxRowsModified
}

func (q QInnodbTrx) TrxConcurrencyTickets() path.Int64Path {
	return q.trxConcurrencyTickets
}

func (q QInnodbTrx) TrxIsolationLevel() path.StringPath {
	return q.trxIsolationLevel
}

func (q QInnodbTrx) TrxUniqueChecks() path.IntPath {
	return q.trxUniqueChecks
}

func (q QInnodbTrx) TrxForeignKeyChecks() path.IntPath {
	return q.trxForeignKeyChecks
}

func (q QInnodbTrx) TrxLastForeignKeyError() path.StringPath {
	return q.trxLastForeignKeyError
}

func (q QInnodbTrx) TrxAdaptiveHashLatched() path.IntPath {
	return q.trxAdaptiveHashLatched
}

func (q QInnodbTrx) TrxAdaptiveHashTimeout() path.Int64Path {
	return q.trxAdaptiveHashTimeout
}

func (q QInnodbTrx) TrxIsReadOnly() path.IntPath {
	return q.trxIsReadOnly
}

func (q QInnodbTrx) TrxAutocommitNonLocking() path.IntPath {
	return q.trxAutocommitNonLocking
}
