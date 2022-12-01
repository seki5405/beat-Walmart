package purchase

import (
	"cuboulder/csci5253/project/marketnode/market"
	"strconv"
)

type PurchaseTicket struct {
	city    string
	state   string
	counter int
	product int
}

func NewPurchaseTicket(m market.Market, counter int, product int) PurchaseTicket {
	purchaseTicket := PurchaseTicket{
		city:    m.GetCity(),
		state:   m.GetState(),
		counter: counter,
		product: product,
	}
	return purchaseTicket
}

func (p *PurchaseTicket) GetPurchaseTicket() PurchaseTicket {
	return PurchaseTicket{
		city:    p.city,
		state:   p.state,
		counter: p.counter,
		product: p.product,
	}
}

func (p *PurchaseTicket) KafkaMessagePurchaseTicketInfo() string {
	return p.city + "," + p.state + "," + strconv.Itoa(p.counter) + "," + strconv.Itoa(p.product)
}
