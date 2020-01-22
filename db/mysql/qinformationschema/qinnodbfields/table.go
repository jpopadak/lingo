// Code generated by Lingo for table information_schema.INNODB_FIELDS - DO NOT EDIT

package qinnodbfields

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QInnodbFields {
	return newQInnodbFields(alias)
}

func New() QInnodbFields {
	return newQInnodbFields("")
}

func newQInnodbFields(alias string) QInnodbFields {
	q := QInnodbFields{_alias: alias}
	q.indexId = NewUnknownPathType(q, "INDEX_ID")
	q.name = path.NewStringPath(q, "NAME")
	q.pos = path.NewInt64Path(q, "POS")
	return q
}

type QInnodbFields struct {
	_alias  string
	indexId UnknownPathType
	name    path.StringPath
	pos     path.Int64Path
}

// core.Table Functions

func (q QInnodbFields) GetColumns() []core.Column {
	return []core.Column{
		q.indexId,
		q.name,
		q.pos,
	}
}

func (q QInnodbFields) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QInnodbFields) GetAlias() string {
	return q._alias
}

func (q QInnodbFields) GetName() string {
	return "INNODB_FIELDS"
}

func (q QInnodbFields) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbFields) IndexId() UnknownPathType {
	return q.indexId
}

func (q QInnodbFields) Name() path.StringPath {
	return q.name
}

func (q QInnodbFields) Pos() path.Int64Path {
	return q.pos
}