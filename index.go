package pet

import "github.com/gopherjs/gopherjs/js"

type Pet struct {
	name string
}

func (p *Pet) Name() string {
	return p.name
}

func (p *Pet) SetName(newName string) {
	p.name = newName
}

func New(name string) *js.Object {
	return js.MakeWrapper(&Pet{name})
}
