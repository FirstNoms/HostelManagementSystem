package main

import (
	"fmt"
	"hostelManagement/hostelManagement"
)
import _ "hostelManagement/hostelManagement"

func main() {
	fmt.Print("Enter a name: ")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Print("Enter a age: ")
	var age int
	fmt.Scanf("%d", &age)
	fmt.Print("Enter a matric Number: ")
	var matricNumber int
	fmt.Scanf("%s", &matricNumber)

	pupil := hostelManagement.NewStudent(name, age, matricNumber)
	firstStudent := hostelManagement.SaveStudent(pupil)
	studentsRoom := hostelManagement.AssignRoom()
	studentsBed := hostelManagement.AssignBed()
	fmt.Println("Hello", firstStudent.GetName(), "\nYou have been assigned to room ", studentsRoom, "\nand bed space ", studentsBed)
}
