package bread

import (
	"io"
	"log"
)

type trail struct {
	crumbs map[string]*crumb
	order  []*crumb
	counts []int
}

func newTrail() *trail {
	return &trail{
		crumbs: make(map[string]*crumb),
		order:  make([]*crumb, 0),
		counts: make([]int, 0),
	}
}

func (t *trail) crumb(writer io.Writer, drop CrumbDrop) {
	c, ok := t.crumbs[drop.Id]
	if !ok {
		c = t.addCrumb(writer, drop.Id)
	}
	t.counts[c.num]++
	c.drop(drop)
	if c.num > 0 && t.counts[c.num-1] < t.counts[c.num] {
		for _, c := range t.order {
			c.spot(drop.Run)
		}
	}
	log.Println(t.counts)
}

func (t *trail) addCrumb(writer io.Writer, id string) *crumb {
	c := newCrumb(id, len(t.order), writer)
	t.crumbs[id] = c
	t.order = append(t.order, c)
	t.counts = append(t.counts, 0)
	return c
}
