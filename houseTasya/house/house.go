package house

import (
	"fmt"
	"houseTasya/houseTasya/appliance"
	"houseTasya/houseTasya/clothes"
	"houseTasya/houseTasya/family"
	"houseTasya/houseTasya/furniture"
	"houseTasya/houseTasya/relatives"
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

	printFamilyInfo("Описание членов семьи:", house.FamilyInfo)
	printRelativesInfo("Описание родственников:", house.RelativesInfo)
	printFurnitureInfo("Описание мебели:", house.FurnitureInfo)
	printApplianceInfo("Описание техники:", house.ApplianceInfo)
	printClothesInfo("Описание одежды:", house.ClothesInfo)
}

func printFamilyInfo(header string, objects []family.Family) {
	fmt.Println(header)
	for _, obj := range objects {
		fmt.Printf("- %s: %s\n", obj.Member, obj.Name)
	}
}

func printRelativesInfo(header string, objects []relatives.Relatives) {
	fmt.Println(header)
	for _, obj := range objects {
		fmt.Printf("- %s: %s\n", obj.Member, obj.Name)
	}
}

func printFurnitureInfo(header string, objects []furniture.Furniture) {
	fmt.Println(header)
	for _, obj := range objects {
		fmt.Printf("- %s: %.2f кв. м, %d шт.\n", obj.Name, obj.Size, obj.Count)
	}
}

func printApplianceInfo(header string, objects []appliance.Appliance) {
	fmt.Println(header)
	for _, obj := range objects {
		fmt.Printf("- %s: %s, %d шт.\n", obj.Name, obj.Brand, obj.Count)
	}
}

func printClothesInfo(header string, objects []clothes.Clothes) {
	fmt.Println(header)
	for _, obj := range objects {
		fmt.Printf("- %s: %d, %d шт.\n", obj.Name, obj.Size, obj.Count)
	}
}
