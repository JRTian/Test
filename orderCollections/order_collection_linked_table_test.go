package orderCollections

import (
	"asctest/structs"
	"testing"
)

/*
   unit test for testing OrderCollectionLinkedTable.go
*/

// test adding in buy orders
func TestLinkedTableCollection_BuyOrders(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orders := orderCollection.GetBuyOrders()
	if orders[0].ID != "1" {
        t.Errorf("expected 1 but got :  %s", orders[0].ID)
	}
	if orders[1].ID != "2" {
        t.Errorf("expected 2 but got :  %s", orders[1].ID)
	}
	
}

// test adding in buy orders
func TestLinkedTableCollection_BuyOrders_1(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :8,
		Type   :structs.OrderTypeBuy,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orders := orderCollection.GetBuyOrders()
	if orders[0].ID != "2" {
        t.Errorf("expected 2 but got :  %s", orders[0].ID )
	}
	if orders[1].ID != "1" {
        t.Errorf("expected 1 but got :  %s", orders[1].ID)
	}
	
}

// test adding in sell orders
func TestLinkedTableCollection_SellOrders(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orders := orderCollection.GetSellOrders()
	if orders[0].ID != "1" {
        t.Errorf("expected 1 but got :  %s", orders[0].ID )
	}
	if orders[1].ID != "2" {
        t.Errorf("expected 1 but got :  %s", orders[1].ID)
	}
	
	
}

// test adding in sell orders
func TestLinkedTableCollection_SellBuyOrders_1(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :8,
		Type   :structs.OrderTypeSell,
	}

	order_3 := &structs.Order{
		ID    :"3",
		Amount :100,
		Price  :9,
		Type   :structs.OrderTypeSell,
	}

	order_4 := &structs.Order{
		ID    :"4",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orderCollection.Process(order_3)
	orderCollection.Process(order_4)

	orders := orderCollection.GetSellOrders()
	if orders[0].ID != "2" {
        t.Errorf("expected 2 but got :  %s", orders[0].ID )
	}
	if orders[1].ID != "3" {
        t.Errorf("expected 3 but got :  %s", orders[1].ID)
	}

	if orders[2].ID != "1" {
        t.Errorf("expected 1 but got :  %s", orders[2].ID)
	}

	if orders[3].ID != "4" {
        t.Errorf("expected 4 but got :  %s", orders[3].ID)
	}
	
	
}

// test processing orders, assuming the full amount can be traded in one deal
func TestLinkedTableCollection_BuyProcessing_FullAmountDeal(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :8,
		Type   :structs.OrderTypeSell,
	}

	order_3 := &structs.Order{
		ID    :"3",
		Amount :100,
		Price  :9,
		Type   :structs.OrderTypeSell,
	}

	order_4 := &structs.Order{
		ID    :"4",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orderCollection.Process(order_3)
	orderCollection.Process(order_4)

	buy_order := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}
	trades := orderCollection.Process(buy_order)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "2" || trades[0].Amount != 100 || trades[0].Price != 8{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
}

// test processing buy orders, assuming the partial amount can be traded in one deal
// simulate when the buy orders come in and be traded with the existing sell orders
func TestLinkedTableCollection_BuyProcessing_PartialAmountDeal(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :8,
		Type   :structs.OrderTypeSell,
	}

	order_3 := &structs.Order{
		ID    :"3",
		Amount :100,
		Price  :9,
		Type   :structs.OrderTypeSell,
	}

	order_4 := &structs.Order{
		ID    :"4",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orderCollection.Process(order_3)
	orderCollection.Process(order_4)

	buy_order := &structs.Order{
		ID    :"1",
		Amount :150, //multiple trades
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}
	trades := orderCollection.Process(buy_order)
	
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "2" || trades[0].Amount != 100 || trades[0].Price != 8{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
	if trades[1].ActiveOrderID != "1" || trades[1].PassiveOrderID != "3" || trades[1].Amount != 50 || trades[1].Price != 9{
        t.Errorf("expected trades but got :  %s", trades[1].Serialize())
	}

	sellOrders := orderCollection.GetSellOrders()
	if sellOrders[0].ID != "3" || sellOrders[0].Amount != 50{
        t.Errorf("err sellOrders[0]")
	}
	if sellOrders[1].ID != "1" || sellOrders[1].Amount != 100{
        t.Errorf("err sellOrders[1]")
	}
	if sellOrders[2].ID != "4" || sellOrders[2].Amount != 100{
        t.Errorf("err sellOrders[2]")
	}

	buy_order_1 := &structs.Order{
		ID    :"1",
		Amount :100, //multiple trades
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	trades = orderCollection.Process(buy_order_1)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "3" || trades[0].Amount != 50 || trades[0].Price != 9{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
	if trades[1].ActiveOrderID != "1" || trades[1].PassiveOrderID != "1" || trades[1].Amount != 50 || trades[1].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[1].Serialize())
	}
	
	
	buy_order_2 := &structs.Order{
		ID    :"1",
		Amount :150, //multiple trades
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	trades = orderCollection.Process(buy_order_2)
	// log.Println("trades", trades)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "1" || trades[0].Amount != 50 || trades[0].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
	if trades[1].ActiveOrderID != "1" || trades[1].PassiveOrderID != "4" || trades[1].Amount != 100 || trades[1].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[1].Serialize())
	}

	sellOrders = orderCollection.GetSellOrders()
	if len(sellOrders) != 0{
		t.Errorf("expecte 0 but %d:", len(sellOrders))
	}
	
}

