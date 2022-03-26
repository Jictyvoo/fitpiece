package elements

import (
	"fmt"
)

type joinType string

const (
	JoinAll   joinType = "JOIN"
	JoinInner joinType = "INNER JOIN"
	JoinOuter joinType = "OUTER JOIN"
	JoinLeft  joinType = "LEFT JOIN"
	JoinRight joinType = "RIGHT JOIN"
)

type JoinClause struct {
	JoinType joinType
	Table    TableName
	On       Clause
}

func (clause JoinClause) Build() string {
	return fmt.Sprintf(
		"%s %s ON %s",
		clause.JoinType, clause.Table.String(), clause.On.Build(),
	)
}

func (clause JoinClause) Compare(b JoinClause) int {
	if clause.Table == b.Table &&
		clause.JoinType == b.JoinType &&
		clause.On.Build() == b.On.Build() {
		return 0
	}
	return 1
}
