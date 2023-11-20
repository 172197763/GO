package datastruct

type NumArray struct {
	nums []int
	cal  []int //树状数组
}

func TreeArrConstructor(nums []int) NumArray {
	obj := NumArray{nums, make([]int, len(nums)+1)}
	//初始化树
	for k, v := range nums {
		obj.cal[k+1] += v
		tmp := k + 1
		for true {
			lowbit := lowbit(tmp)
			tmp += lowbit //低位数+lowbit得到覆盖的上位数
			if tmp > len(nums) {
				break
			}
			obj.cal[tmp] += v
		}
	}
	return obj
}

func (this *NumArray) Update(index int, val int) {
	sub := val - this.nums[index] //获取差值
	this.nums[index] = val        //更新
	index++
	for index <= len(this.nums) {
		lowbit := lowbit(index)
		this.cal[index] += sub
		index += lowbit //低位数+lowbit得到覆盖的上位数
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if len(this.nums) == 1 {
		return this.nums[left]
	}
	right++
	res := 0
	for right > 0 {
		lb := lowbit(right)
		res += this.cal[right]
		right -= lb
	}
	for left > 0 {
		lb := lowbit(left)
		res -= this.cal[left]
		left -= lb
	}
	return res
}

// 获取num最低位的1
func lowbit(num int) int {
	return num & -num
}
