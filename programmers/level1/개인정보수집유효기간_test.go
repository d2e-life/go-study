package level1

import (
	"go-study/programmers"
	"testing"
)

func TestSolution개인정보수집유효기간(t *testing.T) {

	expect1 := []int{1, 3}
	programmers.AssertEqualArray(t, Solution개인정보수집유효기간("2022.05.19", []string{}, []string{}), expect1)

}
