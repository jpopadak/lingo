// Code generated by Lingo for table information_schema.COLLATION_CHARACTER_SET_APPLICABILITY - DO NOT EDIT

package qcollationcharactersetapplicability

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QCollationCharacterSetApplicability {
	return newQCollationCharacterSetApplicability(alias)
}

func New() QCollationCharacterSetApplicability {
	return newQCollationCharacterSetApplicability("")
}

func newQCollationCharacterSetApplicability(alias string) QCollationCharacterSetApplicability {
	q := QCollationCharacterSetApplicability{_alias: alias}
	q.collationName = path.NewStringPath(q, "COLLATION_NAME")
	q.characterSetName = path.NewStringPath(q, "CHARACTER_SET_NAME")
	return q
}

type QCollationCharacterSetApplicability struct {
	_alias           string
	collationName    path.StringPath
	characterSetName path.StringPath
}

// core.Table Functions

func (q QCollationCharacterSetApplicability) GetColumns() []core.Column {
	return []core.Column{
		q.collationName,
		q.characterSetName,
	}
}

func (q QCollationCharacterSetApplicability) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QCollationCharacterSetApplicability) GetAlias() string {
	return q._alias
}

func (q QCollationCharacterSetApplicability) GetName() string {
	return "COLLATION_CHARACTER_SET_APPLICABILITY"
}

func (q QCollationCharacterSetApplicability) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QCollationCharacterSetApplicability) CollationName() path.StringPath {
	return q.collationName
}

func (q QCollationCharacterSetApplicability) CharacterSetName() path.StringPath {
	return q.characterSetName
}
