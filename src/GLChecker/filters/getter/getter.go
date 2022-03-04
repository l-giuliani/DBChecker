package filters

import (
	"fmt"
	configDto "GLChecker/dto/config"
)

type Getter struct {
}

func (g *Getter) FilterFunction() {

}

func (g *Getter) InitFilter (config *configDto.Config) {
	fmt.Println("Initializing Getter Filter")
}