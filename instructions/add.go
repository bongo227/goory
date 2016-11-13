package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

type Add struct {
	name string
	lhs  value.Value
	rhs  value.Value
}

// NewAdd creates a new Add operation
func NewAdd(name string, lhs value.Value, rhs value.Value) *Add {
	assertEqual(lhs.Type(), rhs.Type())
	assertInt(lhs.Type())
	return &Add{name, lhs, rhs}
}

func (i *Add) String() string {
	return "add"
}

func (i *Add) IsTerminator() bool {
	return false
}

func (i *Add) Type() types.Type {
	return i.lhs.Type()
}

func (i *Add) Ident() string {
	return "%" + i.name
}

func (i *Add) Llvm() string {
	return fmt.Sprintf("%%%s = add %s %s, %s",
		i.name,
		i.Type(),
		i.lhs.Ident(),
		i.rhs.Ident())
}