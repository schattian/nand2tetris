package symbol

type Type int

const (
	TYPE_INT Type = iota
	TYPE_CHAR
	TYPE_BOOLEAN
	TYPE_CLASS_NAME
)

type Kind int

const (
	KIND_FIELD Kind = iota
	KIND_STATIC
	KIND_LOCAL
	KIND_ARG
)

type Symbol struct {
	Type  Type
	Kind  Kind
	Index int
}
