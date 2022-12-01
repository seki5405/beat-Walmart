package main

import (
	appKafka "cuboulder/csci5253/project/marketnode/kafka"
	"cuboulder/csci5253/project/marketnode/market"
	"cuboulder/csci5253/project/marketnode/market/purchase"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	numOfCounter int = 10
	numOfProduct int = 5
)

func main() {
	fmt.Println("MarketNode started")
	boulder_co_market := market.NewMarket("Boulder", "CO", true)
	fmt.Println("Market info: ", boulder_co_market)

	// purchase_ticket := purchase.NewPurchaseTicket(boulder_co_market, 1, 1)
	// fmt.Println("Purchase ticket info: ", purchase_ticket)

	// appKafka.Publish(appKafka.KafkaTopic, purchase_ticket.KafkaMessagePurchaseTicketInfo())

	var wg sync.WaitGroup
	wg.Add(numOfCounter)

	for i := 1; i <= numOfCounter; i++ {
		go func(counter int) {
			defer wg.Done()
			triggerPurchase(boulder_co_market, counter)
		}(i)
	}
	fmt.Println("HI")

	// After waiting 1 minute, close the market
	time.Sleep(1 * time.Minute)
	boulder_co_market.SetOpen(false)

	wg.Wait()

}

func publishKafkaMessage(purchaseTicket purchase.PurchaseTicket) {
	appKafka.Publish(appKafka.KafkaTopic, purchaseTicket.KafkaMessagePurchaseTicketInfo())
}

func triggerPurchase(market market.Market, counter int) {
	for market.GetOpen() {
		// Generate a random number between 1 and 5
		productId := rand.Intn(numOfProduct) + 1
		purchaseTicket := purchase.NewPurchaseTicket(market, counter, productId)
		fmt.Println("Purchase ticket info: ", purchaseTicket.KafkaMessagePurchaseTicketInfo())
		publishKafkaMessage(purchaseTicket)
		// Sleep for randomly generated time between 1 and 5 seconds
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}
}
