package leetcode

type MapSum struct {
	val  int
	next []*MapSum
}

/** Initialize your data structure here. */
func Constructor677() MapSum {
	return MapSum{
		val:  0,
		next: make([]*MapSum, 32),
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	for i := 0; i < len(key); i++ {
		c := key[i] & 31
		if cur.next[c] == nil {
			cur.next[c] = &MapSum{
				val:  0,
				next: make([]*MapSum, 32),
			}
		}
		cur = cur.next[c]
	}
	cur.val = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] & 31
		if cur.next[c] == nil {
			cur.next[c] = &MapSum{
				val:  0,
				next: make([]*MapSum, 32),
			}
		}
		cur = cur.next[c]
	}
	queue := make([]*MapSum, 1, 10)
	queue[0] = cur
	sum := 0
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		sum += p.val
		for i := 0; i < len(p.next); i++ {
			if p.next[i] != nil {
				queue = append(queue, p.next[i])
			}
		}
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
