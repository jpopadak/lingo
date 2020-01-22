package core


// ========== NAME ==========
type Name interface {
	Name() string
}
// ========== ALIAS ==========
type Alias interface {
	Name() string
	Alias() string
}

// ========== DIALECT ==========
type Dialect interface {
	Select(n ...Name) SelectClause
}

// ==================================================
type SelectClause interface {
	FromTable(n Name) JoinClause
}

// ==================================================
type JoinClause interface {
	LeftJoinOn() JoinClause
	RightJoinOn() JoinClause
	Where() WhereClause
}
type WhereClause interface {
	And()
	Or()
}

// ==================================================

func doThings() {
	var d Dialect = mysql{}
	var tRG = TResourceGroup{}
	var tRR = TResourceReference{}
	d.Select(tRG.).FromTable(t).LeftJoinOn().Where()
}