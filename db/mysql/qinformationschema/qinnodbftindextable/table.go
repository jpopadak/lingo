// Code generated by Lingo for table information_schema.INNODB_FT_INDEX_TABLE - DO NOT EDIT

package qinnodbftindextable

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbFtIndexTable {
	return newQInnodbFtIndexTable(alias)
}

func New() QInnodbFtIndexTable {
	return newQInnodbFtIndexTable("")
}

func newQInnodbFtIndexTable(alias string) QInnodbFtIndexTable {
	q := QInnodbFtIndexTable{_alias: alias}
	q.word = path.NewStringPath(q, "WORD")
	q.firstDocId = path.NewInt64Path(q, "FIRST_DOC_ID")
	q.lastDocId = path.NewInt64Path(q, "LAST_DOC_ID")
	q.docCount = path.NewInt64Path(q, "DOC_COUNT")
	q.docId = path.NewInt64Path(q, "DOC_ID")
	q.position = path.NewInt64Path(q, "POSITION")
	return q
}

type QInnodbFtIndexTable struct {
	_alias     string
	word       path.StringPath
	firstDocId path.Int64Path
	lastDocId  path.Int64Path
	docCount   path.Int64Path
	docId      path.Int64Path
	position   path.Int64Path
}

// core.Table Functions

func (q QInnodbFtIndexTable) GetColumns() []core.Column {
	return []core.Column{
		q.word,
		q.firstDocId,
		q.lastDocId,
		q.docCount,
		q.docId,
		q.position,
	}
}

func (q QInnodbFtIndexTable) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbFtIndexTable) GetAlias() string {
	return q._alias
}

func (q QInnodbFtIndexTable) GetName() string {
	return "INNODB_FT_INDEX_TABLE"
}

func (q QInnodbFtIndexTable) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbFtIndexTable) Word() path.StringPath {
	return q.word
}

func (q QInnodbFtIndexTable) FirstDocId() path.Int64Path {
	return q.firstDocId
}

func (q QInnodbFtIndexTable) LastDocId() path.Int64Path {
	return q.lastDocId
}

func (q QInnodbFtIndexTable) DocCount() path.Int64Path {
	return q.docCount
}

func (q QInnodbFtIndexTable) DocId() path.Int64Path {
	return q.docId
}

func (q QInnodbFtIndexTable) Position() path.Int64Path {
	return q.position
}
