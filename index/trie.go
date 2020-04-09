package index

type Trie struct {
	Val int
	Children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{0, [26]*Trie{}}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := range word {
		if cur.Children[word[i] - 'a'] == nil {
			cur.Children[word[i] - 'a'] = &Trie{0, [26]*Trie{}}
		}
		cur = cur.Children[word[i] - 'a']
	}
	cur.Val = 1
}

func (this *Trie) ReverseInsert(word string) {
	cur := this
	for i := len(word) - 1; i>=0; i-=1 {
		if cur.Children[word[i] - 'a'] == nil {
			cur.Children[word[i] - 'a'] = &Trie{0, [26]*Trie{}}
		}
		cur = cur.Children[word[i] - 'a']
	}
	cur.Val = 1
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := range word {
		if cur.Children[word[i] - 'a'] != nil {
			cur = cur.Children[word[i] - 'a']
		} else {
			return false
		}
	}
	if cur.Val == 1 {
		return true
	}

	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := range prefix {
		if cur.Children[prefix[i] - 'a'] != nil {
			cur = cur.Children[prefix[i] - 'a']
		} else {
			return false
		}
	}
	return true
}

func (this *Trie) ExactlyStartsWith(prefix string) bool {
	cur := this
	for i := range prefix {
		if cur.Children[prefix[i] - 'a'] != nil {
			cur = cur.Children[prefix[i] - 'a']
		} else {
			return false
		}
	}
	for i := range cur.Children {
		if cur.Children[i] != nil {
			return true
		}
	}
	return false
}

func (this *Trie) ReverseStartsWith(prefix string) bool {
	cur := this
	for i := len(prefix) - 1; i >= 0; i-=1 {
		if cur.Children[prefix[i] - 'a'] != nil {
			cur = cur.Children[prefix[i] - 'a']
		} else {
			return false
		}
	}
	for i := range cur.Children {
		if cur.Children[i] != nil {
			return true
		}
	}
	return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */