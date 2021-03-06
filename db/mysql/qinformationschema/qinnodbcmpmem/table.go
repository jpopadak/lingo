// Code generated by Lingo for table information_schema.INNODB_CMPMEM - DO NOT EDIT

package qinnodbcmpmem

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbCmpmem {
	return newQInnodbCmpmem(alias)
}

func New() QInnodbCmpmem {
	return newQInnodbCmpmem("")
}

func newQInnodbCmpmem(alias string) QInnodbCmpmem {
	q := QInnodbCmpmem{_alias: alias}
	q.pageSize = path.NewIntPath(q, "page_size")
	q.bufferPoolInstance = path.NewIntPath(q, "buffer_pool_instance")
	q.pagesUsed = path.NewIntPath(q, "pages_used")
	q.pagesFree = path.NewIntPath(q, "pages_free")
	q.relocationOps = path.NewInt64Path(q, "relocation_ops")
	q.relocationTime = path.NewIntPath(q, "relocation_time")
	return q
}

type QInnodbCmpmem struct {
	_alias             string
	pageSize           path.IntPath
	bufferPoolInstance path.IntPath
	pagesUsed          path.IntPath
	pagesFree          path.IntPath
	relocationOps      path.Int64Path
	relocationTime     path.IntPath
}

// core.Table Functions

func (q QInnodbCmpmem) GetColumns() []core.Column {
	return []core.Column{
		q.pageSize,
		q.bufferPoolInstance,
		q.pagesUsed,
		q.pagesFree,
		q.relocationOps,
		q.relocationTime,
	}
}

func (q QInnodbCmpmem) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbCmpmem) GetAlias() string {
	return q._alias
}

func (q QInnodbCmpmem) GetName() string {
	return "INNODB_CMPMEM"
}

func (q QInnodbCmpmem) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbCmpmem) PageSize() path.IntPath {
	return q.pageSize
}

func (q QInnodbCmpmem) BufferPoolInstance() path.IntPath {
	return q.bufferPoolInstance
}

func (q QInnodbCmpmem) PagesUsed() path.IntPath {
	return q.pagesUsed
}

func (q QInnodbCmpmem) PagesFree() path.IntPath {
	return q.pagesFree
}

func (q QInnodbCmpmem) RelocationOps() path.Int64Path {
	return q.relocationOps
}

func (q QInnodbCmpmem) RelocationTime() path.IntPath {
	return q.relocationTime
}
