package main

import (
	appKafka "cuboulder/csci5253/project/marketnode/kafka"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Market information struct
// Fetch environment of "city" and "state" or "null" if not found
// Has the number of 4 different products named p1, p2, p3 and p4
type Market struct {
	city  string
	state string
	p1    int
	p2    int
	p3    int
	p4    int
	open  bool
}

// PurchaseTicket struct
// TO send through grpc to the kafka server
// Includes the market information(city, state) and the product which was purchased
type PurchaseTicket struct {
	counter int
	city    string
	state   string
	product int
}

func main() {
	// appKafka.Publish(appKafka.KafkaTopic, "MarketNode started")
	// Create a new market
	city, err := os.LookupEnv("city")
	if !err {
		city = "null"
	}
	state, err := os.LookupEnv("state")
	if !err {
		state = "null"
	}
	market := Market{
		city:  city,
		state: state,
		p1:    10,
		p2:    20,
		p3:    30,
		p4:    40,
		open:  true,
	}

	// Print market information
	fmt.Println(market)

	// Keep running the number of counters and trigger a purchase
	// Use go routines to operate the counters synchronously
	var wg sync.WaitGroup
	numOfCounter := 10
	wg.Add(numOfCounter)
	for i := 0; i < numOfCounter; i++ {
		go triggerPurchase(market, i+1)
	}

	// After 1 minute, close the market
	time.Sleep(1 * time.Minute)
	market.open = false

	appKafka.Publish(appKafka.KafkaTopic, "MarketNode started")

}

// triggerPurchase is a function that triggers a purchase
// It checkes if the numnber of products is greater than 0 before triggering the purchase
func triggerPurchase(market Market, counterNum int) {
	for market.open {
		// randomly choose a product
		pick := rand.Intn(4)
		alertFlag := false
		productToBuy := -1
		switch pick {
		case 0:
			if market.p1 > 0 {
				market.p1--
				productToBuy = 1
			} else {
				alertFlag = true
			}
		case 1:
			if market.p2 > 0 {
				market.p2--
				productToBuy = 2
			} else {
				alertFlag = true
			}
		case 2:
			if market.p3 > 0 {
				market.p3--
				productToBuy = 3
			} else {
				alertFlag = true
			}
		case 3:
			if market.p4 > 0 {
				market.p4--
				productToBuy = 4
			} else {
				alertFlag = true
			}
		}

		// if any product is out of stock, alert the user
		if alertFlag {
			fmt.Println("Out of stock")
			alertOutOfStock(market)
		} else {
			// send message to kafka 0server
			pt := PurchaseTicket{
				counter: counterNum,
				city:    market.city,
				state:   market.state,
				product: productToBuy,
			}
			// pt_marshalled, _ := json.Marshal(pt)
			// fmt.Println(string(pt_marshalled))
			fmt.Println("Purchase product ", pt.product, " on counter ", pt.counter, " in market ", pt.city, "/", pt.state)
		}

		// wait for randomly 0.5 to 2 seconds
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(1500)+500) * time.Millisecond)
	}
	// Market is closed
	fmt.Println("Market is closed")
}

// alertOutOfStock is a function that alerts the user if a product is out of stock
// It sends a message to the kafka server with grpc
func alertOutOfStock(market Market) {
	// check if any product is out of stock
	productOutOfStock := -1
	if market.p1 == 0 {
		// send message to kafka server
		productOutOfStock = 1
	}
	if market.p2 == 0 {
		// send message to kafka server
		productOutOfStock = 2
	}
	if market.p3 == 0 {
		// send message to kafka server
		productOutOfStock = 3
	}
	if market.p4 == 0 {
		// send message to kafka server
		productOutOfStock = 4
	}
	fmt.Println("Product ", productOutOfStock, " is out of stock in market ", market.city, "/", market.state)
}
