package bread

import (
	"io"
)

type crumb struct {
	id     string
	num    int
	data   map[string]string
	writer io.Writer
}

func newCrumb(id string, num int, writer io.Writer) *crumb {
	return &crumb{
		id:     id,
		num:    num,
		data:   make(map[string]string),
		writer: writer,
	}
}

func (c *crumb) drop(drop CrumbDrop) {
	c.data[drop.Run] = drop.Data
}

func (c *crumb) spot(run string) {
	var idBytes [32]byte
	copy(idBytes[:], c.id)
	_, err := c.writer.Write(append(idBytes[:], []byte(c.data[run])...))
	if err != nil {
		panic(err)
	}
}
