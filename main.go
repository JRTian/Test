package main

import (
	"asctest/httpHandlers"
	"asctest/interfaces"
	"asctest/orderCollections"
	"log"
	"net/http"
)

const HTTP_LISTEN_PORT = "8080" // http service listen on the port

/*
   use interface here for extending the collection to other types
*/
var OrderCollection interfaces.OrderCollection

func main() {
	startService()
}

/*
    start the http service
*/
func startService(){
	initApp()
	err := startHttpService()
	if err != nil {
		log.Println("Failed to start http service, error ", err)
		return
	}
}

/*
    init the global variables
*/
func initApp() {
	OrderCollection = orderCollections.GetOrderCollection("LinkedTable")
}

func startHttpService() error{
	mux := http.NewServeMux()

	// use PostOrderHandler to handle the order input
	mux.Handle("/order", httpHandlers.NewPostOrderHandler(OrderCollection))
	err := http.ListenAndServe("localhost:" + HTTP_LISTEN_PORT, mux)

	return err
}
 