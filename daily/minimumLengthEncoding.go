package daily

import (
	"Goproject/leetcode/index"
)

func MinimumLengthEncoding(words []string) int {
	trie := index.Constructor()
	for _, s := range words {
		trie.ReverseInsert(s)
	}

	mp := make(map[string]interface{})

	res := 0
	for _, s := range words {
		if _, ok := mp[s]; ok {
			continue
		}
		if trie.ReverseStartsWith(s) == false {
			res += len(s)
			res += 1
			mp[s] = struct {}{}
		}
	}
	return res
}
