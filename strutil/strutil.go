package strutil

// SplitLoop: 按intervalN间隔，以splitChar为界限分割str
func SplitLoop(str string, splitChar rune, interval int) (result []string) {
	hid := 0
	currentStr := ""
	for _, s := range str {
		currentStr = currentStr + string(s)
		if s == splitChar {
			hid++
			if hid%interval == 0 {
				currentStr = string([]rune(currentStr)[0 : len([]rune(currentStr))-1])
				result = append(result, currentStr)
				currentStr = ""
			}
		}
	}
	if currentStr != "" {
		result = append(result, currentStr)
	}
	return
}

// Contains: 寻找是str中否存在substr
func Contains(str string, substr string) (isMatching bool) {
	if len(substr) == 0 {
		return false
	}
	if str == substr {
		return true
	}
	if len(str) < len(substr) {
		return false
	}
	for i, s := range []rune(str) {
		if s == []rune(substr)[0] {
			if i+len([]rune(substr)) > len([]rune(str)) {
				return false
			}
			for k, v := range []rune(str)[i : i+len([]rune(substr))] {

				if v != []rune(substr)[k] {
					break
				}
				if k == len([]rune(substr))-1 {
					isMatching = true
				}
			}
		}
		if isMatching {
			break
		}
	}
	return isMatching
}

// ContainsAnd: 寻找是str中否存在substr,如果都存在，则返回true
func ContainsAnd(str string, substr ...string) (isMatching bool) {
	for _, k := range substr {
		isMatching = Contains(str, k)
		if !isMatching {
			return false
		}
	}
	return true
}

// ContainsOr: 寻找是str中否存在substr,如果至少存在一个，则返回true
func ContainsOr(str string, substr ...string) (isMatching bool) {
	for _, k := range substr {
		isMatching = Contains(str, k)
		if isMatching {
			return true
		}
	}
	return false
}
