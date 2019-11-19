package leetcode

type LRUNode struct {
	prev *LRUNode
	next *LRUNode
	key  int
	val  int
}

type LRUCache struct {
	head     *LRUNode
	tail     *LRUNode
	size     int
	capacity int
	nodeMap  map[int]*LRUNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:     nil,
		tail:     nil,
		size:     0,
		capacity: capacity,
		nodeMap:  make(map[int]*LRUNode),
	}
}

func (this *LRUCache) Get(key int) int {
	lruNode := this.nodeMap[key]
	if lruNode != nil {
		// 不在尾部时，需要移动到尾部
		if lruNode != this.tail {
			if lruNode == this.head {
				this.head = lruNode.next
				lruNode.next.prev = nil
				lruNode.next = nil
			} else {
				lruNode.prev.next = lruNode.next
				lruNode.next.prev = lruNode.prev
			}
			lruNode.prev = this.tail
			this.tail.next = lruNode
			this.tail = lruNode
		}
		return lruNode.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.nodeMap[key] != nil {
		// 借助Get放到尾部
		this.nodeMap[key].val = value
		this.Get(key)
		return
	}
	if this.capacity == 1 && this.head != nil {
		this.nodeMap[this.head.key] = nil
		this.head = nil
	}
	if this.head == nil {
		this.head = &LRUNode{
			prev: nil,
			next: nil,
			key:  key,
			val:  value,
		}
		this.tail = this.head
		this.nodeMap[key] = this.head
		this.size = 1
		return
	}

	// 判断是否已满，若是，则剔除前面的值
	if this.capacity == this.size {
		this.nodeMap[this.head.key] = nil
		this.head = this.head.next
		this.head.prev = nil
	} else {
		this.size++
	}
	// 加入新节点到末尾
	node := &LRUNode{
		prev: this.tail,
		next: nil,
		key:  key,
		val:  value,
	}
	this.tail.next = node
	this.tail = node
	this.nodeMap[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
