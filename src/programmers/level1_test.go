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

	isEqaul := func(arr1 []int, arr2 []int) bool {
		var result bool = true

		for i, v := range arr1 {
			if arr2[i] != v {
				result = false
				break
			}
		}
		return result
	}

	expect1 := []int{9, 4}
	if caseResult := Solution대충만든자판([]string{"ABACD", "BCEFD"}, []string{"ABCD", "AABB"}); !isEqaul(caseResult, expect1) {
		t.Errorf("예상 정답과 다릅니다 %v -> %v", caseResult, expect1)
	}

	expect2 := []int{-1}
	if caseResult := Solution대충만든자판([]string{"AA"}, []string{"B"}); !isEqaul(caseResult, expect2) {
		t.Errorf("예상 정답과 다릅니다 %v -> %v", caseResult, expect2)
	}

	expect3 := []int{4, 6}
	if caseResult := Solution대충만든자판([]string{"AGZ", "BSSS"}, []string{"ASA", "BGZ"}); !isEqaul(caseResult, expect3) {
		t.Errorf("예상 정답과 다릅니다 %v -> %v", caseResult, expect3)
	}
}
