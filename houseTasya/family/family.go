package family

type Family struct {
	Member string
	Name   string
}

func AddFamilyInfo() []Family {
	return []Family{
		{Member: "Мать", Name: "Маша"},
		{Member: "Отец", Name: "Игорь"},
		{Member: "Старший сын", Name: "Влад"},
		{Member: "Младший сын", Name: "Миша"},
		{Member: "Дочь", Name: "Лиза"},
	}
}
