// Code generated by Lingo for table information_schema.INNODB_METRICS - DO NOT EDIT

package qinnodbmetrics

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QInnodbMetrics {
	return newQInnodbMetrics(alias)
}

func New() QInnodbMetrics {
	return newQInnodbMetrics("")
}

func newQInnodbMetrics(alias string) QInnodbMetrics {
	q := QInnodbMetrics{_alias: alias}
	q.name = path.NewStringPath(q, "NAME")
	q.subsystem = path.NewStringPath(q, "SUBSYSTEM")
	q.count = path.NewInt64Path(q, "COUNT")
	q.maxCount = path.NewInt64Path(q, "MAX_COUNT")
	q.minCount = path.NewInt64Path(q, "MIN_COUNT")
	q.avgCount = path.NewFloat64Path(q, "AVG_COUNT")
	q.countReset = path.NewInt64Path(q, "COUNT_RESET")
	q.maxCountReset = path.NewInt64Path(q, "MAX_COUNT_RESET")
	q.minCountReset = path.NewInt64Path(q, "MIN_COUNT_RESET")
	q.avgCountReset = path.NewFloat64Path(q, "AVG_COUNT_RESET")
	q.timeEnabled = path.NewTimePath(q, "TIME_ENABLED")
	q.timeDisabled = path.NewTimePath(q, "TIME_DISABLED")
	q.timeElapsed = path.NewInt64Path(q, "TIME_ELAPSED")
	q.timeReset = path.NewTimePath(q, "TIME_RESET")
	q.status = path.NewStringPath(q, "STATUS")
	q.__type = path.NewStringPath(q, "TYPE")
	q.comment = path.NewStringPath(q, "COMMENT")
	return q
}

type QInnodbMetrics struct {
	_alias        string
	name          path.StringPath
	subsystem     path.StringPath
	count         path.Int64Path
	maxCount      path.Int64Path
	minCount      path.Int64Path
	avgCount      path.Float64Path
	countReset    path.Int64Path
	maxCountReset path.Int64Path
	minCountReset path.Int64Path
	avgCountReset path.Float64Path
	timeEnabled   path.TimePath
	timeDisabled  path.TimePath
	timeElapsed   path.Int64Path
	timeReset     path.TimePath
	status        path.StringPath
	__type        path.StringPath
	comment       path.StringPath
}

// core.Table Functions

func (q QInnodbMetrics) GetColumns() []core.Column {
	return []core.Column{
		q.name,
		q.subsystem,
		q.count,
		q.maxCount,
		q.minCount,
		q.avgCount,
		q.countReset,
		q.maxCountReset,
		q.minCountReset,
		q.avgCountReset,
		q.timeEnabled,
		q.timeDisabled,
		q.timeElapsed,
		q.timeReset,
		q.status,
		q.__type,
		q.comment,
	}
}

func (q QInnodbMetrics) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QInnodbMetrics) GetAlias() string {
	return q._alias
}

func (q QInnodbMetrics) GetName() string {
	return "INNODB_METRICS"
}

func (q QInnodbMetrics) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QInnodbMetrics) Name() path.StringPath {
	return q.name
}

func (q QInnodbMetrics) Subsystem() path.StringPath {
	return q.subsystem
}

func (q QInnodbMetrics) Count() path.Int64Path {
	return q.count
}

func (q QInnodbMetrics) MaxCount() path.Int64Path {
	return q.maxCount
}

func (q QInnodbMetrics) MinCount() path.Int64Path {
	return q.minCount
}

func (q QInnodbMetrics) AvgCount() path.Float64Path {
	return q.avgCount
}

func (q QInnodbMetrics) CountReset() path.Int64Path {
	return q.countReset
}

func (q QInnodbMetrics) MaxCountReset() path.Int64Path {
	return q.maxCountReset
}

func (q QInnodbMetrics) MinCountReset() path.Int64Path {
	return q.minCountReset
}

func (q QInnodbMetrics) AvgCountReset() path.Float64Path {
	return q.avgCountReset
}

func (q QInnodbMetrics) TimeEnabled() path.TimePath {
	return q.timeEnabled
}

func (q QInnodbMetrics) TimeDisabled() path.TimePath {
	return q.timeDisabled
}

func (q QInnodbMetrics) TimeElapsed() path.Int64Path {
	return q.timeElapsed
}

func (q QInnodbMetrics) TimeReset() path.TimePath {
	return q.timeReset
}

func (q QInnodbMetrics) Status() path.StringPath {
	return q.status
}

func (q QInnodbMetrics) Type() path.StringPath {
	return q.__type
}

func (q QInnodbMetrics) Comment() path.StringPath {
	return q.comment
}