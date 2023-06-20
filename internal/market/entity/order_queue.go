package entity

type OrderQeue struct {
	Orders []*Order
}

func (oq *OrderQeue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

func (oq *OrderQeue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq *OrderQeue) Len() int {
	return len(oq.Orders)
}

func (oq *OrderQeue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

func (oq *OrderQeue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQeue {
	return &OrderQeue{}
}
