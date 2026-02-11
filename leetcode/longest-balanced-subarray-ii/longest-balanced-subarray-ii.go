// 3721. Longest Balanced Subarray II
// https://leetcode.com/problems/longest-balanced-subarray-ii/

package longestbalancedsubarrayii

// Copied from Editorial's Approach: Prefix Sum + Segment Tree (need to work on that more)
type lazyTag struct {
	delta int
}

func (t *lazyTag) add(other *lazyTag) {
	t.delta += other.delta
}

func (t *lazyTag) has() bool {
	return t.delta != 0
}

func (t *lazyTag) clear() {
	t.delta = 0
}

type segmentTreeNode struct {
	minValue int
	maxValue int
	tag      *lazyTag
}

func newSegmentTreeNode() *segmentTreeNode {
	return &segmentTreeNode{
		minValue: 0,
		maxValue: 0,
		tag:      &lazyTag{},
	}
}

type segmentTree struct {
	n    int
	tree []*segmentTreeNode
}

func newSegmentTree(data []int) *segmentTree {
	n := len(data)
	tree := make([]*segmentTreeNode, n*4+1)
	for i := range tree {
		tree[i] = newSegmentTreeNode()
	}
	seg := &segmentTree{n: n, tree: tree}
	seg.build(data, 1, n, 1)
	return seg
}

func (seg *segmentTree) addRange(l, r, val int) {
	tag := &lazyTag{delta: val}
	seg.update(l, r, tag, 1, seg.n, 1)
}

func (seg *segmentTree) findLast(start, val int) int {
	if start > seg.n {
		return -1
	}
	return seg.find(start, seg.n, val, 1, seg.n, 1)
}

func (seg *segmentTree) applyTag(i int, tag *lazyTag) {
	seg.tree[i].minValue += tag.delta
	seg.tree[i].maxValue += tag.delta
	seg.tree[i].tag.add(tag)
}

func (seg *segmentTree) pushdown(i int) {
	if seg.tree[i].tag.has() {
		tag := &lazyTag{delta: seg.tree[i].tag.delta}
		seg.applyTag(i<<1, tag)
		seg.applyTag((i<<1)|1, tag)
		seg.tree[i].tag.clear()
	}
}

func (seg *segmentTree) pushup(i int) {
	left := seg.tree[i<<1]
	right := seg.tree[(i<<1)|1]
	seg.tree[i].minValue = min(left.minValue, right.minValue)
	seg.tree[i].maxValue = max(left.maxValue, right.maxValue)
}

func (seg *segmentTree) build(data []int, l, r, i int) {
	if l == r {
		seg.tree[i].minValue = data[l-1]
		seg.tree[i].maxValue = data[l-1]
		return
	}

	mid := l + ((r - l) >> 1)
	seg.build(data, l, mid, i<<1)
	seg.build(data, mid+1, r, (i<<1)|1)
	seg.pushup(i)
}

func (seg *segmentTree) update(targetL, targetR int, tag *lazyTag, l, r, i int) {
	if targetL <= l && r <= targetR {
		seg.applyTag(i, tag)
		return
	}

	seg.pushdown(i)
	mid := l + ((r - l) >> 1)
	if targetL <= mid {
		seg.update(targetL, targetR, tag, l, mid, i<<1)
	}
	if targetR > mid {
		seg.update(targetL, targetR, tag, mid+1, r, (i<<1)|1)
	}
	seg.pushup(i)
}

func (seg *segmentTree) find(targetL, targetR, val, l, r, i int) int {
	if seg.tree[i].minValue > val || seg.tree[i].maxValue < val {
		return -1
	}

	if l == r {
		return l
	}

	seg.pushdown(i)
	mid := l + ((r - l) >> 1)

	if targetR >= mid+1 {
		res := seg.find(targetL, targetR, val, mid+1, r, (i<<1)|1)
		if res != -1 {
			return res
		}
	}

	if l <= targetR && mid >= targetL {
		return seg.find(targetL, targetR, val, l, mid, i<<1)
	}

	return -1
}

func paritySign(x int) int {
	if x%2 == 0 {
		return 1
	}
	return -1
}

func longestBalanced(nums []int) int {
	occurrences := make(map[int][]int)

	maxLen := 0
	prefixSum := make([]int, len(nums))
	prefixSum[0] = paritySign(nums[0])
	occurrences[nums[0]] = append(occurrences[nums[0]], 1)

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1]
		occ := occurrences[nums[i]]
		if len(occ) == 0 {
			prefixSum[i] += paritySign(nums[i])
		}
		occurrences[nums[i]] = append(occ, i+1)
	}

	seg := newSegmentTree(prefixSum)
	for i := 0; i < len(nums); i++ {
		maxLen = max(maxLen, seg.findLast(i+maxLen, 0)-i)
		nextPos := len(nums) + 1
		occurrences[nums[i]] = occurrences[nums[i]][1:]
		if len(occurrences[nums[i]]) > 0 {
			nextPos = occurrences[nums[i]][0]
		}

		seg.addRange(i+1, nextPos-1, -paritySign(nums[i]))
	}

	return maxLen
}
