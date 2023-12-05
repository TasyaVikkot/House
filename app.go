package main

import "fmt"

func newHouse(house House) {
	fmt.Printf("Описание дома:\n")
	fmt.Printf("Количество комнат: %d\n", house.RoomsCount)
	fmt.Printf("Количество этажей: %d\n", house.FloorsCount)
	fmt.Printf("Площадь дома: %.2f кв. м\n", house.Area)

	fmt.Println("Описание членов семьи:")
	for _, member := range house.FamilyInfo {
		fmt.Printf("- %s: %s\n", member.Member, member.Name)
	}

	fmt.Println("Описание родственников:")
	for _, member := range house.RelativesInfo {
		fmt.Printf("- %s: %s\n", member.Member, member.Name)
	}

	fmt.Println("Описание мебели:")
	for _, furniture := range house.FurnitureInfo {
		fmt.Printf("- %s: %.2f кв. м, %d шт.\n", furniture.Name, furniture.Size, furniture.Count)
	}

	fmt.Println("Описание техники:")
	for _, appliance := range house.ApplianceInfo {
		fmt.Printf("- %s: %s, %d шт.\n", appliance.Name, appliance.Brand, appliance.Count)
	}

	fmt.Println("Описание одежды:")
	for _, clothes := range house.ClothesInfo {
		fmt.Printf("- %s: %d, %d шт.\n", clothes.Name, clothes.Size, clothes.Count)
	}
}

func main() {
	HouseTasya := createHouse()
	newHouse(HouseTasya)
}
