package level1

// https://school.programmers.co.kr/learn/courses/30/lessons/155652
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
