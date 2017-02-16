package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type selec struct {
	distinct bool
	columns  []columnOrColumnAs
}

func newSelect(columns ...columnOrColumnAs) selec {
	return selec{
		distinct: false,
		columns:  columns,
	}
}

func newSelectDistinct(columns ...columnOrColumnAs) selec {
	return selec{
		distinct: true,
		columns:  columns,
	}
}

func (s selec) node() generator.Node {
	cs := clauseNodes(s)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return generator.NewParallelNodes(fs...)
}

func (s selec) myNode() generator.Node {
	fs := make([]generator.Node, len(s.columns))
	for i, c := range s.columns {
		fs[i] = c.node()
	}
	k := generator.NewExpression(keyword.Select)
	if s.distinct {
		k = k.Append(generator.NewExpression(keyword.Distinct))
	}
	return generator.NewContainer(
		k,
		generator.NewComma(fs...),
	)
}

func (s selec) previous() clause {
	return nil
}

func (s selec) From(t joiner) from {
	f := newFrom(t)
	f.prev = s
	return f
}
