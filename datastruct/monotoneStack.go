package datastruct

// 单调栈
type monotoneStack struct {
	arr []int
}

// 获取递增栈
func (this *monotoneStack) GetMonoResult(t string, seq string) []int {
	//t：asc 单调递增 desc 单调递减
	res := make([]int, len(this.arr))
	if t == "asc" {
		if seq == "r" {
			min := this.arr[0]
			for i := 0; i < len(res); i++ {
				if min < this.arr[i] {
					res[i] = min
				}
				if min > this.arr[i] {
					min = this.arr[i]
				}
			}
		} else {
			min := this.arr[len(res)-1]
			for i := len(res) - 1; i >= 0; i-- {
				if min < this.arr[i] {
					res[i] = min
				}
				if min > this.arr[i] {
					min = this.arr[i]
				}
			}
		}

	} else {
		if seq == "r" {
			max := this.arr[len(res)-1]
			for i := len(res) - 1; i >= 0; i-- {
				if max > this.arr[i] {
					res[i] = max
				}
				if max < this.arr[i] {
					max = this.arr[i]
				}
			}
		} else {
			max := this.arr[0]
			for i := 0; i < len(res); i++ {
				if max > this.arr[i] {
					res[i] = max
				}
				if max < this.arr[i] {
					max = this.arr[i]
				}
			}
		}

	}
	return res
}
func MonoStackConstructor(arr []int) monotoneStack {
	return monotoneStack{arr}
}
