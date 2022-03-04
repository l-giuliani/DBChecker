package context

import (
	filters "GLChecker/filters"
)

type Context struct {
	Filters []*filters.Filter
}

var context Context
var initialized bool = false 

func InitContext () {
	context.Filters = make([]*filters.Filter, 0)

	initialized = true
}

func AddFilter (filter *filters.Filter) {
	if !initialized {
		panic("Context not initialized")
	}
	context.Filters = append(context.Filters, filter)
}