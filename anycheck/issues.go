package anycheck

import "go/token"

type Issue interface {
	Pos() token.Pos
	Message() string
}

type interfaceNotAllowed struct {
	pos token.Pos
}

func (i interfaceNotAllowed) Pos() token.Pos {
	return i.pos
}

func (i interfaceNotAllowed) Message() string {
	return "Use of 'interface{}' detected. Replace with 'any'."
}

type anyNotAllowed struct {
	pos token.Pos
}

func (i anyNotAllowed) Pos() token.Pos {
	return i.pos
}

func (i anyNotAllowed) Message() string {
	return "Use of 'any' detected. Replace with 'interface{}'."
}
