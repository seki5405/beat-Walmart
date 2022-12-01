package market

type Market struct {
	city  string
	state string
	open  bool
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

func NewMarket(city string, state string, open bool) Market {
	market := Market{
		city:  city,
		state: state,
		open:  open,
	}
	return market
}

func (m *Market) SetMarket(city string, state string, open bool) {
	m.city = city
	m.state = state
	m.open = open
}

func (m *Market) KafkaMessageMarketInfo() string {
	if m.open {
		return m.city + "," + m.state + "," + "open"
	}
	return m.city + "," + m.state + "," + "closed"
}
