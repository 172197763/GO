package leetcode

type Node45 struct {
	val  int
	pre  *Node45
	next *Node45
}

// https://leetcode.cn/problems/design-front-middle-back-queue/description/
type FrontMiddleBackQueue struct {
	head  *Node45
	tail  *Node45
	mid   *Node45
	count int
}

func Constructor45() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{nil, nil, nil, 0}
}
func (this *FrontMiddleBackQueue) Init(tmp_node Node45) {
	this.head = &tmp_node
	this.tail = &tmp_node
	this.mid = &tmp_node
}
func (this *FrontMiddleBackQueue) InitNil() {
	this.head = nil
	this.tail = nil
	this.mid = nil
}
func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.count++
	tmp_node := Node45{val, nil, nil}
	if this.count == 1 {
		this.Init(tmp_node)
		return
	}
	tmp_node.next = this.head
	this.head.pre = &tmp_node
	this.head = &tmp_node
	if this.count%2 == 0 {
		this.mid = this.mid.pre
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	this.count++
	tmp_node := Node45{val, nil, nil}
	if this.count == 1 {
		this.Init(tmp_node)
		return
	}
	if this.count == 2 {
		tmp_node.next = this.tail
		this.head = &tmp_node
		this.tail.pre = &tmp_node
		this.mid = &tmp_node
		return
	}

	if this.count%2 == 1 {
		tmp_node.next = this.mid.next
		this.mid.next.pre = &tmp_node
		this.mid.next = &tmp_node
		tmp_node.pre = this.mid
		this.mid = this.mid.next
	} else {
		tmp_node.next = this.mid
		tmp_node.pre = this.mid.pre
		this.mid.pre.next = &tmp_node
		this.mid.pre = &tmp_node
		this.mid = &tmp_node
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.count++
	tmp_node := Node45{val, nil, nil}
	if this.count == 1 {
		this.Init(tmp_node)
		return
	}
	this.tail.next = &tmp_node
	tmp_node.pre = this.tail
	this.tail = &tmp_node
	if this.count%2 == 1 {
		this.mid = this.mid.next
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.count == 0 {
		return -1
	}
	this.count--
	val := this.head.val
	if this.count == 0 {
		this.InitNil()
		return val
	}
	this.head = this.head.next
	this.head.pre = nil
	if this.count%2 == 1 {
		this.mid = this.mid.next
	}
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.count == 0 {
		return -1
	}
	this.count--
	val := this.mid.val
	if this.count == 0 {
		this.InitNil()
		return val
	}
	if this.count == 1 {
		this.Init(*this.tail)
		return val
	}
	tmp_pre := this.mid.pre
	tmp_next := this.mid.next
	this.mid.pre.next = tmp_next
	this.mid.next.pre = tmp_pre
	if this.count%2 == 0 {
		this.mid = tmp_pre
	} else {
		this.mid = tmp_next
	}
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.count == 0 {
		return -1
	}
	this.count--
	val := this.tail.val
	if this.count == 0 {
		this.InitNil()
		return val
	}
	this.tail = this.tail.pre
	this.tail.next = nil
	if this.count%2 == 0 {
		this.mid = this.mid.pre
	}
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
