package filters

import (
	"fmt"
	configDto "GLChecker/dto/config"
	libs "GLChecker/libs"
	getterFilter "GLChecker/filters/getter"
)

type FilterInterface interface {
	InitFilter(*configDto.Config)
	FilterFunction()
}

type Filter struct {
	FilterInterface FilterInterface
}

func InitFilter (filterType int, config *configDto.Config) *Filter{
	var cfilter FilterInterface = nil
	if(filterType == libs.GETTERFILTER){
		cfilter = new(getterFilter.Getter)
	}
	if cfilter == nil {
		panic("Filter Not Found")
	}
	cfilter.InitFilter(config)

	filter := new(Filter)
	filter.FilterInterface = cfilter

	return filter
}

func (g *Filter) Start () {
	fmt.Println("Filter started")
}

func (g *Filter) Stop() {

}