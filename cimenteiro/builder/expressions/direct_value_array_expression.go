package expressions

type ValueFieldType interface {
	~string | ~int | ~uint
}

type DirectValueArrayExpression[T ValueFieldType] struct {
	values []T
	wrapIn [2]rune
}

func (expression DirectValueArrayExpression[T]) Build() string {
	return ArrayElementExpression[T]{
		values: expression.values,
		wrapIn: expression.wrapIn,
	}.Build()
}

func (expression DirectValueArrayExpression[T]) BuildPlaceholder(placeholder string) (string, []any) {
	return expression.Build(), []any{}
}
