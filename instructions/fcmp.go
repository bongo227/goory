package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Float comparisons
type Fcmp struct {
	name string
	mode string
	lhs  value.Value
	rhs  value.Value
}

func NewFcmp(name string, mode string, lhs value.Value, rhs value.Value) *Fcmp {
	return &Fcmp{name, mode, lhs, rhs}
}

func (i *Fcmp) String() string {
	return "fcmp"
}

func (i *Fcmp) IsTerminator() bool {
	return false
}

func (i *Fcmp) Type() types.Type {
	return types.NewBoolType()
}

func (i *Fcmp) Ident() string {
	return "%" + i.name
}

func (i *Fcmp) Llvm() string {
	return fmt.Sprintf("%%%s = fcmp %s %s %s, %s", i.name, i.mode, i.Type().String(), i.lhs.Ident(), i.rhs.Ident())
}
