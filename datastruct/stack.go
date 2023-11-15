package datastruct

type Stack struct {
	tail int
	arr  []int
}

// 新增元素
func (this *Stack) StackPut(num int) {
	this.tail++
	if len(this.arr) == this.tail+1 {
		tmp := make([]int, len(this.arr)*2)
		copy(tmp, this.arr)
		this.arr = tmp
	}
	this.arr[this.tail] = num
}

// 弹出根节点
func (this *Stack) StackPop() (int, bool) {
	if this.tail == -1 {
		return 0, false
	}
	val := this.arr[this.tail]
	this.tail--
	return val, true
}

// 获取栈顶元素，不出栈
func (this *Stack) StackTop() (int, bool) {
	if this.tail == -1 {
		return 0, false
	}
	val := this.arr[this.tail]
	return val, true
}

// 获取剩余栈中的元素
func (this *Stack) StackGet() []int {
	if this.tail == -1 {
		return []int{}
	}
	return this.arr[:this.tail+1]
}
func StackConstructor() Stack {
	arr := make([]int, 32)
	return Stack{-1, arr}
}
