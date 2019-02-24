package bread

import (
	"io"
)

type loaf struct {
	trails map[string]*trail
	writer io.Writer
	crumbs chan CrumbDrop
}

func newLoaf(w io.Writer) *loaf {
	return &loaf{
		trails: make(map[string]*trail),
		writer: w,
		crumbs: make(chan CrumbDrop),
	}
}

func (l *loaf) baking() {
	for {
		select {
		case msg := <-l.crumbs:
			trail := l.trails[msg.Trail]
			if trail == nil {
				trail = newTrail()
				l.trails[msg.Trail] = trail
			}
			trail.crumb(l.writer, msg)
		default:
		}
	}
}
