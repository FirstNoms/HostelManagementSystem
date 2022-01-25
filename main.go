package main

import (
	"fmt"
	"hostelManagement/hostelManagement"
)
import _ "hostelManagement/hostelManagement"

func main() {
	var room [5]bool
	//var bed [5]bool
	//var bedInput int
	var roomInput int

	firstStudent := hostelManagement.NewStudent("Ada", 20, 2)
	fmt.Println(firstStudent.CreateStudent())

	fmt.Println("Choose a room number between 1 - 5: ")
	fmt.Scanf("%d", roomInput)

	roomGenerated := hostelManagement.GenerateRoomNumbers(room)

	if roomInput >= 1 && roomInput <= len(room)-1 {
		hostelManagement.AssignRoom(roomGenerated, roomInput, room)
	}

	//fmt.Println("Choose a bed number between 1 - 5:")
	//fmt.Scanf("%d", bedInput)
	//
	//bedGenerated := hostelManagement.GenerateBedSpaceNumbers(bed)
	//
	//if bedInput >= 1 && bedInput <= 5{
	//	hostelManagement.GenerateBedSpaceNumbers(bedGenerated, bedInput, bed )
	//}else{
	//	fmt.Println("We're Out of beds!")
	//}
}
