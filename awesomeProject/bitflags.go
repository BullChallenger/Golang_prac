package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

func TurnLightsOn(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다.")
	}

	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다.")
	}

	if IsLightOn(rooms, BathRoom) {
		fmt.Println("욕실에 불을 켠다.")
	}

	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은 방에 불을 켠다.")
	}
}

func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	fmt.Println(rooms)
	rooms = SetLight(rooms, BathRoom)
	fmt.Println(rooms)
	rooms = SetLight(rooms, SmallRoom)
	fmt.Println(rooms)
	rooms = ResetLight(rooms, SmallRoom)
	fmt.Println(rooms)
	TurnLightsOn(rooms)
	fmt.Println(rooms)
}
