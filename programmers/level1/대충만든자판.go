package level1

// https://school.programmers.co.kr/learn/courses/30/lessons/160586
func Solution대충만든자판(keymap []string, targets []string) []int {
	var result []int = []int{}

	// keymap 의 각 key 에 접근할 때 가장 낮은 수를 알고 있는 map 을 선언
	var keymapLeast map[string]int = map[string]int{}
	for _, keys := range keymap {
		for index, key := range keys {
			if already := keymapLeast[string(key)]; already == 0 || already > index {
				keymapLeast[string(key)] = index + 1
			}
		}
	}
	// 만약 map 에 키가 없는것이 있다면 result -1
	for _, target := range targets {

		var leastApproach = 0
		for _, targetKey := range target {
			val, exist := keymapLeast[string(targetKey)]
			if !exist {
				leastApproach = -1
				break
			}
			leastApproach += val
		}

		result = append(result, leastApproach)
	}

	return result
}