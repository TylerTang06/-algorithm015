package week02

func isAnagram(s string, t string) bool {
	myMap := map[rune]int{}
	for _, str := range s {
		if _, ok := myMap[str]; ok {
			myMap[str]++
		} else {
			myMap[str] = 1
		}
	}

	for _, str := range t {
		if _, ok := myMap[str]; ok {
			myMap[str]--
			if myMap[str] == 0 {
				delete(myMap, str)
			}
		} else {
			return false
		}
	}
	if len(myMap) != 0 {
		return false
	}

	return true
}
