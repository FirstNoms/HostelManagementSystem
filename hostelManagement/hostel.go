package hostelManagement
import (
	"fmt"
	"math/rand"
	"time"
)

type Student struct{
	name string
	age int
	matricNumber int
}

func createStudent(firstStudent *Student)int64{
	var pupils [5]int64
	var total int64
	for _, pupil:= range  pupils{
		total+= pupil
	}
	return total
}

func generateRoomNumbers(room [5]bool) int{
	rand.Seed(time.Now().UnixNano())
	roomNumber:= rand.Intn(5)
	if room[roomNumber]== false{
		generateRoomNumbers(room)
	}
	return roomNumber
}


func assignRoom(roomNumber, userInput int, room [5]bool)string{
	if userInput <= 5 && roomNumber <=5{
		room[roomNumber] = true
		fmt.Println("You Room is room ", roomNumber+1)
	}else{
		fmt.Println("You have to pick between room 1 - 5!")
	}
	return ""
}

func generateBedSpaceNumbers(bed []bool)int {
	rand.Seed(time.Now().UnixNano())
	bedNumber := rand.Intn(5)
	if bed[bedNumber] == false {
		generateBedSpaceNumbers(bed)
	}
	return bedNumber
}

func assignBed(bedNumber, userInput int, bed [5]bool)string{
	if userInput <= 5 && bedNumber <=5{
		bed[bedNumber] = true
		fmt.Println("You Bed Number is ", bedNumber+1)
	}else{
		fmt.Println("You have to pick between bed 1 - 5!")
	}
	return ""
}