package main

type House struct {
	RoomsCount    int
	FloorsCount   int
	Area          float64
	FamilyInfo    []Family
	RelativesInfo []Relatives
	FurnitureInfo []Furniture
	ApplianceInfo []Appliance
	ClothesInfo   []Clothes
}

func createHouse() House {
	return House{
		RoomsCount:    10,
		FloorsCount:   3,
		Area:          100.55,
		FamilyInfo:    AddFamilyInfo(),
		RelativesInfo: AddRelativesInfo(),
		FurnitureInfo: AddFurnitureInfo(),
		ApplianceInfo: AddApplianceInfo(),
		ClothesInfo:   AddClothesInfo(),
	}
}
