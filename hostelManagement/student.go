package hostelManagement

type Student struct {
	name         string
	age          int
	matricNumber int
}

func SaveStudent(pupil Student) Student {
	return pupil
}

func NewStudent(name string, age, matricNumber int) Student {
	return Student{name, age, matricNumber}
}

func (student *Student) GetName() string {
	return student.name
}
