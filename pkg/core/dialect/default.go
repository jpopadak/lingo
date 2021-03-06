package dialect

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

type Default struct{}

func (Default) GetName() string {
	return "Default"
}

func (Default) ValueFormat() string {
	return "?"
}

func (Default) SetValueFormat() string {
	return "="
}

func (Default) ExpandTable(table core.Table) (core.SQL, error) {
	return ExpandEntity(table)
}

func (Default) ExpandColumn(column core.Column) (core.SQL, error) {
	return ExpandColumn(column)
}

func (Default) Operator(left core.SQL, op operator.Operand, values []core.SQL) (core.SQL, error) {
	switch op {
	// No special operations needed beyond ANSI SQL
	}

	return Operator(left, op, values)
}

func (m Default) Value(value interface{}) (core.SQL, error) {
	return Value(m, value)
}

func (m Default) Join(left core.SQL, joinType expression.JoinType, on core.SQL) (core.SQL, error) {
	return Join(left, genericJoinTypeToStr[joinType], on)
}

func (m Default) OrderBy(left core.SQL, direction sort.Direction) (core.SQL, error) {
	return OrderBy(left, direction)
}

func (m Default) Set(left core.SQL, value core.SQL) (core.SQL, error) {
	return Set(m, left, value)
}
