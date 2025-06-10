package main

// Leetcode #146
type Node struct {
	Key, Value int
	Prev, Next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.remove(node)
		this.addToHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.Value = value
		this.remove(node)
		this.addToHead(node)
	} else {
		if len(this.cache) == this.capacity {
			// Remove the least recently used item
			lru := this.tail.Prev
			delete(this.cache, lru.Key)
			this.remove(lru)
		}
		newNode := &Node{Key: key, Value: value}
		this.cache[key] = newNode
		this.addToHead(newNode)
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */