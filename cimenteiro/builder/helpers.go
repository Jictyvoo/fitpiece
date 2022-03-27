package builder

type Writer interface {
	WriteRune(rune) (int, error)
	WriteString(string) (int, error)
}

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

func removeBrackets(str string) string {
	if len(str) < 1 {
		return str
	}
	var start, end uint = 0, uint(len(str))
	if str[0] == '(' {
		start = 1
	}
	if str[len(str)-1] == ')' {
		end = uint(len(str) - 1)
	}
	return str[start:end]
}
