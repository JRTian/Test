package main

import (
	"asctest/structs"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

/*
   test main.go
*/
func TestTradingEngine_1Buy_1Sell(t *testing.T){
	// test 1 buy order and 1 sell order, basic test
	go startService()
	// wait for the service to be started, can be improved in the production
	time.Sleep((5 * time.Second)) 

	buy_order := structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}
	requestBody := bytes.NewBuffer(buy_order.Marshal())
	resp, err := http.Post(getOrderUrl(), "application/json", requestBody)
	if err != nil {
		t.Errorf("expect err is nil but %s", err.Error())
	}

	sell_order := structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}
	requestBody = bytes.NewBuffer(sell_order.Marshal())
	resp, err = http.Post(getOrderUrl(), "application/json", requestBody)
	if err != nil {
		t.Errorf("expect err is nil but %s", err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expect err is nil but %s", err.Error())
	}

	var trades []structs.Trade
	err = json.Unmarshal(respBody, &trades)
	if err != nil{
		t.Errorf("expect err is nil but %s,  respBody :%s", err.Error(), respBody)
	}
	// log.Println("return body", string(respBody))
	
	if trades[0].ActiveOrderID != "1" || trades[0].PassiveOrderID != "1" || trades[0].Amount != 100 || trades[0].Price != 10{
        t.Errorf("expected trades but got :  %s", trades[0].Serialize())
	}
}

func TestTradingEngine_100Buy_100Sell(t *testing.T){
	// launch 100 buy order and 100 sell order concurrently, for making sure the trading engine is thread safe
	requestNumber := 100
	go startHttpService()
	// wait for the service to be started, can be improved in the production
	time.Sleep((5 * time.Second)) 

	var wg sync.WaitGroup
	buy_order := structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeBuy,
	}
	var trades []structs.Trade

	for i:=0; i<requestNumber; i++{
		wg.Add(1)
		go func(seq int){
			// log.Println("seq :", seq)
			defer wg.Done()
			buyRequestBody := bytes.NewBuffer(buy_order.Marshal())
			resp, err := http.Post(getOrderUrl(), "application/json", buyRequestBody)
			if err != nil {
				log.Println("err happened ", resp)
				t.Errorf("expect err is nil but %s", err.Error())
				return
			}

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("expect err is nil but %s", err.Error())
			}

			var tradeArr []structs.Trade
			err = json.Unmarshal(respBody, &tradeArr)
			if err != nil{
				t.Errorf("expect err is nil but %s", err.Error())
			}
            if len(tradeArr)!=0{
				trades = append(trades, tradeArr[0])
			}
		}(i)
	}
	

	sell_order := structs.Order{
		ID    :"1",
		Amount :100,
		Price  :10,
		Type   :structs.OrderTypeSell,
	}
	
	for i:=0; i<requestNumber; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			sellRequestBody := bytes.NewBuffer(sell_order.Marshal())
			resp, err := http.Post(getOrderUrl(), "application/json", sellRequestBody)
			if err != nil {
				t.Errorf("expect err is nil but %s", err.Error())
			}
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("expect err is nil but %s", err.Error())
			}

			var tradeArr []structs.Trade
			err = json.Unmarshal(respBody, &tradeArr)
			if err != nil{
				t.Errorf("expect err is nil but %s", err.Error())
			}
            if len(tradeArr)!=0{
				trades = append(trades, tradeArr[0])
			}
		}()
	}

	wg.Wait()
	// log.Println("done", len(trades))
	tradesLen := len(trades)
	if tradesLen!=requestNumber{
        t.Errorf("expect %d but %d", requestNumber, tradesLen)
	}

	for i:=0; i<requestNumber; i++{
		if trades[i].Amount != 100{
			t.Errorf("expect %d but %d", requestNumber, tradesLen)
		}
	}
}

func getOrderUrl()string{
	return "http://localhost:" + HTTP_LISTEN_PORT + "/order"
}