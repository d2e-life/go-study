package programmers

/*
https://school.programmers.co.kr/learn/courses/30/lessons/155652
두 문자열 s와 skip, 그리고 자연수 index가 주어질 때, 다음 규칙에 따라 문자열을 만들려 합니다. 암호의 규칙은 다음과 같습니다.

문자열 s의 각 알파벳을 index만큼 뒤의 알파벳으로 바꿔줍니다.
index만큼의 뒤의 알파벳이 z를 넘어갈 경우 다시 a로 돌아갑니다.
skip에 있는 알파벳은 제외하고 건너뜁니다.
예를 들어 s = "aukks", skip = "wbqd", index = 5일 때, a에서 5만큼 뒤에 있는 알파벳은 f지만 [b, c, d, e, f]에서 'b'와 'd'는 skip에 포함되므로 세지 않습니다. 따라서 'b', 'd'를 제외하고 'a'에서 5만큼 뒤에 있는 알파벳은 [c, e, f, g, h] 순서에 의해 'h'가 됩니다. 나머지 "ukks" 또한 위 규칙대로 바꾸면 "appy"가 되며 결과는 "happy"가 됩니다.

두 문자열 s와 skip, 그리고 자연수 index가 매개변수로 주어질 때 위 규칙대로 s를 변환한 결과를 return하도록 solution 함수를 완성해주세요.
*/
func Solution둘만의암호(s string, skip string, index int) string {

	var skipCodes []int = []int{}
	for _, skipAlphabet := range skip {
		skipCodes = append(skipCodes, int(skipAlphabet))
	}

	var answerCodes []rune = []rune{}
	for _, alphabet := range s {
		var asciiCode int = int(alphabet)

		var remain int = index
		for remain > 0 {
			asciiCode = asciiCode + 1

			if asciiCode > 122 {
				asciiCode = asciiCode%122 + 96
			}

			var isSkip bool = false
			for _, skipCode := range skipCodes {
				if asciiCode == skipCode {
					isSkip = true
					break
				}
			}

			if isSkip {
				continue
			}

			remain = remain - 1
		}

		answerCodes = append(answerCodes, rune(asciiCode))

	}
	return string(answerCodes)
}

/*
https://school.programmers.co.kr/learn/courses/30/lessons/160586
휴대폰의 자판은 컴퓨터 키보드 자판과는 다르게 하나의 키에 여러 개의 문자가 할당될 수 있습니다. 키 하나에 여러 문자가 할당된 경우, 동일한 키를 연속해서 빠르게 누르면 할당된 순서대로 문자가 바뀝니다.

예를 들어, 1번 키에 "A", "B", "C" 순서대로 문자가 할당되어 있다면 1번 키를 한 번 누르면 "A", 두 번 누르면 "B", 세 번 누르면 "C"가 되는 식입니다.

같은 규칙을 적용해 아무렇게나 만든 휴대폰 자판이 있습니다. 이 휴대폰 자판은 키의 개수가 1개부터 최대 100개까지 있을 수 있으며, 특정 키를 눌렀을 때 입력되는 문자들도 무작위로 배열되어 있습니다. 또, 같은 문자가 자판 전체에 여러 번 할당된 경우도 있고, 키 하나에 같은 문자가 여러 번 할당된 경우도 있습니다. 심지어 아예 할당되지 않은 경우도 있습니다. 따라서 몇몇 문자열은 작성할 수 없을 수도 있습니다.

이 휴대폰 자판을 이용해 특정 문자열을 작성할 때, 키를 최소 몇 번 눌러야 그 문자열을 작성할 수 있는지 알아보고자 합니다.

1번 키부터 차례대로 할당된 문자들이 순서대로 담긴 문자열배열 keymap과 입력하려는 문자열들이 담긴 문자열 배열 targets가 주어질 때, 각 문자열을 작성하기 위해 키를 최소 몇 번씩 눌러야 하는지 순서대로 배열에 담아 return 하는 solution 함수를 완성해 주세요.

단, 목표 문자열을 작성할 수 없을 때는 -1을 저장합니다.
*/
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
