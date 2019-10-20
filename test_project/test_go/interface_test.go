package test_go

import (
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

func TestAAA(t *testing.T) {
	var asDspModules []AlgorithmModule

	asDspModules = append(asDspModules, &AddModule{})
	asDspModules = append(asDspModules, &SubModule{})
	asDspModules = append(asDspModules, &MulModule{})
	asDspModules = append(asDspModules, &DivModule{})

	for i := 0; i < len(asDspModules); i++ {
		fmt.Println(asDspModules[i].Run(100, 90))
	}

}
