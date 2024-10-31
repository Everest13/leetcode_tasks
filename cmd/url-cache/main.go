package main

func main() {
	//test data 1
	lruCache := Constructor(2)
	//lruCache.Put(1,1) //null
	//lruCache.Put(2,2) //null
	//lruCache.Get(1) //1
	//lruCache.Put(3,3) //null
	//lruCache.Get(2) //-1
	//lruCache.Put(4,4) //null
	//lruCache.Get(1) //-1
	//lruCache.Get(3) //3
	//lruCache.Get(4) //4

	//test data 2
	//lruCache = Constructor(1)
	//lruCache.Put(2,1) //null
	//lruCache.Get(2) //1
	//lruCache.Put(3,2) //null
	//lruCache.Get(2) //1
	//lruCache.Get(3) //1

	//test data 3
	//["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get"]
	// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6]]
	//	//                   [null, null, null, null,null,  -1,   null,  19,  17,  null,  -1,  null,  null, null, -1,   null, -1,    5,   -1,   12,  null,  null,  3,   5,     5,  null,  null,  1,   null, -1,   null, 30,   5,    30,   null, null, null, -1,   null,  -1,  24,  null,  null, 18,   null, null, null, null, -1,   null, null, 18,   null,null,  -1,  null,  null, null, null, null, 18,   null, null, -1,   null,  4,    29,  30,   null,  12,   -1,  null, null, null, null, 29,   null, null, null, null,  17,   22,  18,   null, null, null, -1,   null, null, null, 20,   null, null, null,  -1,  18,   18,   null, null, null, null, 20,   null, null, null, null, null, null, null]
	//data := [][]int{{10,13},{3,17},{6,11},{10,5},{9,10},{13},{2,19},{2},{3},{5,25},{8},{9,22},{5,5},{1,30},{11},{9,12},{7},{5},{8},{9},{4,30},{9,3},{9},{10},{10},{6,14},{3,1},{3},{10,11},{8},{2,14},{1},{5},{4},{11,4},{12,24},{5,18},{13},{7,23},{8},{12},{3,27},{2,12},{5},{2,9},{13,4},{8,18},{1,7},{6},{9,29},{8,21},{5},{6,30},{1,12},{10},{4,15},{7,22},{11,26},{8,17},{9,29},{5},{3,4},{11,30},{12},{4,29},{3},{9},{6},{3,4},{1},{10},{3,29},{10,28},{1,20},{11,13},{3},{3,12},{3,8},{10,9},{3,26},{8},{7},{5},{13,17},{2,27},{11,15},{12},{9,19},{2,15},{3,16},{1},{12,17},{9,1},{6,19},{4},{5},{5},{8,1},{11,7},{5,2},{9,28},{1},{2,2},{7,4},{4,22},{7,24},{9,26},{13,28},{11,26}}
	//commands := []string{"put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"}
	//lruCache := Constructor(10)
	////                   [null, null, null, null,null,  -1,   null,  19,  17,  null,  -1,  null,  null, null, -1,   null, -1,    5,   -1,   12,  null,  null,  3,   5,     5,  null,  null,  1,   null, -1,   null, 30,   5,    30,   null, null, null, -1,   null,  -1,  24,  null,  null, 18,   null, null, null, null, -1,   null, null, 18,   null,null,  -1,  null,  null, null, null, null, 18,   null, null, -1,   null,  4,    29,  30,   null,  12,   -1,  null, null, null, null, 29,   null, null, null, null,  17,   22,  18,   null, null, null, -1,   null, null, null, 20,   null, null, null,  -1,  18,   18,   null, null, null, null, 20,   null, null, null, null, null, null, null]
	//commands := []string{"put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get"} //"put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"}
	//data := [][]int{{10,13},{3,17},{6,11},{10,5},{9,10},{13},{2,19},{2},{3},{5,25},{8},{9,22},{5,5},{1,30},{11},{9,12},{7},{5},{8},{9},{4,30},{9,3},{9},{10},{10},{6,14},{3,1},{3},{10,11},{8},{2,14},{1},{5},{4},{11,4},{12,24},{5,18},{13},{7,23},{8},{12},{3,27},{2,12},{5},{2,9},{13,4},{8,18},{1,7},{6},{9,29},{8,21},{5},{6,30},{1,12},{10},{4,15},{7,22},{11,26},{8,17},{9,29},{5},{3,4},{11,30},{12},{4,29},{3},{9},{6},{3,4},{1},{10},{3,29},{10,28},{1,20},{11,13},{3},{3,12},{3,8},{10,9},{3,26},{8},{7},{5},{13,17},{2,27},{11,15},{12},{9,19},{2,15},{3,16},{1},{12,17},{9,1},{6,19},{4},{5},{5},{8,1},{11,7},{5,2},{9,28},{1},{2,2},{7,4},{4,22},{7,24},{9,26},{13,28},{11,26}}

	//test data 4
	lruCache = Constructor(1)
	commands := []string{"get","get","put","get","put","put","put","put","get","put"}
	data := [][]int{{6},{8},{12,1},{2},{15,11},{5,2},{1,15},{4,2},{4},{15,15}}

	//test data 5
	//commands = []string{"put","get"}
	//data = [][]int{{2,1},{2}}

	//test data 6
	//commands = []string{"put","get","put","get","get"}
	//data = [][]int{{2,1},{2}, {3,2}, {2}, {3}}

	for i, com := range commands {
		dataCom := data[i]
		if com == "put" {
			lruCache.Put(dataCom[0], dataCom[1])
			continue
		}
		if com == "get" {
			lruCache.Get(dataCom[0])
		}
	}
}

/*
	146. LRU Cache
	Medium
	https://leetcode.com/problems/lru-cache/description/
	Runtime 71ms Beats 89.62%, Memory 73.05MB Beats 70.69%
*/
type LRUCache struct {
	capacity int
	nodeMap map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		nodeMap: map[int]*Node{},
	}
}

type Node struct {
	key int
	value int
	next *Node
	prev *Node
}

func (this *LRUCache) moveNode(node *Node) {
	if this.head == node {
		return
	}

	if this.tail == node {
		this.tail = node.prev
	} else {
		node.prev.next = node.next
		if node.next != nil {
			node.next.prev = node.prev
		}
	}

	node.prev = nil
	node.next = this.head
	this.head.prev = node
	this.head = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		this.moveNode(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodeMap[key]; ok {
		this.moveNode(node)
		node.value = value
		return
	}

	if len(this.nodeMap) == this.capacity {
		delete(this.nodeMap, this.tail.key)

		if this.tail == this.head {
			this.tail = nil
			this.head = nil
		} else {
			last := this.tail
			this.tail = last.prev
		}
	}

	node := &Node{
		key: key,
		value: value,
	}
	this.nodeMap[key] = node

	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}

	node.next = this.head
	this.head.prev = node
	this.head = node
}