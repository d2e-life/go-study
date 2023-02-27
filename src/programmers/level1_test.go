package programmers

import (
	"testing"
)

func TestSolution둘만의암호(t *testing.T) {
	expect1 := "happy"
	if case1 := Solution둘만의암호("aukks", "wbqd", 5); case1 != expect1 {
		t.Errorf("예상 정답과 다릅니다 %s -> %s", case1, expect1)
	}
}

func TestSolution대충만든자판(t *testing.T) {

	expect1 := []int{9, 4}
	assertEqualArray(t, Solution대충만든자판([]string{"ABACD", "BCEFD"}, []string{"ABCD", "AABB"}), expect1)

	expect2 := []int{-1}
	assertEqualArray(t, Solution대충만든자판([]string{"AA"}, []string{"B"}), expect2)

	expect3 := []int{4, 6}
	assertEqualArray(t, Solution대충만든자판([]string{"AGZ", "BSSS"}, []string{"ASA", "BGZ"}), expect3)

}

func TestSolution개인정보수집유효기간(t *testing.T) {

	expect1 := []int{1, 3}
	assertEqualArray(t, Solution개인정보수집유효기간("2022.05.19", []string{}, []string{}), expect1)

}

func assertEqualArray(t *testing.T, caseResult []int, expect []int) bool {
	var result bool = true

	for i, v := range expect {
		if len(caseResult) <= i || caseResult[i] != v {
			result = false
			t.Errorf("예상 정답과 다릅니다 %v -> %v", caseResult, expect)
			break
		}
	}
	return result
}
