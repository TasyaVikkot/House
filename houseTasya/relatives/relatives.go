package relatives

type Relatives struct {
	Member string
	Name   string
}

func AddRelativesInfo() []Relatives {
	return []Relatives{
		{Member: "Бабушка", Name: "Елена"},
		{Member: "Дедушка", Name: "Игорь"},
		{Member: "Дядя", Name: "Максим"},
		{Member: "Тетя", Name: "Ксюша"},
		{Member: "Прабабушка", Name: "Рая"},
	}
}
