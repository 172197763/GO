package datastruct

import "fmt"

type heap struct {
	arr   []kv
	order string //大根堆desc 小根堆asc
}
type kv struct {
	val int
	key string
}

//新增元素
func (this *heap) HeapPut(num int, key string) {
	this.arr = append(this.arr, kv{num, key})
	this.heapUpChange(len(this.arr) - 1)
}

//弹出根节点
func (this *heap) HeapPop() (string, int) {
	k := this.arr[0].key
	v := this.arr[0].val
	this.HeapDel(0)
	return k, v
}

//获取根节点
func (this *heap) HeapRootGet() (string, int) {
	k := this.arr[0].key
	v := this.arr[0].val
	return k, v
}

//删除元素
func (this *heap) HeapDel(idx int) {
	if idx >= len(this.arr) || idx < 0 {
		return
	}
	if len(this.arr) == 1 {
		this.arr = make([]kv, 0)
		return
	}
	this.arr[idx] = this.arr[len(this.arr)-1] //末尾元素取代删除值
	this.arr = append(this.arr[:len(this.arr)-1])
	if len(this.arr) > 1 {
		this.heapDownChange(idx)
	}
}
func pow(num int, pow int) int {
	if pow == 0 {
		return 1
	}
	for pow == 0 {
		num *= num
		pow--
	}
	return num
}

//元素下更改触发调整
func (this *heap) heapDownChange(idx int) {
	if len(this.arr)-1 < idx*2+1 {
		return
	}
	left := idx*2 + 1
	right := left + 1
	min := left
	max := right
	//检查节点下的子节点
	if len(this.arr) > right {
		if this.arr[left].val > this.arr[right].val {
			min = right
			max = left
		}
	} else {
		max = left
	}
	if this.order == "asc" {
		if this.arr[min].val < this.arr[idx].val {
			tmp := this.arr[min]
			this.arr[min] = this.arr[idx]
			this.arr[idx] = tmp
			this.heapDownChange(min)
		}
	} else {
		if this.arr[max].val > this.arr[idx].val {
			tmp := this.arr[max]
			this.arr[max] = this.arr[idx]
			this.arr[idx] = tmp
			this.heapDownChange(max)
		}
	}
	return
}

//元素上更改触发调整
func (this *heap) heapUpChange(idx int) {
	p := (idx - 1) / 2
	if this.order == "asc" {
		if this.arr[p].val > this.arr[idx].val { //比父节点小则交换&继续检测
			tmp := this.arr[p]
			this.arr[p] = this.arr[idx]
			this.arr[idx] = tmp
			if p == 0 {
				return
			}
			this.heapUpChange(p)
		}
	} else {
		if this.arr[p].val < this.arr[idx].val { //比父节点大则交换&继续检测
			tmp := this.arr[p]
			this.arr[p] = this.arr[idx]
			this.arr[idx] = tmp
			if p == 0 {
				return
			}
			this.heapUpChange(p)
		}
	}
	return
}
func (this *heap) HeapShow() {
	fmt.Println(this.arr)
}

//order asc小根堆 desc大根堆
func HeapConstructor(order string) heap {
	if order != "desc" {
		order = "asc"
	}
	return heap{make([]kv, 0), order}
}
