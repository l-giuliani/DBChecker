package services

import (
	context "GLChecker/context"
	libs "GLChecker/libs"
	getterFilter "GLChecker/filters/getter"
)

func InitSystem (){
	context.InitContext()

	config := libs.GetConfig()

	getters := config.Getters
	for _, getterConf := range getters {
		getter := new(getterFilter.Getter)
		getter.InitGetter(getterConf)
		context.AddGetterFilter(getter)
		getter.Start()
	}
}