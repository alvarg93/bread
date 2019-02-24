package bread

import (
	"io"
)

type recipe struct {
	Tag    string
	Writer io.Writer
}

func Recipe(tag string, writer io.Writer) *recipe {
	return &recipe{
		Tag:    tag,
		Writer: writer,
	}
}
