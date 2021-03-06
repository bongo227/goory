package instructions

import (
	"fmt"

	"github.com/bongo227/goory/types"

	"github.com/bongo227/goory/value"
)

// Conditional Branch statement
type CondBr struct {
	block      value.Value
	name       string
	condition  value.Value
	trueBlock  value.Value
	falseBlock value.Value
}

func NewCondBr(block value.Value, name string, condition value.Value, trueBlock value.Value, falseBlock value.Value) *CondBr {
	if !condition.Type().Equal(types.NewBoolType()) {
		panic("condition is not a bool type")
	}

	if !trueBlock.Type().Equal(types.NewBlockType()) {
		panic("trueBlock is not a block type")
	}

	if !falseBlock.Type().Equal(types.NewBlockType()) {
		panic("falseBlock is not a block type")
	}

	return &CondBr{block, name, condition, trueBlock, falseBlock}
}

func (i *CondBr) Block() value.Value {
	return i.block
}

func (i *CondBr) IsTerminator() bool {
	return true
}

func (i *CondBr) Type() types.Type {
	return types.VOID
}

func (i *CondBr) Ident() string {
	return "%" + i.name
}

func (i *CondBr) Llvm() string {
	return fmt.Sprintf("br i1 %s, label %%%s, label %%%s",
		i.condition.Ident(),
		i.trueBlock.Ident(),
		i.falseBlock.Ident())
}
