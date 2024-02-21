package main

import "fmt"

func main() {
	tries := Constructor()
	tries.Insert("fire")
	tries.Insert("firework")
	fmt.Println(tries.Search("firew"))
	fmt.Println(tries.Search("fire"))
	fmt.Println(tries.StartsWith("firew"))
}

type Trie struct {
	val  byte
	next []*Trie
}

func Constructor() Trie {
	return Trie{
		next: make([]*Trie, 0),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.next = append(this.next, nil)
		return
	}
	now := word[0]
	var matchedTrie *Trie
	for _, trie := range this.next {
		if trie == nil {
			continue
		}
		if now == trie.val {
			matchedTrie = trie
			break
		}
	}
	if matchedTrie != nil {
		matchedTrie.Insert(word[1:])
	} else {
		newTrie := &Trie{
			val:  now,
			next: make([]*Trie, 0),
		}
		this.next = append(this.next, newTrie)
		newTrie.Insert(word[1:])
	}
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		for _, trie := range this.next {
			if trie == nil {
				return true
			}
		}
		return false
	}
	now := word[0]
	for _, trie := range this.next {
		if trie == nil {
			continue
		}
		if trie.val == now {
			return trie.Search(word[1:])
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	now := prefix[0]
	for _, trie := range this.next {
		if trie == nil {
			continue
		}
		if trie.val == now {
			return trie.StartsWith(prefix[1:])
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
