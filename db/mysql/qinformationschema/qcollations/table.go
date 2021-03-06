// Code generated by Lingo for table information_schema.COLLATIONS - DO NOT EDIT

package qcollations

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QCollations {
	return newQCollations(alias)
}

func New() QCollations {
	return newQCollations("")
}

func newQCollations(alias string) QCollations {
	q := QCollations{_alias: alias}
	q.collationName = path.NewStringPath(q, "COLLATION_NAME")
	q.characterSetName = path.NewStringPath(q, "CHARACTER_SET_NAME")
	q.id = path.NewInt64Path(q, "ID")
	q.isDefault = path.NewStringPath(q, "IS_DEFAULT")
	q.isCompiled = path.NewStringPath(q, "IS_COMPILED")
	q.sortlen = path.NewInt64Path(q, "SORTLEN")
	return q
}

type QCollations struct {
	_alias           string
	collationName    path.StringPath
	characterSetName path.StringPath
	id               path.Int64Path
	isDefault        path.StringPath
	isCompiled       path.StringPath
	sortlen          path.Int64Path
}

// core.Table Functions

func (q QCollations) GetColumns() []core.Column {
	return []core.Column{
		q.collationName,
		q.characterSetName,
		q.id,
		q.isDefault,
		q.isCompiled,
		q.sortlen,
	}
}

func (q QCollations) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QCollations) GetAlias() string {
	return q._alias
}

func (q QCollations) GetName() string {
	return "COLLATIONS"
}

func (q QCollations) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QCollations) CollationName() path.StringPath {
	return q.collationName
}

func (q QCollations) CharacterSetName() path.StringPath {
	return q.characterSetName
}

func (q QCollations) Id() path.Int64Path {
	return q.id
}

func (q QCollations) IsDefault() path.StringPath {
	return q.isDefault
}

func (q QCollations) IsCompiled() path.StringPath {
	return q.isCompiled
}

func (q QCollations) Sortlen() path.Int64Path {
	return q.sortlen
}
