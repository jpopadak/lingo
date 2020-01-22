// Code generated by Lingo for table information_schema.COLUMN_STATISTICS - DO NOT EDIT

package qcolumnstatistics

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/path"
)

func As(alias string) QColumnStatistics {
	return newQColumnStatistics(alias)
}

func New() QColumnStatistics {
	return newQColumnStatistics("")
}

func newQColumnStatistics(alias string) QColumnStatistics {
	q := QColumnStatistics{_alias: alias}
	q.schemaName = path.NewStringPath(q, "SCHEMA_NAME")
	q.tableName = path.NewStringPath(q, "TABLE_NAME")
	q.columnName = path.NewStringPath(q, "COLUMN_NAME")
	q.histogram = path.NewJSONPath(q, "HISTOGRAM")
	return q
}

type QColumnStatistics struct {
	_alias     string
	schemaName path.StringPath
	tableName  path.StringPath
	columnName path.StringPath
	histogram  path.JSONPath
}

// core.Table Functions

func (q QColumnStatistics) GetColumns() []core.Column {
	return []core.Column{
		q.schemaName,
		q.tableName,
		q.columnName,
		q.histogram,
	}
}

func (q QColumnStatistics) GetSQL(d core.Dialect, sql core.SQL) error {
	return path.ExpandTableWithDialect(d, q, sql)
}

func (q QColumnStatistics) GetAlias() string {
	return q._alias
}

func (q QColumnStatistics) GetName() string {
	return "COLUMN_STATISTICS"
}

func (q QColumnStatistics) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QColumnStatistics) SchemaName() path.StringPath {
	return q.schemaName
}

func (q QColumnStatistics) TableName() path.StringPath {
	return q.tableName
}

func (q QColumnStatistics) ColumnName() path.StringPath {
	return q.columnName
}

func (q QColumnStatistics) Histogram() path.JSONPath {
	return q.histogram
}