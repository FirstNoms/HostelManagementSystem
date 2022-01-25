package hostelManagement

import (
	"fmt"
	"hostelManagement/hostelManagement"
)

func main() {
	fmt.Println("Choose a room number: ")
	var userInput int64
	fmt.Scanf("%d", &userInput)
	firstPupil:= hostelManagement.Student{"Nomso", 20, 2}
	fmt.Println(hostelManagement.createStudent(&firstPupil))
	hostelManagement.assignRoom(&userInput,)
}
