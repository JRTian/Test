This program is to demo the stock trading process.

// how to run
go run main.go
but "go test" can basically show the process

// http service main.go

1. the orders will be input by sending http request
2. main.go will start the http service for the order input

// order process - OrderCollectionLinkedTable.go 
3. when a new order (like buy order) is recieved, it will check out if another queued order (like sell order) can make the trade happen 
4. if the full amount can be traded, it will return the trade data with the full amount 
5. if the partial amount can be traded, it will return the trade data with the partial amount and the amount left will be stored in the buy/sell orders queue 
6. the collection uses two way linked table to queue the buy/sell orders 
7. the orders will be sorted by pricing and time
   price from high to low in the buy order queue 
   price from low to high in the sell order queue
   sorting the orders will make the trading process be quicker
8. process function is thread safe to handle concurrent orders input 
9. the collection is abstracted as in orderCollection.go interface 
10. factory pattern is applied for extend the collection type

// unit test 
11. main_test.go is for testing the main.go function includes the trading processing and the thread safe input 
12. order_collection_linked_table_test.go is testing order_collection_linked_table.go

Note:
This is a basic demo of the trading system.
In the production enviroment, the coding can be improved especially the
order searching process (heap sorting or half-interval search might be applied)
more test need to be added to test it thoroughly
Now the trading process is a synchronized process but should be considered as an asynchronous one.
