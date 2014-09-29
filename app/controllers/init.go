package controllers

import (
	"github.com/revel/revel"
)

var Container *Dic = new(Dic)

func init() {
	revel.OnAppStart(func () { Container.Init() })
}
