package queue

type Queue []*Order

type Order struct {
	priority int
	product  string
}

func (o *Order) New(priority int, product string) {
	o.priority = priority
	o.product = product
}

func (q *Queue) Add(order *Order) {
	switch len(*q) {
	case 0:
		*q = append(*q, order)
	default:
		appended := false
		var i int
		var addedOrder *Order
		// addedOrder = order
		for i, addedOrder = range *q {
			if order.priority > addedOrder.priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}
	}
}
