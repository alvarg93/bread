package bread

import (
	"sync"
)

type oven struct {
	loafs sync.Map
}

var o *oven
var once sync.Once

func Bake(r recipe) {
	once.Do(func() {
		o = &oven{}
	})
	l := newLoaf(r.Writer)
	go l.baking()
	o.loafs.Store(r.Tag, l)
}

func Crumb(trail, data string) {
	o.loafs.Range(func(key, l interface{}) bool {
		l.(*loaf).crumbs <- crumbMsg{trail, data}
		return true
	})
}
