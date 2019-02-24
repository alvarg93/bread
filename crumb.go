package bread

import (
	"io"
)

type crumb struct {
	id string
}

func newCrumb(id string) *crumb {
	return &crumb{
		id: id,
	}
}

func (c *crumb) spot(w io.Writer) {
	w.Write([]byte(c.id))
}
