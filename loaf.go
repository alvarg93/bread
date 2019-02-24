package bread

import (
	"io"
)

type loaf struct {
	trails map[string]*trail
	writer io.Writer
	crumbs chan crumbMsg
}

type crumbMsg struct {
	trail string
	data  string
}

func newLoaf(w io.Writer) *loaf {
	return &loaf{
		trails: make(map[string]*trail),
		writer: w,
		crumbs: make(chan crumbMsg),
	}
}

func (l *loaf) baking() {
	for {
		select {
		case msg := <-l.crumbs:
			trail := l.trails[msg.trail]
			if trail == nil {
				trail = newPath()
				l.trails[msg.trail] = trail
			}
			trail.crumb(l.writer, msg.data)
		default:
		}
	}
}
