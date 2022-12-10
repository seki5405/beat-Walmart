package purchase

import (
	"cuboulder/csci5253/project/marketnode/market"
	"encoding/json"
	"math/rand"
)

type PurchaseTicket struct {
	City      string
	State     string
	Counter   int
	Cart      map[int]int
	Inventory map[int]int
}

func NewPurchaseTicket(mk market.Market, counter int) PurchaseTicket {
	purchaseTicket := PurchaseTicket{
		City:      mk.GetCity(),
		State:     mk.GetState(),
		Counter:   counter,
		Cart:      map[int]int{},
		Inventory: mk.GetInventory(),
	}
	return purchaseTicket
}

func (p *PurchaseTicket) GetPurchaseTicket() PurchaseTicket {
	return PurchaseTicket{
		City:      p.City,
		State:     p.State,
		Counter:   p.Counter,
		Cart:      p.Cart,
		Inventory: p.Inventory,
	}
}

func (p *PurchaseTicket) PutProductsInCart(numOfProduct, bias, maxBuying int) {
	for i := 0; i < numOfProduct; i++ {
		if i == bias {
			p.Cart[i] = rand.Intn(maxBuying) + (maxBuying / 2)
		} else {
			p.Cart[i] = rand.Intn(maxBuying)
		}
	}
}

func (p *PurchaseTicket) KafkaMessagePurchaseTicketInfo() string {
	jsonCartBytes, _ := json.Marshal(p)
	jsonCartString := string(jsonCartBytes)
	return jsonCartString
}
