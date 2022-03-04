package context

import (
	getterFilter "GLChecker/filters/getter"
)

type Context struct {
	GetterFilter []*getterFilter.Getter
}

var context Context
var initialized bool = false 

func InitContext () {
	context.GetterFilter = make([]*getterFilter.Getter, 0)

	initialized = true
}

func AddGetterFilter (getter *getterFilter.Getter) {
	if !initialized {
		panic("Context not initialized")
	}
	context.GetterFilter = append(context.GetterFilter, getter)
}