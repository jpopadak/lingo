// Code generated by Lingo for table information_schema.TABLES - DO NOT EDIT

package qtables

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QTables {
	return newQTables(alias)
}

func New() QTables {
	return newQTables("")
}

func newQTables(alias string) QTables {
	q := QTables{_alias: alias}
	q.tableCatalog = path.NewStringPath(q, "TABLE_CATALOG")
	q.tableSchema = path.NewStringPath(q, "TABLE_SCHEMA")
	q.tableName = path.NewStringPath(q, "TABLE_NAME")
	q.tableType = path.NewStringPath(q, "TABLE_TYPE")
	q.engine = path.NewStringPath(q, "ENGINE")
	q.version = path.NewInt64Path(q, "VERSION")
	q.rowFormat = path.NewStringPath(q, "ROW_FORMAT")
	q.tableRows = path.NewInt64Path(q, "TABLE_ROWS")
	q.avgRowLength = path.NewInt64Path(q, "AVG_ROW_LENGTH")
	q.dataLength = path.NewInt64Path(q, "DATA_LENGTH")
	q.maxDataLength = path.NewInt64Path(q, "MAX_DATA_LENGTH")
	q.indexLength = path.NewInt64Path(q, "INDEX_LENGTH")
	q.dataFree = path.NewInt64Path(q, "DATA_FREE")
	q.autoIncrement = path.NewInt64Path(q, "AUTO_INCREMENT")
	q.createTime = path.NewTimePath(q, "CREATE_TIME")
	q.updateTime = path.NewTimePath(q, "UPDATE_TIME")
	q.checkTime = path.NewTimePath(q, "CHECK_TIME")
	q.tableCollation = path.NewStringPath(q, "TABLE_COLLATION")
	q.checksum = path.NewInt64Path(q, "CHECKSUM")
	q.createOptions = path.NewStringPath(q, "CREATE_OPTIONS")
	q.tableComment = path.NewStringPath(q, "TABLE_COMMENT")
	return q
}

type QTables struct {
	_alias         string
	tableCatalog   path.StringPath
	tableSchema    path.StringPath
	tableName      path.StringPath
	tableType      path.StringPath
	engine         path.StringPath
	version        path.Int64Path
	rowFormat      path.StringPath
	tableRows      path.Int64Path
	avgRowLength   path.Int64Path
	dataLength     path.Int64Path
	maxDataLength  path.Int64Path
	indexLength    path.Int64Path
	dataFree       path.Int64Path
	autoIncrement  path.Int64Path
	createTime     path.TimePath
	updateTime     path.TimePath
	checkTime      path.TimePath
	tableCollation path.StringPath
	checksum       path.Int64Path
	createOptions  path.StringPath
	tableComment   path.StringPath
}

// core.Table Functions

func (q QTables) GetColumns() []core.Column {
	return []core.Column{
		q.tableCatalog,
		q.tableSchema,
		q.tableName,
		q.tableType,
		q.engine,
		q.version,
		q.rowFormat,
		q.tableRows,
		q.avgRowLength,
		q.dataLength,
		q.maxDataLength,
		q.indexLength,
		q.dataFree,
		q.autoIncrement,
		q.createTime,
		q.updateTime,
		q.checkTime,
		q.tableCollation,
		q.checksum,
		q.createOptions,
		q.tableComment,
	}
}

func (q QTables) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QTables) GetAlias() string {
	return q._alias
}

func (q QTables) GetName() string {
	return "TABLES"
}

func (q QTables) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QTables) TableCatalog() path.StringPath {
	return q.tableCatalog
}

func (q QTables) TableSchema() path.StringPath {
	return q.tableSchema
}

func (q QTables) TableName() path.StringPath {
	return q.tableName
}

func (q QTables) TableType() path.StringPath {
	return q.tableType
}

func (q QTables) Engine() path.StringPath {
	return q.engine
}

func (q QTables) Version() path.Int64Path {
	return q.version
}

func (q QTables) RowFormat() path.StringPath {
	return q.rowFormat
}

func (q QTables) TableRows() path.Int64Path {
	return q.tableRows
}

func (q QTables) AvgRowLength() path.Int64Path {
	return q.avgRowLength
}

func (q QTables) DataLength() path.Int64Path {
	return q.dataLength
}

func (q QTables) MaxDataLength() path.Int64Path {
	return q.maxDataLength
}

func (q QTables) IndexLength() path.Int64Path {
	return q.indexLength
}

func (q QTables) DataFree() path.Int64Path {
	return q.dataFree
}

func (q QTables) AutoIncrement() path.Int64Path {
	return q.autoIncrement
}

func (q QTables) CreateTime() path.TimePath {
	return q.createTime
}

func (q QTables) UpdateTime() path.TimePath {
	return q.updateTime
}

func (q QTables) CheckTime() path.TimePath {
	return q.checkTime
}

func (q QTables) TableCollation() path.StringPath {
	return q.tableCollation
}

func (q QTables) Checksum() path.Int64Path {
	return q.checksum
}

func (q QTables) CreateOptions() path.StringPath {
	return q.createOptions
}

func (q QTables) TableComment() path.StringPath {
	return q.tableComment
}
