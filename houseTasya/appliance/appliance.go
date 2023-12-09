package appliance

type Appliance struct {
	Name  string
	Brand string
	Count int
}

func AddApplianceInfo() []Appliance {
	return []Appliance{
		{Name: "Посудомоечная машина", Brand: "LG", Count: 1},
		{Name: "Стиральная машина", Brand: "Bosch", Count: 1},
		{Name: "Холодильник", Brand: "Samsung", Count: 1},
		{Name: "Микроволновка", Brand: "Bosch", Count: 1},
		{Name: "Тостер", Brand: "Bosch", Count: 1},
	}
}
