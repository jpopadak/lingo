package core

// ========== TABLE ==========
type TResourceGroup struct {}
func (t TResourceGroup) Name() string { return "resource_group" }
func (t TResourceGroup) As(a string) Alias { return alias{name: t.Name(), alias: a}}
func (t TResourceGroup) UUID() BinaryPath {}


type TResourceReference struct {}
func (t TResourceReference) Name() string { return "resource_reference" }
func (t TResourceReference) As(a string) Alias { return alias{name: t.Name(), alias: a}}

// ========== BINARY PATH ==========
type BinaryPath interface {
	Name() string
	As(alias string) Alias
	Value() []byte
}

// ========== TABLE ALIAS ==========
type alias struct {
	name string
	alias string
}
func (ta alias) Name() string { return ta.name }
func (ta alias) Alias() string { return ta.alias }

// ========== DIALECT ==========
type mysql struct{}
func (m mysql) Select(name ...Name) SelectClause { return selectClause{} }

// ========== SELECT CLAUSE ==========
type selectClause struct{}
func (s selectClause) FromTable(n Name) WhereClause { return whereClause{} }

// ========== JOIN CLAUSE ==========
type joinClause struct {}
func (s joinClause) LeftJoin() JoinClause { }
func (s joinClause) RightJoin() JoinClause
func (s joinClause) Where() WhereClause

// ========== JOIN CLAUSE ==========

// ========== WHERE CLAUSE ==========
type whereClause struct {}
func (w whereClause) Where() {}
