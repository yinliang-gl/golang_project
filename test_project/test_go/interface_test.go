package test_go

import (
	"errors"
	"fmt"
	"testing"
)

type AlgorithmModule interface {
	Run(a int, b int) (result int)
}

type AddModule struct {
}

func (this *AddModule) Run(a int, b int) (result int) {
	return a + b
}

type SubModule struct {
}

func (this *SubModule) Run(a int, b int) (result int) {
	return a - b
}

type MulModule struct {
}

func (this *MulModule) Run(a int, b int) (result int) {
	return a * b
}

type DivModule struct {
}

func (this *DivModule) Run(a int, b int) (result int) {
	return a / b
}

func TestInterface01(t *testing.T) {
	var algorithmModules []AlgorithmModule

	algorithmModules = append(algorithmModules, &AddModule{})
	algorithmModules = append(algorithmModules, &SubModule{})
	algorithmModules = append(algorithmModules, &MulModule{})
	algorithmModules = append(algorithmModules, &DivModule{})

	for i := 0; i < len(algorithmModules); i++ {
		fmt.Println(algorithmModules[i].Run(100, 90))
	}

}

func GetModule(module_name string) (Algor AlgorithmModule, err error) {
	if module_name == "add" {
		return &AddModule{}, nil
	} else if module_name == "sub" {
		return &SubModule{}, nil
	} else {
		return nil, errors.New("unsupport module:" + module_name)
	}
}

func TestInterface02(t *testing.T) {
	var moduleNames []string
	moduleNames = append(moduleNames, "add")
	moduleNames = append(moduleNames, "sub")
	moduleNames = append(moduleNames, "aaa")

	for i := 0; i < len(moduleNames); i++ {
		module, err := GetModule(moduleNames[i])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("%s, %d\n", moduleNames[i], module.Run(100, 90))
		// fmt.Println(GetModule(moduleNames[i]).Run(100, 90))
	}
}
