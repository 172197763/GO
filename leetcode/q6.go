package leetcode

//LFU算法实现
type LFUCache struct {
	cap  int
	vmap map[int]*Vmap
	link *IdexMap //索引
}

// 单链表
type IdexMap struct {
	cnt  int
	head *IdxLink //节点下的数据链表(头)
	tail *IdxLink //节点下的数据链表(尾)
	next *IdexMap //下一节点
}

// 双向链表索引
type IdxLink struct {
	cnt  int      //命中次数
	key  int      //值
	pre  *IdxLink //前指针
	next *IdxLink //后指针
}

// 哈希结构
type Vmap struct {
	num int //值
	cnt int //命中次数
}

func (this *LFUCache) Get(key int) int {
	tmp, ok := this.vmap[key]
	if ok == false {
		return -1
	}
	this.vmap[key].cnt = tmp.cnt + 1 //命中次数加一
	this.PutLink(key, this.vmap[key].cnt)
	return tmp.num
}

func (this *LFUCache) Put(key int, value int) {
	_, ok := this.vmap[key]
	if ok == true { //存在则更新值
		this.vmap[key].cnt++
		this.vmap[key].num = value
		this.PutLink(key, this.vmap[key].cnt)
		return
	}
	if len(this.vmap) == this.cap { //先删除
		delete(this.vmap, this.GetDelLinkIdx())
	}
	this.vmap[key] = &Vmap{value, 1}
	this.PutLink(key, 1)
	return
}
func (this *LFUCache) PutLink(key int, cnt int) {
	headlink := this.link
	sw := true
	var selectNode *IdexMap
	for sw {
		if headlink.cnt == cnt { //先判断是否当前节点的
			selectNode = headlink
			break
		}
		if headlink.next == nil {
			selectNode = &IdexMap{cnt, nil, nil, nil}
			headlink.next = selectNode
			break
		}
		if headlink.cnt < cnt && headlink.next.cnt > cnt {
			selectNode = &IdexMap{cnt, nil, nil, headlink.next}
			headlink.next = selectNode
			break
		}

		headlink = headlink.next
	}
	if selectNode.head == nil {
		newnode := &IdxLink{cnt, key, nil, nil}
		selectNode.head = newnode
		selectNode.tail = newnode
	} else {
		newnode := &IdxLink{cnt, key, selectNode.tail, nil} //新元素前指针指向原队尾元素
		selectNode.tail.next = newnode                      //队尾元素指向新元素
		selectNode.tail = newnode                           //更新队尾元素指针指向
	}
}
func (this *LFUCache) GetDelLinkIdx() int {
	tempNode := this.link
	var tempLinkNode *IdxLink
	res := -1
	for true {
		if tempNode.head == nil {
			tempNode = tempNode.next
			tempLinkNode = tempNode.head
			continue
		}
		if this.vmap[tempLinkNode.key].cnt == tempLinkNode.cnt { //校验是否最新的索引
			res = tempLinkNode.key
			if tempLinkNode.next == nil {
				tempNode.head = nil
				tempNode.tail = nil
			} else {
				tempNode.head = tempLinkNode.next
			}
		}
		if res > -1 {
			return res
		}
		if tempLinkNode.next == nil {
			tempNode.head = nil
			tempNode.tail = nil
		}
		tempLinkNode = tempLinkNode.next
	}
	return res
}
func Constructor(capacity int) LFUCache {
	return LFUCache{capacity, make(map[int]*Vmap, capacity), &IdexMap{0, nil, nil, nil}}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// func main() {
// 	op := []string{"LFUCache", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "get", "get", "get", "get", "get", "get", "get", "get", "get", "get", "get"}
// 	params := [][]int{
// 		{10}, {0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}, {0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10},
// 	}
// 	capacity := 10
// 	var obj LFUCache
// 	for k, v := range op {
// 		// fmt.Println(k, v, params[k][0])
// 		switch v {
// 		case "LFUCache":
// 			obj = Constructor(capacity)
// 		case "put":
// 			obj.Put(params[k][0], params[k][1])
// 			fmt.Println("null")
// 		case "get":
// 			fmt.Println(obj.Get(params[k][0]))
// 		}
// 	}
// 	// return
// 	// obj := Constructor(capacity)
// 	// fmt.Println(obj.Get(1))
// 	// obj.Put(1, 1)
// 	// fmt.Println(obj.Get(1))
// 	// fmt.Println(obj)
// 	// var vmap map[int]Vmap
// 	// vmap = make(map[int]Vmap)
// 	// vmap[1] = Vmap{2, 1}
// 	// for v := range vmap {
// 	// 	fmt.Println("idx:", v, vmap[v].num, vmap[v].cnt)
// 	// }
// 	// t, ok := vmap[1]
// 	// fmt.Println(t, ok)
// 	// t, ok = vmap[2]
// 	// fmt.Println(t, ok)
// }
