// Code generated by Lingo for table information_schema.INNODB_SESSION_TEMP_TABLESPACES - DO NOT EDIT

package qinnodbsessiontemptablespaces

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QInnodbSessionTempTablespaces {
	return newQInnodbSessionTempTablespaces(alias)
}

func New() QInnodbSessionTempTablespaces {
	return newQInnodbSessionTempTablespaces("")
}

func newQInnodbSessionTempTablespaces(alias string) QInnodbSessionTempTablespaces {
	q := QInnodbSessionTempTablespaces{_alias: alias}
	q.id = path.NewIntPath(q, "ID")
	q.space = path.NewIntPath(q, "SPACE")
	q.path = path.NewStringPath(q, "PATH")
	q.size = path.NewInt64Path(q, "SIZE")
	q.state = path.NewStringPath(q, "STATE")
	q.purpose = path.NewStringPath(q, "PURPOSE")
	return q
}

type QInnodbSessionTempTablespaces struct {
	_alias  string
	id      path.IntPath
	space   path.IntPath
	path    path.StringPath
	size    path.Int64Path
	state   path.StringPath
	purpose path.StringPath
}

// core.Table Functions

func (q QInnodbSessionTempTablespaces) GetColumns() []core.Column {
	return []core.Column{
		q.id,
		q.space,
		q.path,
		q.size,
		q.state,
		q.purpose,
	}
}

func (q QInnodbSessionTempTablespaces) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QInnodbSessionTempTablespaces) GetAlias() string {
	return q._alias
}

func (q QInnodbSessionTempTablespaces) GetName() string {
	return "INNODB_SESSION_TEMP_TABLESPACES"
}

func (q QInnodbSessionTempTablespaces) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbSessionTempTablespaces) Id() path.IntPath {
	return q.id
}

func (q QInnodbSessionTempTablespaces) Space() path.IntPath {
	return q.space
}

func (q QInnodbSessionTempTablespaces) Path() path.StringPath {
	return q.path
}

func (q QInnodbSessionTempTablespaces) Size() path.Int64Path {
	return q.size
}

func (q QInnodbSessionTempTablespaces) State() path.StringPath {
	return q.state
}

func (q QInnodbSessionTempTablespaces) Purpose() path.StringPath {
	return q.purpose
}