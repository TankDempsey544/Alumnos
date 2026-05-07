package main

import "fmt"

type Subject struct {
	Name  string
	Grade float64
}

type Student struct {
	Name     string
	Subjects map[string]Subject
}

func (s *Student) AddSubject(name string, grade float64) {
	if s.Subjects == nil {
		s.Subjects = make(map[string]Subject)
	}
	s.Subjects[name] = Subject{
		Name:  name,
		Grade: grade,
	}
}

func (s Student) Average() float64 {
	total := 0.0
	for _, sub := range s.Subjects {
		total += sub.Grade
	}
	if len(s.Subjects) == 0 {
		return 0
	}
	return total / float64(len(s.Subjects))
}

func (s Student) Print() {
	fmt.Println("Alumno:", s.Name)
	for _, sub := range s.Subjects {
		fmt.Println("-", sub.Name, ":", sub.Grade)
	}
	fmt.Println("Promedio:", s.Average())
}

func (s Student) IsPassing() bool {
	return s.Average() >= 70
}
func (s Student) FindSubject(name string) (Subject, bool) {
	sub, found := s.Subjects[name]
	return sub, found
}

func (s *Student) RemoveSubject(name string) bool {
	if _, found := s.Subjects[name]; !found {
		return false
	}
	delete(s.Subjects, name)
	return true
}

func (s *Student) UpdateGrade(name string, newGrade float64) bool {
	sub, found := s.Subjects[name]
	if !found {
		return false
	}
	sub.Grade = newGrade
	s.Subjects[name] = sub
	return true
}

func (s Student) TopSubject() (Subject, bool) {
	if len(s.Subjects) == 0 {
		return Subject{}, false
	}
	var top Subject
	for _, sub := range s.Subjects {
		if sub.Grade > top.Grade {
			top = sub
		}
	}
	return top, true
}

func main() {
	student := Student{Name: "Carlos"}
	student.AddSubject("Matemáticas", 90)
	student.AddSubject("Programación", 95)
	student.AddSubject("Física", 80)
	student.Print()

	fmt.Println("\n¿Aprobado?", student.IsPassing())

	if sub, found := student.FindSubject("Física"); found {
		fmt.Println("\nMateria encontrada:", sub.Name, "- Calificación:", sub.Grade)
	}

	student.UpdateGrade("Física", 72)
	fmt.Println("\nFísica actualizada:")
	student.Print()

	student.RemoveSubject("Matemáticas")
	fmt.Println("\nDespués de eliminar Matemáticas:")
	student.Print()

	if top, ok := student.TopSubject(); ok {
		fmt.Println("\nMateria con mayor calificación:", top.Name, "-", top.Grade)
	}
}
