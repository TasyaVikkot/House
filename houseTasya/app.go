package main

import (
	"houseTasya/houseTasya/house"
)

func main() {
	newHouse := house.CreateHouse()
	house.MyHouse(newHouse)
}
