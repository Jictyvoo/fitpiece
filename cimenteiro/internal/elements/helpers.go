package elements

type Expression interface {
	Build() string
	BuildPlaceholder(placeholder string) (string, []any)
}
