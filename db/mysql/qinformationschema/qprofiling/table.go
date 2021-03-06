// Code generated by Lingo for table information_schema.PROFILING - DO NOT EDIT

package qprofiling

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QProfiling {
	return newQProfiling(alias)
}

func New() QProfiling {
	return newQProfiling("")
}

func newQProfiling(alias string) QProfiling {
	q := QProfiling{_alias: alias}
	q.queryId = path.NewIntPath(q, "QUERY_ID")
	q.seq = path.NewIntPath(q, "SEQ")
	q.state = path.NewStringPath(q, "STATE")
	q.duration = path.NewBinaryPath(q, "DURATION")
	q.cpuUser = path.NewBinaryPath(q, "CPU_USER")
	q.cpuSystem = path.NewBinaryPath(q, "CPU_SYSTEM")
	q.contextVoluntary = path.NewIntPath(q, "CONTEXT_VOLUNTARY")
	q.contextInvoluntary = path.NewIntPath(q, "CONTEXT_INVOLUNTARY")
	q.blockOpsIn = path.NewIntPath(q, "BLOCK_OPS_IN")
	q.blockOpsOut = path.NewIntPath(q, "BLOCK_OPS_OUT")
	q.messagesSent = path.NewIntPath(q, "MESSAGES_SENT")
	q.messagesReceived = path.NewIntPath(q, "MESSAGES_RECEIVED")
	q.pageFaultsMajor = path.NewIntPath(q, "PAGE_FAULTS_MAJOR")
	q.pageFaultsMinor = path.NewIntPath(q, "PAGE_FAULTS_MINOR")
	q.swaps = path.NewIntPath(q, "SWAPS")
	q.sourceFunction = path.NewStringPath(q, "SOURCE_FUNCTION")
	q.sourceFile = path.NewStringPath(q, "SOURCE_FILE")
	q.sourceLine = path.NewIntPath(q, "SOURCE_LINE")
	return q
}

type QProfiling struct {
	_alias             string
	queryId            path.IntPath
	seq                path.IntPath
	state              path.StringPath
	duration           path.BinaryPath
	cpuUser            path.BinaryPath
	cpuSystem          path.BinaryPath
	contextVoluntary   path.IntPath
	contextInvoluntary path.IntPath
	blockOpsIn         path.IntPath
	blockOpsOut        path.IntPath
	messagesSent       path.IntPath
	messagesReceived   path.IntPath
	pageFaultsMajor    path.IntPath
	pageFaultsMinor    path.IntPath
	swaps              path.IntPath
	sourceFunction     path.StringPath
	sourceFile         path.StringPath
	sourceLine         path.IntPath
}

// core.Table Functions

func (q QProfiling) GetColumns() []core.Column {
	return []core.Column{
		q.queryId,
		q.seq,
		q.state,
		q.duration,
		q.cpuUser,
		q.cpuSystem,
		q.contextVoluntary,
		q.contextInvoluntary,
		q.blockOpsIn,
		q.blockOpsOut,
		q.messagesSent,
		q.messagesReceived,
		q.pageFaultsMajor,
		q.pageFaultsMinor,
		q.swaps,
		q.sourceFunction,
		q.sourceFile,
		q.sourceLine,
	}
}

func (q QProfiling) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QProfiling) GetAlias() string {
	return q._alias
}

func (q QProfiling) GetName() string {
	return "PROFILING"
}

func (q QProfiling) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QProfiling) QueryId() path.IntPath {
	return q.queryId
}

func (q QProfiling) Seq() path.IntPath {
	return q.seq
}

func (q QProfiling) State() path.StringPath {
	return q.state
}

func (q QProfiling) Duration() path.BinaryPath {
	return q.duration
}

func (q QProfiling) CpuUser() path.BinaryPath {
	return q.cpuUser
}

func (q QProfiling) CpuSystem() path.BinaryPath {
	return q.cpuSystem
}

func (q QProfiling) ContextVoluntary() path.IntPath {
	return q.contextVoluntary
}

func (q QProfiling) ContextInvoluntary() path.IntPath {
	return q.contextInvoluntary
}

func (q QProfiling) BlockOpsIn() path.IntPath {
	return q.blockOpsIn
}

func (q QProfiling) BlockOpsOut() path.IntPath {
	return q.blockOpsOut
}

func (q QProfiling) MessagesSent() path.IntPath {
	return q.messagesSent
}

func (q QProfiling) MessagesReceived() path.IntPath {
	return q.messagesReceived
}

func (q QProfiling) PageFaultsMajor() path.IntPath {
	return q.pageFaultsMajor
}

func (q QProfiling) PageFaultsMinor() path.IntPath {
	return q.pageFaultsMinor
}

func (q QProfiling) Swaps() path.IntPath {
	return q.swaps
}

func (q QProfiling) SourceFunction() path.StringPath {
	return q.sourceFunction
}

func (q QProfiling) SourceFile() path.StringPath {
	return q.sourceFile
}

func (q QProfiling) SourceLine() path.IntPath {
	return q.sourceLine
}
