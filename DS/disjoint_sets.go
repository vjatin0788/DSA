package DS

import "fmt"

type DisjointSets struct {
	Set  []int64
	Size int64
}

func (set *DisjointSets) Find(x int64) int64 {
	if set == nil {
		return -1
	}
	if !(x >= 0 && x <= set.Size) {
		return -1
	}
	if set.Set[x] == x {
		return x
	} else {
		return set.Find(set.Set[x])
	}

}

//O(mn)
func (set *DisjointSets) FastFind(x int64) int64 {
	if set == nil {
		return 0
	}
	if !(x >= 0 && x <= set.Size) {
		return 0
	}
	if set.Set[x] <= -1 {
		return x
	} else {
		return set.FastFind(set.Set[x])
	}

}

//O(n+mLogn)
func (set *DisjointSets) CompressFind(x int64) int64 {
	if set == nil {
		return 0
	}
	if !(x >= 0 && x <= set.Size) {
		return 0
	}
	if set.Set[x] <= -1 {
		return x
	} else {
		set.Set[x] = set.FastFind(set.Set[x])
		return set.Set[x]
	}

}

//O(mLogn)
func (set *DisjointSets) UnionBySize(x, y int64) {
	if set == nil {
		return
	}

	if !(x >= 0 && x <= set.Size) && !(y >= 0 && y <= set.Size) {
		return
	}

	xp := set.FastFind(x)
	yp := set.FastFind(y)

	if xp == yp {
		return
	}

	if set.Set[xp] < set.Set[yp] {
		temp := set.Set[yp]
		set.Set[yp] = xp
		set.Set[xp] += temp
	} else {
		temp := set.Set[xp]
		set.Set[xp] = yp
		set.Set[yp] += temp
	}
}

// time complexity -O(mLogn)
func (set *DisjointSets) UnionByRank(x, y int64) {
	if set == nil {
		return
	}

	if !(x >= 0 && x <= set.Size) && !(y >= 0 && y <= set.Size) {
		return
	}

	xp := set.FastFind(x)
	yp := set.FastFind(y)

	fmt.Println(set.Set[xp], set.Set[yp])

	if set.Set[xp] < set.Set[yp] {
		set.Set[yp] = xp
	} else {
		if set.Set[xp] == set.Set[yp] {
			set.Set[yp] += -1
		}
		set.Set[xp] = yp
	}
}

//O(mn)
func (set *DisjointSets) Union(x, y int64) {
	if set == nil {
		return
	}
	if set.FastFind(x) == set.FastFind(y) {
		return
	}
	if !(x >= 0 && x <= set.Size) && !(y >= 0 && y <= set.Size) {
		return
	}
	set.Set[x] = y
}

func MakeSetWithSlowFind(size int64) *DisjointSets {
	arr := make([]int64, size)
	for idx := range arr {
		arr[idx] = int64(idx)
	}
	return &DisjointSets{
		Set:  arr,
		Size: size,
	}
}

func MakeSetWithFastFind(size int64) *DisjointSets {
	arr := make([]int64, size)
	for idx := range arr {
		arr[idx] = -1
	}
	return &DisjointSets{
		Set:  arr,
		Size: size,
	}
}
