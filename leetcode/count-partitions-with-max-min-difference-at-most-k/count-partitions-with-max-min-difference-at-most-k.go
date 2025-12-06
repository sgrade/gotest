package countpartitionswithmaxmindifferenceatmostk

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

// Based on Editorial's Approach 1: Sliding Window + Dynamic Programming
func countPartitions(nums []int, k int) int {
	MOD := int64(1e9 + 7)
	n := len(nums)
	validPartitions := make([]int64, n+1)
	validPartitionsForPrefix := make([]int64, n+1)
	validPartitions[0], validPartitionsForPrefix[0] = 1, 1

	elementsInWindow := redblacktree.NewWith(utils.IntComparator)

	for right, left := 0, 0; right < n; right++ {
		addTree(elementsInWindow, nums[right])
		for elementsInWindow.Right().Key.(int)-elementsInWindow.Left().Key.(int) > k {
			removeTree(elementsInWindow, nums[left])
			left++
		}
		if left > 0 {
			validPartitions[right+1] = (validPartitionsForPrefix[right] - validPartitionsForPrefix[left-1] + MOD) % MOD
		} else {
			validPartitions[right+1] = validPartitionsForPrefix[right] % MOD
		}
		validPartitionsForPrefix[right+1] = (validPartitionsForPrefix[right] + validPartitions[right+1]) % MOD
	}

	return int(validPartitions[n])
}

func addTree(tree *redblacktree.Tree, val int) {
	if cnt, found := tree.Get(val); found {
		tree.Put(val, cnt.(int)+1)
	} else {
		tree.Put(val, 1)
	}
}

func removeTree(tree *redblacktree.Tree, val int) {
	if cnt, found := tree.Get(val); found {
		if cnt.(int) > 1 {
			tree.Put(val, cnt.(int)-1)
		} else {
			tree.Remove(val)
		}
	}
}
