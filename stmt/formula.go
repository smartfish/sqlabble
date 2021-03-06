package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Formula struct {
	arithmetic keyword.Operator
	numeric    []ValOrColOrSubOrFormula
}

func (f Formula) nodeize() (tokenizer.Tokenizer, []interface{}) {
	ts := make(tokenizer.Tokenizers, len(f.numeric))
	values := []interface{}{}
	for i, num := range f.numeric {
		t, vals := num.nodeize()
		if _, ok := num.(Formula); ok {
			t = tokenizer.NewParentheses(t)
		}
		ts[i] = t
		values = append(values, vals...)
	}
	return tokenizer.NewTokenizers(ts...).Prefix(
		token.Word(f.arithmetic),
	), values
}

func (f Formula) keyword() keyword.Operator {
	return f.arithmetic
}

func NewAdd(numeric ...ValOrColOrSubOrFormula) Formula {
	return Formula{
		arithmetic: keyword.Add,
		numeric:    numeric,
	}
}

func NewSub(numeric ...ValOrColOrSubOrFormula) Formula {
	return Formula{
		arithmetic: keyword.Sub,
		numeric:    numeric,
	}
}

func NewMul(numeric ...ValOrColOrSubOrFormula) Formula {
	return Formula{
		arithmetic: keyword.Mul,
		numeric:    numeric,
	}
}

func NewDiv(numeric ...ValOrColOrSubOrFormula) Formula {
	return Formula{
		arithmetic: keyword.Div,
		numeric:    numeric,
	}
}

func NewIntegerDiv(numeric ...ValOrColOrSubOrFormula) Formula {
	return Formula{
		arithmetic: keyword.IntegerDiv,
		numeric:    numeric,
	}
}

func NewMod(numeric ...ValOrColOrSubOrFormula) Formula {
	return Formula{
		arithmetic: keyword.Mod,
		numeric:    numeric,
	}
}

// isValOrColOrAliasOrSubOrForm always returns true.
// This method exists only to implement the interface ValOrColOrAliasOrSubOrForm.
// This is a shit of duck typing, but anyway it works.
func (f Formula) isValOrColOrSubOrFormula() bool {
	return true
}

// isColOrAliasOrFuncOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Formula) isValOrColOrAliasOrFuncOrSubOrFormula() bool {
	return true
}
