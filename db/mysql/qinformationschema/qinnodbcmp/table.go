// Code generated by Lingo for table information_schema.INNODB_CMP - DO NOT EDIT

package qinnodbcmp

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbCmp {
	return newQInnodbCmp(alias)
}

func New() QInnodbCmp {
	return newQInnodbCmp("")
}

func newQInnodbCmp(alias string) QInnodbCmp {
	q := QInnodbCmp{_alias: alias}
	q.pageSize = path.NewIntPath(q, "page_size")
	q.compressOps = path.NewIntPath(q, "compress_ops")
	q.compressOpsOk = path.NewIntPath(q, "compress_ops_ok")
	q.compressTime = path.NewIntPath(q, "compress_time")
	q.uncompressOps = path.NewIntPath(q, "uncompress_ops")
	q.uncompressTime = path.NewIntPath(q, "uncompress_time")
	return q
}

type QInnodbCmp struct {
	_alias         string
	pageSize       path.IntPath
	compressOps    path.IntPath
	compressOpsOk  path.IntPath
	compressTime   path.IntPath
	uncompressOps  path.IntPath
	uncompressTime path.IntPath
}

// core.Table Functions

func (q QInnodbCmp) GetColumns() []core.Column {
	return []core.Column{
		q.pageSize,
		q.compressOps,
		q.compressOpsOk,
		q.compressTime,
		q.uncompressOps,
		q.uncompressTime,
	}
}

func (q QInnodbCmp) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbCmp) GetAlias() string {
	return q._alias
}

func (q QInnodbCmp) GetName() string {
	return "INNODB_CMP"
}

func (q QInnodbCmp) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbCmp) PageSize() path.IntPath {
	return q.pageSize
}

func (q QInnodbCmp) CompressOps() path.IntPath {
	return q.compressOps
}

func (q QInnodbCmp) CompressOpsOk() path.IntPath {
	return q.compressOpsOk
}

func (q QInnodbCmp) CompressTime() path.IntPath {
	return q.compressTime
}

func (q QInnodbCmp) UncompressOps() path.IntPath {
	return q.uncompressOps
}

func (q QInnodbCmp) UncompressTime() path.IntPath {
	return q.uncompressTime
}
