package main

import (
	appKafka "cuboulder/csci5253/project/marketnode/kafka"
	"cuboulder/csci5253/project/marketnode/market"
	"cuboulder/csci5253/project/marketnode/market/purchase"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	numOfCounter int = 10
	numOfProduct int = 5
)

func main() {
	fmt.Println("MarketNode started")
	// Check if CITY and STATE env variables are set and give Boulder and CO as default
	city := os.Getenv("CITY")
	if city == "" {
		city = "Boulder"
	}
	state := os.Getenv("STATE")
	if state == "" {
		state = "CO"
	}
	bias := -1
	if b := os.Getenv("BIAS"); b != "" {
		bias, _ = strconv.Atoi(b)
	}

	maxBuying := 5
	if m := os.Getenv("MAX_BUYING"); m != "" {
		maxBuying, _ = strconv.Atoi(m)
	}

	defaultInventory := 10000
	if d := os.Getenv("DEFAULT_INVENTORY"); d != "" {
		defaultInventory, _ = strconv.Atoi(d)
	}

	market := market.NewMarket(city, state, true, numOfProduct, defaultInventory)
	fmt.Println("Market info: ", market)

	for {
		since := time.Now()
		for market.GetOpen() {
			// runMarket(market, numOfProduct, bias, maxBuying)

			// // Open the market for 10 minutes
			// if time.Since(since).Minutes() > 10 {
			// 	market.SetOpen(false)
			// 	break
			// }
			var wg sync.WaitGroup
			wg.Add(numOfCounter)
			mu := sync.Mutex{}

			for i := 1; i <= numOfCounter; i++ {
				go func(counter int) {
					defer wg.Done()
					for market.GetOpen() {
						mu.Lock()
						triggerPurchase(market, counter, numOfProduct, bias, maxBuying)
						mu.Unlock()
						time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
					}
				}(i)
			}
			// if 1 minute passed, close the market
			for {
				// fmt.Println("Time Check :", time.Since(since).Minutes())
				if time.Since(since).Minutes() > 1 {
					market.SetOpen(false)
					break
				}
				time.Sleep(time.Duration(16) * time.Second)
			}
			wg.Wait()

		}
		// Close the market for 1 minute
		fmt.Println("Market closed")
		time.Sleep(time.Duration(8) * time.Minute)
	}
}

// func runMarket(market market.Market, numOfProduct, bias, maxBuying int) {
// 	var wg sync.WaitGroup
// 	wg.Add(numOfCounter)
// 	mu := sync.Mutex{}

// 	for i := 1; i <= numOfCounter; i++ {
// 		go func(counter int) {
// 			defer wg.Done()
// 			for market.GetOpen() {
// 				mu.Lock()
// 				triggerPurchase(market, counter, numOfProduct, bias, maxBuying)
// 				mu.Unlock()
// 				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
// 			}
// 		}(i)
// 	}
// 	wg.Wait()

// 	// wait for 1 minute
// 	// time.Sleep(time.Duration(1) * time.Minute)
//

func publishKafkaMessage(purchaseTicket purchase.PurchaseTicket) {
	appKafka.Publish(appKafka.KafkaTopic, purchaseTicket.KafkaMessagePurchaseTicketInfo())
}

func triggerPurchase(mk market.Market, counter, numOfProduct, bias, maxBuying int) {
	purchaseTicket := purchase.NewPurchaseTicket(mk, counter)
	purchaseTicket.PutProductsInCart(numOfProduct, bias, maxBuying)
	// Pull the products  from the market
	mk.PullProductFromInventory(purchaseTicket.Cart)
	fmt.Println("Purchase ticket info: ", purchaseTicket.KafkaMessagePurchaseTicketInfo())
	publishKafkaMessage(purchaseTicket)
}
