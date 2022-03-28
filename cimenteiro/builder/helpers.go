package builder

type Writer interface {
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}

const (
	ClauseCreator __clauseCreator = 0
)

func buildSelectColumns(writer Writer, builder QueryBuilder) {
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

func buildJoinClauses(writer Writer, builder QueryBuilder) {
	for _, joinClause := range builder.joins {
		_, _ = writer.WriteRune(' ')
		_, _ = writer.WriteString(joinClause.Build())
	}
}
