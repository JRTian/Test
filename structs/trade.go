package structs

import (
	"encoding/json"
	"fmt"
)

/*
   Trade is to describe the deal, include the order id and amount and price
*/
type Trade struct {
	ActiveOrderID string `json:"taker_order_id"`
	PassiveOrderID string `json:"maker_order_id"`
	Amount       uint64 `json:"amount"`
	Price        float32 `json:"price"`
}

 
/*
    return the string contains the trade data
*/
func (t *Trade) Serialize()string {
	return fmt.Sprintf("MakerOrderID: %s, TakerOrderID: %s,  Amount: %d, Price: %d", t.PassiveOrderID, t.ActiveOrderID, t.Amount, t.Price)
}

/*
    from json to order
*/
func (t *Trade) Unmarshal(msg []byte) error {
	// should handle the err in the production
	return json.Unmarshal(msg, t)
}

/*
    from order to json
*/
func (t *Trade) Marshal() []byte {
	// should handle the err in the production
	str, _ := json.Marshal(t)
	return str
}