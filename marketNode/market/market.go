package market

import "fmt"

type Market struct {
	city  string
	state string
	open  bool
	// inventory is a map of productID to inventory
	inventory map[int]int
}

func (m *Market) GetCity() string {
	return m.city
}

func (m *Market) GetState() string {
	return m.state
}

func (m *Market) GetOpen() bool {
	return m.open
}

func (m *Market) SetCity(city string) {
	m.city = city
}

func (m *Market) SetState(state string) {
	m.state = state
}

func (m *Market) SetOpen(open bool) {
	m.open = open
}

func (m *Market) GetMarket() Market {
	return Market{
		city:  m.city,
		state: m.state,
		open:  m.open,
	}
}

func NewMarket(city string, state string, open bool, numOfProduct, defaultInventory int) Market {
	inventory := make(map[int]int)
	for i := 0; i < numOfProduct; i++ {
		inventory[i] = defaultInventory
	}
	market := Market{
		city:      city,
		state:     state,
		open:      open,
		inventory: inventory,
	}
	return market
}

func (m *Market) SetMarket(city string, state string, open bool) {
	m.city = city
	m.state = state
	m.open = open
}

func (m *Market) GetInventory() map[int]int {
	return m.inventory
}

func (m *Market) PullProductFromInventory(cart map[int]int) {
	for productID, quantity := range cart {
		if m.inventory[productID] >= quantity {
			m.inventory[productID] -= quantity
		} else {
			m.inventory[productID] = 0
			fmt.Println("Out of Stock productID: ", productID)
		}
	}
}

func (m *Market) KafkaMessageMarketInfo() string {
	if m.open {
		return m.city + "," + m.state + "," + "open"
	}
	return m.city + "," + m.state + "," + "closed"
}
