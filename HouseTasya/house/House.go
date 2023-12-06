package house

import (
	"HouseTasya/HouseTasya/appliance"
	"HouseTasya/HouseTasya/clothes"
	"HouseTasya/HouseTasya/family"
	"HouseTasya/HouseTasya/furniture"
	"HouseTasya/HouseTasya/relatives"
	"fmt"
)

type House struct {
	RoomsCount    int
	FloorsCount   int
	Area          float64
	FamilyInfo    []family.Family
	RelativesInfo []relatives.Relatives
	FurnitureInfo []furniture.Furniture
	ApplianceInfo []appliance.Appliance
	ClothesInfo   []clothes.Clothes
}

func CreateHouse() House {
	return House{
		RoomsCount:    10,
		FloorsCount:   3,
		Area:          100.55,
		FamilyInfo:    family.AddFamilyInfo(),
		RelativesInfo: relatives.AddRelativesInfo(),
		FurnitureInfo: furniture.AddFurnitureInfo(),
		ApplianceInfo: appliance.AddApplianceInfo(),
		ClothesInfo:   clothes.AddClothesInfo(),
	}
}

func MyHouse(house House) {
	fmt.Printf("Описание дома:\n")
	fmt.Printf("Количество комнат: %d\n", house.RoomsCount)
	fmt.Printf("Количество этажей: %d\n", house.FloorsCount)
	fmt.Printf("Площадь дома: %.2f кв. м\n", house.Area)

	fmt.Println("Описание членов семьи:")
	for _, newObject := range house.FamilyInfo {
		fmt.Printf("- %s: %s\n", newObject.Member, newObject.Name)
	}

	fmt.Println("Описание родственников:")
	for _, newObject := range house.RelativesInfo {
		fmt.Printf("- %s: %s\n", newObject.Member, newObject.Name)
	}

	fmt.Println("Описание мебели:")
	for _, newObject := range house.FurnitureInfo {
		fmt.Printf("- %s: %.2f кв. м, %d шт.\n", newObject.Name, newObject.Size, newObject.Count)
	}

	fmt.Println("Описание техники:")
	for _, newObject := range house.ApplianceInfo {
		fmt.Printf("- %s: %s, %d шт.\n", newObject.Name, newObject.Brand, newObject.Count)
	}

	fmt.Println("Описание одежды:")
	for _, newObject := range house.ClothesInfo {
		fmt.Printf("- %s: %d, %d шт.\n", newObject.Name, newObject.Size, newObject.Count)
	}
}
