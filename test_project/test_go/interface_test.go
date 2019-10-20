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
	var algorithmModules []AlgorithmModule

	algorithmModules = append(algorithmModules, &AddModule{})
	algorithmModules = append(algorithmModules, &SubModule{})
	algorithmModules = append(algorithmModules, &MulModule{})
	algorithmModules = append(algorithmModules, &DivModule{})

	for i := 0; i < len(algorithmModules); i++ {
		fmt.Println(algorithmModules[i].Run(100, 90))
	}

}
