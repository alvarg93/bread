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

type CrumbDrop struct {
	Trail string
	Run   string
	Id    string
	Data  string
}

func Crumb(crumb CrumbDrop) {
	o.loafs.Range(func(key, l interface{}) bool {
		l.(*loaf).crumbs <- crumb
		return true
	})
}
