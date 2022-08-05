package orderCollections

import (
	"asctest/structs"
	"sync"
)

/*
   using linked table to store buy/sell orders, the insertion and deletion time complexity is O(1)
   buy/sell orders have seperate queue
   when a buy/sell order comes in, it will try to process the order by finding out the deal ( prioritize price and time)
   unprocessed orders will be queued, the queue will follow the sequence of price and time
   this is a basic model of dealing with the trading, in the real production, this model can use like heap sorting or bucket or half-interval search
   to improve the seaching time complexity
*/
const MAX_ORDERS = 1000

type OrderCollectionLinkedTable struct {
	BuyOrderHead  *structs.OrderNode  // buy orders queue
	SellOrderHead *structs.OrderNode  // sell orders queue
	sync.RWMutex // Read Write mutex, guards access to collection.
}

/*
    process an order
	find the matched order and make the deal, untraded order will be queued
	return the trades array
*/
func (collection *OrderCollectionLinkedTable)Process(order *structs.Order) []structs.Trade{
	// make it be thread safe by using mutex
	collection.Lock()
	defer collection.Unlock()
	if order.Type == structs.OrderTypeBuy {
		return collection.processBuyOrder(order)
	}

	return collection.processSellOrder(order)
}

/*
    process a buy order
*/
func (collection *OrderCollectionLinkedTable)processBuyOrder(order *structs.Order) []structs.Trade{
	trades := make([]structs.Trade, 0, 1)
	if !collection.hasSellOrders(){
		collection.AddBuyOrder(order)
		return trades
	}
	pointer := collection.SellOrderHead
	for {
        sellOrder := pointer.Order
		// log.Println("sellOrder:", sellOrder)
		if sellOrder.Price > order.Price{
			break
		}
		if sellOrder.Amount >= order.Amount {
			trades = append(trades, structs.Trade{order.ID, sellOrder.ID, order.Amount, sellOrder.Price})
			sellOrder.Amount = sellOrder.Amount - order.Amount
			if sellOrder.Amount == 0 {
				collection.removeSellNode(pointer)
			}
			return trades
		}

		trades = append(trades, structs.Trade{order.ID, sellOrder.ID, sellOrder.Amount, sellOrder.Price})
		order.Amount -= sellOrder.Amount
		nodeToRemove := pointer
		if pointer.Next == nil{
			collection.removeSellNode(nodeToRemove)
			break
		}
		pointer = pointer.Next
		collection.removeSellNode(nodeToRemove)
		
	}
	return trades
}

/*
    process a sell order
*/
func (collection *OrderCollectionLinkedTable)processSellOrder(order *structs.Order) []structs.Trade{
	trades := make([]structs.Trade, 0, 1)
	if !collection.hasBuyOrders(){
		collection.AddSellOrder(order)
		return trades
	}
	pointer := collection.BuyOrderHead
	for {
        buyOrder := pointer.Order
		if buyOrder.Price < order.Price{
			break
		}
		if buyOrder.Amount >= order.Amount {
			trades = append(trades, structs.Trade{order.ID, buyOrder.ID, order.Amount, buyOrder.Price})
			buyOrder.Amount -= order.Amount
			if buyOrder.Amount == 0 {
				collection.removeBuyNode(pointer)
			}
			return trades
		}

		trades = append(trades, structs.Trade{order.ID, buyOrder.ID, buyOrder.Amount, buyOrder.Price})
		order.Amount -= buyOrder.Amount
		nodeToRemove := pointer
		if pointer.Next == nil{
			collection.removeBuyNode(nodeToRemove)
			break
		}
		pointer = pointer.Next
		collection.removeBuyNode(nodeToRemove)
	}
	return trades
}

