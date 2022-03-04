package services

import (
	context "GLChecker/context"
	libs "GLChecker/libs"
	//getterFilter "GLChecker/filters/getter"
	filters "GLChecker/filters"
)

func InitSystem (){
	context.InitContext()

	config := libs.GetConfig()

	getters := config.Getters
	for range getters {
		//getter := new(getterFilter.Getter)
		//getter.InitGetter(getterConf)
		filter := filters.InitFilter(libs.GETTERFILTER, config)
		context.AddFilter(filter)
		filter.Start()
	}
}