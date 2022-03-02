package libs

import (
	"os"
	"encoding/json"
	configDto "GLChecker/dto/config"
)

var config configDto.Config
var initialized bool = false 

func initConfig(){
	confName := "config.json"

	conf, err := os.ReadFile(confName)
	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(conf), &config)
    initialized = true
}

func GetConfig() *configDto.Config{
	if !initialized {
		initConfig()
	}
	return &config
}