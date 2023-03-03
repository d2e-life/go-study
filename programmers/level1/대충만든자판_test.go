package level1

import (
	"go-study/programmers"
	"testing"
)

func TestSolution대충만든자판(t *testing.T) {

	expect1 := []int{9, 4}
	programmers.AssertEqualArray(t, Solution대충만든자판([]string{"ABACD", "BCEFD"}, []string{"ABCD", "AABB"}), expect1)

	expect2 := []int{-1}
	programmers.AssertEqualArray(t, Solution대충만든자판([]string{"AA"}, []string{"B"}), expect2)

	expect3 := []int{4, 6}
	programmers.AssertEqualArray(t, Solution대충만든자판([]string{"AGZ", "BSSS"}, []string{"ASA", "BGZ"}), expect3)

}
