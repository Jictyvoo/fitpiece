package builder

import "github.com/wrapped-owls/fitpiece/cimenteiro/internal/utils"

const (
	ClauseCreator __clauseCreator = false
)

func buildSelectColumns(writer utils.Writer, builder QueryBuilder) {
	if len(builder.fields) < 1 {
		_, _ = writer.WriteRune('*')
	} else {
		for index, fieldName := range builder.fields {
			if index > 0 {
				_, _ = writer.WriteRune(',')
				_, _ = writer.WriteRune(' ')
			}
			_, _ = writer.WriteString(fieldName.Build())
		}
	}
}

func buildJoinClauses(writer utils.Writer, builder QueryBuilder) {
	for _, joinClause := range builder.joins {
		_, _ = writer.WriteRune(' ')
		_, _ = writer.WriteString(joinClause.Build())
	}
}
