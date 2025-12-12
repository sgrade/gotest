// 380. Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/

package insertdeletegetrandomo1

import (
	"math/rand"
)

type RandomizedSet struct {
	valToIdx map[int]int
	idxToVal []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIdx: make(map[int]int),
		idxToVal: []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.valToIdx[val]; ok {
		return false
	}

	idx := len(rs.idxToVal)
	rs.valToIdx[val] = idx
	rs.idxToVal = append(rs.idxToVal, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	valIdx, ok := rs.valToIdx[val]
	if !ok {
		return false
	}

	lastVal := rs.idxToVal[len(rs.idxToVal)-1]
	rs.valToIdx[lastVal] = valIdx
	rs.idxToVal[valIdx] = lastVal

	rs.idxToVal = rs.idxToVal[:len(rs.idxToVal)-1]
	delete(rs.valToIdx, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(rs.idxToVal))
	return rs.idxToVal[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
