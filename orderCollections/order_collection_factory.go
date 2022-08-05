package orderCollections

import (
	"asctest/interfaces"
)

const MAX_ID_NUMBER = 100000

/*
    factory partern to create the collection instance
	if there is another collection type like HeapSortingCollection, the GetOrderCollection can be changed accordingly 
*/
func GetOrderCollection(collectionType string) interfaces.OrderCollection{
	if collectionType == "LinkedTable" {
		return &OrderCollectionLinkedTable{
			BuyOrderHead:  nil,
			SellOrderHead: nil,
		}
	}
	return nil
}