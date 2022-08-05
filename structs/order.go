package structs

import "encoding/json"

type OrderType int32

/*
    enum order type, only buy or sell
*/
const (
	OrderTypeBuy  OrderType = 0
	OrderTypeSell OrderType = 1
)

type Order struct {
	ID     string    `json:"id"`  // the ID might not be passed through, here the ID is for demo the trading process
	Amount uint64    `json:"amount"`
	Price  float32   `json:"price"`
	Type   OrderType `json:"type"`
}

func NewOrder()*Order{
	return  &Order{
		ID    :"0",
		Amount :0,
		Price  :0,
		Type   :OrderTypeBuy,
	}
}

type OrderNode struct {
	Order *Order
	Prev  *OrderNode
	Next  *OrderNode
}

/*
    create an order node
*/
func NewOrderNode(order *Order) *OrderNode {
	orderNode := new(OrderNode)
	orderNode.Prev = nil
	orderNode.Next = nil
	orderNode.Order = copyOrder(order)
	return orderNode
}

/*
    deeply duplicate an order
*/
func copyOrder(order *Order) *Order {
	copiedOrder := new(Order)
	copiedOrder.Amount = order.Amount
	copiedOrder.ID = order.ID
	copiedOrder.Price = order.Price
	copiedOrder.Type = order.Type
	return copiedOrder
}

/*
    from json to order
*/
func (order *Order) Unmarshal(msg []byte) error {
	// should handle the err in the production
	return json.Unmarshal(msg, order)
}

/*
    from order to json
*/
func (order *Order) Marshal() []byte {
	// should handle the err in the production
	str, _ := json.Marshal(order)
	return str
}