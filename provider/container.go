package provider

import "go.uber.org/dig"

type ProviderRegisterFuns []interface{}

type ProviderInterface interface {
	Register() ProviderRegisterFuns
}

var container *dig.Container

func SetContainer(c *dig.Container) {
	container = c
}

func check() {

	if container == nil {
		panic("container对象未初始化")
	}
}

func Cont() *dig.Container {
	return container
}
