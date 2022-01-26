package hostelManagement

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRoomNumbers(room [5]bool) int {
	rand.Seed(time.Now().UnixNano())
	roomNumber := rand.Intn(5)
	if room[roomNumber] == true {
		generateRoomNumbers(room)
	}
	return roomNumber
}

func AssignRoom() int {
	var room [5]bool
	roomNumber := generateRoomNumbers(room)
	fmt.Println(roomNumber)
	if roomNumber <= (len(room))-1 {
		room[roomNumber] = true
		return roomNumber + 1
	}
	return 0
}

func generateBedSpaceNumbers(bed [5]bool) int {
	rand.Seed(time.Now().UnixNano())
	bedNumber := rand.Intn(5)
	if bed[bedNumber] == true {
		generateBedSpaceNumbers(bed)
	}
	return bedNumber
}

func AssignBed() int {
	var bed [5]bool
	bedNumber := generateBedSpaceNumbers(bed)
	if bedNumber <= (len(bed))-1 {
		bed[bedNumber] = true
		return bedNumber + 1
	}
	return 0
}
