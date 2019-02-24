package bread

import (
	"io"
)

type trail struct {
	crumbs map[string]*crumb
	order  []*crumb
}

func newPath() *trail {
	return &trail{
		crumbs: make(map[string]*crumb),
		order:  make([]*crumb, 0),
	}
}

func (t *trail) crumb(writer io.Writer, id string) {
	c, ok := t.crumbs[id]
	if !ok {
		c = newCrumb(id)
		t.order = append(t.order)
	}
	c.spot(writer)
}
