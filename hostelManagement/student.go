package hostelManagement

type Student struct {
	name         string
	age          int
	matricNumber int
}

func NewStudent(studentName string, age int, matricNumber int) *Student {
	return &Student{studentName, age, matricNumber}
}
func (student *Student) CreateStudent() string {
	studentToSave := Student{student.name, student.age, student.matricNumber}
	var students [5]Student
	var nullStudent Student
	for i := 0; i <= (len(students) - 1); i++ {
		if students[i] == nullStudent {
			students[i] = studentToSave
		}
	}
	return "That'll be all for now!"
}
