package interfaces

import "asctest/structs"

/*
    orders can be stored in OrderCollection
	OrderCollection is the abstraction of the buy/sell orders queue
	multiple collection can implement the interface for extending
	in this program, only OrderCollectionLinkedTable is the implementor
	but other collection can replace the OrderCollectionLinkedTable to extend
*/

type OrderCollection interface {
	GetBuyOrders()[]*structs.Order
	GetSellOrders()[]*structs.Order
	Process(order *structs.Order) []structs.Trade
}