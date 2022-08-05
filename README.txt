This program is to demo the stock trading process.

// how to run
go run main.go
but "go test" can basically show the process

// http service main.go

1. the orders will be input from http request
2. main.go will start the http service for the order input

// order process - OrderCollectionLinkedTable.go 
3. when a new order is recieved, it will check out if another queued order can make the trade happen 
4. if the full amount can be traded, it will return the trade data with the full amount 
5. if the partial amount can be traded, it will return the trade data with the partial amount and the amount left will be stored in the buy/sell orders queue 
6. the collection uses two way linked table to queue the buy/sell orders 
7. process function is thread safe to handle concurrent orders input 
8. the collection is abstracted as in orderCollection.go interface 
9. factory pattern is applied for extend the collection type

// unit test 
10. main_test.go is for testing the main.go function includes the trading processing and the thread safe input 
11. 11. order_collection_linked_table_test.go is testing order_collection_linked_table.go

In the production enviroment, the coding can be improved especially the
order searching process (heap sorting or half-interval search might be applied)
