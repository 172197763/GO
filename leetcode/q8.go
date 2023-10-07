package leetcode

//https://leetcode.cn/problems/online-stock-span/description/
type StockSpanner struct {
	slink_tail *SLink
}
type SLink struct {
	val   int    //分数
	score int    //权重分值
	pre   *SLink //上一元素
}

func Constructor8() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	tmp_node := SLink{price, 1, nil}
	count := 1
	if this.slink_tail == nil {
		this.slink_tail = &tmp_node
	} else {
		tmp_node.pre = this.slink_tail
		this.slink_tail = &tmp_node
		tmp := this.slink_tail
		for true {
			if tmp.pre.val > price {
				this.slink_tail.score = count
				this.slink_tail.pre = tmp.pre
				break
			} else {
				count += tmp.pre.score
				if tmp.pre.pre == nil {
					this.slink_tail.pre = nil
					this.slink_tail.score = count
					break
				} else {
					tmp = tmp.pre
				}
			}
		}
	}
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