// test processing sell orders, assuming the partial amount can be traded in one deal
// simulate when the sell orders come in and be traded with the existing buy orders
func TestLinkedTableCollection_SellProcessing_FullAmountDeal(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :8,
		Type   :structs.OrderTypeBuy,
	}

	order_3 := &structs.Order{
		ID    :"3",
		Amount :100,
		Price  :9,
		Type   :structs.OrderTypeBuy,
	}

	order_4 := &structs.Order{
		ID    :"4",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orderCollection.Process(order_3)
	orderCollection.Process(order_4)

	sell_order := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}
	trades := orderCollection.Process(sell_order)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "1" || trades[0].Amount != 100 || trades[0].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
}

func TestLinkedTableCollection_SellProcessing_PartAmountDeal(t *testing.T){
	order_1 := &structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	order_2 := &structs.Order{
		ID    :"2",
		Amount :100,
		Price  :8,
		Type   :structs.OrderTypeBuy,
	}

	order_3 := &structs.Order{
		ID    :"3",
		Amount :100,
		Price  :9,
		Type   :structs.OrderTypeBuy,
	}

	order_4 := &structs.Order{
		ID    :"4",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}

	orderCollection := GetOrderCollection("LinkedTable")
	orderCollection.Process(order_1)
	orderCollection.Process(order_2)
	orderCollection.Process(order_3)
	orderCollection.Process(order_4)

	sell_order := &structs.Order{
		ID    :"1",
		Amount :150, //multiple trades
		Price  :10,
		Type   :structs.OrderTypeSell,
	}
	trades := orderCollection.Process(sell_order)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "1" || trades[0].Amount != 100 || trades[0].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
	if trades[1].ActiveOrderID != "1" || trades[1].PassiveOrderID != "4" || trades[1].Amount != 50 || trades[1].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[1].Serialize())
	}

	buyOrders := orderCollection.GetBuyOrders()
	// log.Println("sellOrders", len(buyOrders), buyOrders[0],  buyOrders[1], buyOrders[2])
	if buyOrders[0].ID != "4" || buyOrders[0].Amount != 50{
        t.Errorf("err sellOrders[0]")
	}
	if buyOrders[1].ID != "3" || buyOrders[1].Amount != 100{
        t.Errorf("err sellOrders[1]")
	}
	if buyOrders[2].ID != "2" || buyOrders[2].Amount != 100{
        t.Errorf("err sellOrders[2]")
	}

	sell_order_1 := &structs.Order{
		ID    :"1",
		Amount :100, //multiple trades
		Price  :9,
		Type   :structs.OrderTypeSell,
	}

	trades = orderCollection.Process(sell_order_1)
	// log.Println("trades", trades)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "4" || trades[0].Amount != 50 || trades[0].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
	if trades[1].ActiveOrderID != "1" || trades[1].PassiveOrderID != "3" || trades[1].Amount != 50 || trades[1].Price != 9{
        t.Errorf("expected trades but got :  %s", trades[1].Serialize())
	}
	
	sell_order_2 := &structs.Order{
		ID    :"1",
		Amount :150, //multiple trades
		Price  :7,
		Type   :structs.OrderTypeSell,
	}

	trades = orderCollection.Process(sell_order_2)
	// log.Println("trades", trades)
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "3" || trades[0].Amount != 50 || trades[0].Price != 9{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
	if trades[1].ActiveOrderID != "1" || trades[1].PassiveOrderID != "2" || trades[1].Amount != 100 || trades[1].Price != 8{
        t.Errorf("expected trades but got :  %s", trades[1].Serialize())
	}

	buyOrders = orderCollection.GetBuyOrders()
	if len(buyOrders) != 0{
		t.Errorf("expecte 0 but %d:", len(buyOrders))
	}

	sell_order_3 := &structs.Order{
		ID    :"1",
		Amount :150, //multiple trades
		Price  :7,
		Type   :structs.OrderTypeSell,
	}
	trades = orderCollection.Process(sell_order_3)
	if len(trades)!=0{
		t.Errorf("expect 0 but : %d", len(trades))
	}
	
    sellOrders := orderCollection.GetSellOrders()
	if len(sellOrders) != 1{
		t.Errorf("expecte 1 but %d:", len(sellOrders))
	}
}