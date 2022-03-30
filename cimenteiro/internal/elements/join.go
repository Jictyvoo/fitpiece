package elements

import (
	"github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"
	"strings"
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

func (clause JoinClause) Build(writer utils.Writer) int {
	totalLength := 0

	length, _ := writer.WriteString(string(clause.JoinType))
	totalLength += length

	length, _ = writer.WriteRune(' ')
	totalLength += length

	length, _ = writer.WriteString(clause.Table.String())
	totalLength += length

	length, _ = writer.WriteString(" ON ")
	totalLength += length

	totalLength += clause.On.Build(writer)
	return totalLength
}

func (clause JoinClause) Compare(b JoinClause) int {
	if clause.Table == b.Table &&
		clause.JoinType == b.JoinType {
		builderA, builderB := strings.Builder{}, strings.Builder{}
		clause.On.Build(&builderA)
		b.On.Build(&builderB)
		if builderA.String() == builderB.String() {
			return 0
		}
		return -1
	}
	return 1
}

func (clause JoinClause) String() string {
	builder := strings.Builder{}
	clause.Build(&builder)
	return builder.String()
}
