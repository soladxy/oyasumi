package main

const interfaceTmpl = `package {{.LowerName}}

import (
	"context"
	"github.com/soladxy/oyasumi/biz/container"
)

type I{{.FirstCharUpperName}} interface {
	Sample(ctx context.Context) error
}

func New{{.FirstCharUpperName}}Service(c *container.Container) I{{.FirstCharUpperName}} {
	return new{{.FirstCharUpperName}}(c)
}
`
const implTmpl = `package {{.LowerName}}

import (
	"context"
	"errors"
	"github.com/soladxy/oyasumi/biz/container"
)

type _{{.LowerName}} struct {
	c *container.Container
}

func new{{.FirstCharUpperName}}(c *container.Container) *_{{.LowerName}} {
	return &_{{.LowerName}}{
		c: c,
	}
}

func (s *_{{.LowerName}}) Sample(ctx context.Context) error {
	return errors.New("not implemented")
}
`
