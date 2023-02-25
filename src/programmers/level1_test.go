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
