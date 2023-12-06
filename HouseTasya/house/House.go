package house

import (
	"HouseTasya/HouseTasya/appliance"
	"HouseTasya/HouseTasya/clothes"
	"HouseTasya/HouseTasya/family"
	"HouseTasya/HouseTasya/furniture"
	"HouseTasya/HouseTasya/relatives"
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
