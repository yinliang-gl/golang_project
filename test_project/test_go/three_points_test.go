package test_go

import (
	"fmt"
	"strconv"
	"testing"
)

func ThreePointsTestFunction(val int, lines ...string) string {
	tmp := strconv.Itoa(val)

	for _, line := range lines {
		tmp += "_"
		tmp += line
	}
	return tmp
}

func TestThreePoints(t *testing.T) {
	fmt.Println(ThreePointsTestFunction(10, "AA", "AB"))
	fmt.Println(ThreePointsTestFunction(10, "AA", "AB", "AC"))
	fmt.Println(ThreePointsTestFunction(10, "AA", "AB", "AC", "AD"))
}
