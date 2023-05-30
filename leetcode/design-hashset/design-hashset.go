// 705. Design HashSet
// https://leetcode.com/problems/design-hashset/

package designhashset

type MyHashSet struct {
	storage [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{storage: make([][]int, 2069)}
}

func (this *MyHashSet) Add(key int) {
	bucket := key % 2069
	for _, element := range this.storage[bucket] {
		if element == key {
			return
		}
	}
	this.storage[bucket] = append(this.storage[bucket], key)
}

func (this *MyHashSet) Remove(key int) {
	bucket := key % 2069
	for idx, element := range this.storage[bucket] {
		if element == key {
			this.storage[bucket] = append(this.storage[bucket][:idx], this.storage[bucket][(idx+1):]...)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	bucket := key % 2069
	for _, element := range this.storage[bucket] {
		if element == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
