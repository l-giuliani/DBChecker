package filters

import (
	"fmt"
	configDto "GLChecker/dto/config"
)

type Getter struct {

}

func (g *Getter) getterFunction() {

}

func (g *Getter) InitGetter (getter configDto.Getter) {
	fmt.Println(getter.Source)
}

func (g *Getter) Start () {
	fmt.Println("Getter started")
}

func (g *Getter) Stop() {

}