// Add a buy order to the order collection
func (collection *OrderCollectionLinkedTable) AddBuyOrder(order *structs.Order) {
	orderNode := structs.NewOrderNode(order)
	if collection.BuyOrderHead == nil{
		collection.BuyOrderHead = orderNode
		orderNode.Prev = nil
		orderNode.Next = nil
		return
	}

	pointer := collection.BuyOrderHead
	for {
        buyOrder := pointer.Order
		// log.Println("buyOrder.Price ", buyOrder.Price, orderNode.Order.Price)
		if buyOrder.Price < orderNode.Order.Price {
			// log.Println("buyOrder.Price < orderNode.Order.Price")
			orderNode.Prev = pointer.Prev
			orderNode.Next = pointer
			if pointer  == collection.BuyOrderHead{
				collection.BuyOrderHead = orderNode
			}else{
				pointer.Prev.Next = orderNode
			}
			
			pointer.Prev = orderNode
			return
		}
		if pointer.Next != nil{
			pointer = pointer.Next
			continue
		}
		break
	}

	// add the node to the end of the link
	pointer.Next = orderNode
	orderNode.Prev = pointer
	orderNode.Next = nil

}

// Add a sell order to the order collection
func (collection *OrderCollectionLinkedTable) AddSellOrder(order *structs.Order) {
	orderNode := structs.NewOrderNode(order)
	if collection.SellOrderHead == nil{
		collection.SellOrderHead = orderNode
		orderNode.Next = nil
		return
	}

	pointer := collection.SellOrderHead
	for {
        sellOrder := pointer.Order
		if sellOrder.Price > orderNode.Order.Price {
			orderNode.Prev = pointer.Prev
			orderNode.Next = pointer
			if pointer  == collection.SellOrderHead{
				collection.SellOrderHead = orderNode
			}else{
				pointer.Prev.Next = orderNode
			}
			pointer.Prev = orderNode
			return
		}
		if pointer.Next != nil{
			pointer = pointer.Next
			continue
		}
		break
	}

	// add the node to the end of the link
	pointer.Next = orderNode
	orderNode.Prev = pointer
	orderNode.Next = nil
}

/*
   return all buy orders
*/
func (collection *OrderCollectionLinkedTable)GetBuyOrders()[]*structs.Order{
	collection.Lock()
	defer collection.Unlock()
	buyOrders := make([]*structs.Order, 0, MAX_ORDERS)
	pointer := collection.BuyOrderHead
	if pointer == nil{
		return buyOrders
	}
	for {
        buyOrder := pointer.Order
		buyOrders = append(buyOrders, buyOrder)
		if pointer.Next != nil{
			pointer = pointer.Next
			continue
		}
		break
	}
	return buyOrders;
}

/*
   return all sell orders
*/
func (collection *OrderCollectionLinkedTable)GetSellOrders()[]*structs.Order{
	collection.Lock()
	defer collection.Unlock()
	sellOrders := make([]*structs.Order, 0, MAX_ORDERS)
	pointer := collection.SellOrderHead
	if pointer == nil{
		return sellOrders
	}
	for {
        sellOrder := pointer.Order
		sellOrders = append(sellOrders, sellOrder)
		if pointer.Next != nil{
			pointer = pointer.Next
			continue
		}
		break
	}
	return sellOrders;
}

// Remove a buy order node
func (collection *OrderCollectionLinkedTable) removeBuyNode(node *structs.OrderNode) {
	if node == collection.BuyOrderHead{
		collection.BuyOrderHead = node.Next
		node = nil
		return
	}

	node.Prev = node.Next
	if node.Next != nil{
		node.Next = node.Prev
	}
	node = nil
}

// Remove a sell order node
func (collection *OrderCollectionLinkedTable) removeSellNode(node *structs.OrderNode) {
	if node == collection.SellOrderHead{
		collection.SellOrderHead = node.Next
		node = nil
		return
	}

	node.Prev = node.Next
	if node.Next != nil{
		node.Next = node.Prev
	}
	node = nil
}

func (collection *OrderCollectionLinkedTable) hasSellOrders()bool{
	if collection.SellOrderHead == nil{
		return false
	}
	return true
}

func (collection *OrderCollectionLinkedTable) hasBuyOrders()bool{
	if collection.BuyOrderHead == nil{
		return false
	}
	return true
}