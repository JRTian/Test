package httpHandlers

import (
	"asctest/interfaces"
	"asctest/structs"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/*
   implement ServeHTTP function to match httpHandler interface
*/
type postOrderHandler struct {
	orderCollection interfaces.OrderCollection
}

func NewPostOrderHandler (orderCollection interfaces.OrderCollection)*postOrderHandler{
    h := &postOrderHandler{
		orderCollection: orderCollection,
	}

	return h
}

/*
    handle the buy / sell order upload
*/
func (p postOrderHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("post data error", err)
		return
	}
	order := structs.NewOrder()
	order.Unmarshal(body)
	trades := p.orderCollection.Process(order)
	str, _ := json.Marshal(trades)

	// fmt.Println("order", order, trades)
	io.WriteString(resp, string(str))
	
}