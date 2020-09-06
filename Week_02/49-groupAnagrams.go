package week02

func groupAnagrams(strs []string) [][]string {
	if strs == nil || len(strs) == 0 {
		return [][]string{}
	}
	// count: current size of the res array
	res, count := [][]string{}, 0
	// the value is the index of anagrams string existed in res
	strsMap := map[[26]int]int{}
	for _, str := range strs {
		strsArry := [26]int{}
		for _, s := range str {
			strsArry[s-'a']++
		}
		if idex, ok := strsMap[strsArry]; ok {
			res[idex] = append(res[idex], str)
		} else {
			strsMap[strsArry] = count
			res = append(res, []string{str})
			count++
		}
	}

	return res
}
