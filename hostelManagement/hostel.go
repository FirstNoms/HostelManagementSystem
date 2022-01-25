package hostelManagement

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRoomNumbers(room [5]bool) int {
	rand.Seed(time.Now().UnixNano())
	roomNumber := rand.Intn(5)
	if room[roomNumber] == false {
		GenerateRoomNumbers(room)
	}
	return roomNumber
}

func AssignRoom(roomNumber, userInput int, room [5]bool) string {
	if userInput <= 5 && roomNumber <= 5 {
		room[roomNumber] = true
		fmt.Println("You Room is room ", roomNumber+1)
	} else {
		fmt.Println("You have to pick between room 1 - 5!")
	}
	return ""
}

func GenerateBedSpaceNumbers(bed []bool) int {
	rand.Seed(time.Now().UnixNano())
	bedNumber := rand.Intn(5)
	if bed[bedNumber] == false {
		GenerateBedSpaceNumbers(bed)
	}
	return bedNumber
}

func AssignBed(bedNumber, userInput int, bed [5]bool) string {
	if userInput <= 5 && bedNumber <= 5 {
		bed[bedNumber] = true
		fmt.Println("You Bed Number is ", bedNumber+1)
	} else {
		fmt.Println("You have to pick between bed 1 - 5!")
	}
	return ""
}
