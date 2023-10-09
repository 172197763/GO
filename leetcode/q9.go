package leetcode

//https://leetcode.cn/problems/stock-price-fluctuation/description/?envType=daily-question&envId=2023-10-08
type StockPrice struct {
	curIdx     int
	priceMap   map[int]int
	valLink    *link
	valMaxLink *link
}
type link struct {
	idx  int
	val  int
	next *link
}

func Constructor9() StockPrice {
	obj := StockPrice{0, make(map[int]int), nil, nil}
	return obj
}

func (this *StockPrice) Update(timestamp int, price int) {
	_, ok := this.priceMap[timestamp]
	if ok {
		this.DelLinkItem(timestamp)
	}
	if timestamp > this.curIdx {
		this.curIdx = timestamp
	}
	this.AddLinkItem(timestamp, price)
	this.priceMap[timestamp] = price
}
func (this *StockPrice) DelLinkItem(timestamp int) {
	tmpLink := this.valLink
	if tmpLink.idx == timestamp {
		this.valLink = tmpLink.next
		return
	}
	for true {
		if tmpLink.next.idx == timestamp {
			if tmpLink.next.next == nil {
				tmpLink.next = nil
				this.valMaxLink = tmpLink
			} else {
				tmpLink.next = tmpLink.next.next
			}
			return
		}
		tmpLink = tmpLink.next
	}
}
func (this *StockPrice) AddLinkItem(timestamp int, price int) {
	tmp := link{timestamp, price, nil}
	if this.valLink == nil {
		this.valLink = &tmp
		this.valMaxLink = &tmp
		return
	}
	tmplink := this.valLink
	if tmplink.val > price {
		tmp.next = tmplink
		this.valLink = &tmp
		return
	}
	for true {
		if tmplink.next == nil {
			tmplink.next = &tmp
			this.valMaxLink = &tmp
			return
		}
		if tmplink.next.val > price {
			tmp.next = tmplink.next
			tmplink.next = &tmp
			return
		}
		tmplink = tmplink.next
	}

}

func (this *StockPrice) Current() int {
	return this.priceMap[this.curIdx]
}

func (this *StockPrice) Maximum() int {
	return this.valMaxLink.val
}

func (this *StockPrice) Minimum() int {
	return this.valLink.val
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
