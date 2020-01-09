package main

type ListNode struct {
	key   int
	value int
	prev  *ListNode
	next  *ListNode
}

type LRUCache struct {
	hashTable         map[int]*ListNode
	head              *ListNode
	tail              *ListNode
	totalItemsInCache int
	maxCapacity       int
}

func Constructor(capacity int) LRUCache {
	hashTable := make(map[int]*ListNode)
	head := &ListNode{}
	tail := &ListNode{}
	head.next = tail
	tail.prev = head
	cache := LRUCache{
		hashTable:         hashTable,
		head:              head,
		tail:              tail,
		totalItemsInCache: 0,
		maxCapacity:       capacity,
	}

	return cache
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.hashTable[key]
	if !ok {
		return -1
	}

	// Item has been accessed. Move to the front of the cache
	cache.moveToHead(node)
	return node.value
}

func (cache *LRUCache) Put(key int, value int) {
	node, ok := cache.hashTable[key]
	if ok {
		// If item is found in the cache, just update it and move it to the head of the list
		node.value = value
		cache.moveToHead(node)
	} else {
		newNode := &ListNode{
			key:   key,
			value: value,
		}

		// Add to the hashtable and the actual list that represents the cache
		cache.hashTable[key] = newNode
		cache.addToFront(newNode)
		cache.totalItemsInCache++

		if cache.totalItemsInCache > cache.maxCapacity {
			cache.removeLRUEntry()
		}
	}
}

func (cache *LRUCache) moveToHead(node *ListNode) {
	cache.removeFromList(node)
	cache.addToFront(node)
}

func (cache *LRUCache) addToFront(node *ListNode) {
	node.prev = cache.head
	node.next = cache.head.next

	cache.head.next.prev = node
	cache.head.next = node
}

func (cache *LRUCache) removeLRUEntry() {
	tail := cache.popTail()
	delete(cache.hashTable, tail.key)
	cache.totalItemsInCache--
}

func (cache *LRUCache) popTail() *ListNode {
	tail := cache.tail.prev
	cache.removeFromList(tail)
	return tail
}

func (cache *LRUCache) removeFromList(node *ListNode) {
	savedPrev := node.prev
	savedNext := node.next

	savedPrev.next = savedNext
	savedNext.prev = savedPrev
}

func main() {

}
