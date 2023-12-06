package furniture

type Furniture struct {
	Name  string
	Size  float64
	Count int
}

func AddFurnitureInfo() []Furniture {
	return []Furniture{
		{Name: "Шкаф", Size: 2, Count: 4},
		{Name: "Стол", Size: 4, Count: 2},
		{Name: "Стул", Size: 0.8, Count: 15},
		{Name: "Кровать", Size: 6, Count: 5},
		{Name: "Ковер", Size: 8, Count: 5},
	}
}
