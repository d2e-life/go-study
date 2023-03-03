package programmers

import "testing"

func AssertEqualArray(t *testing.T, caseResult []int, expect []int) bool {
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