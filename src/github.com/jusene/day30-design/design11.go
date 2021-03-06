package main

import "fmt"

// 过滤器模式
type Person struct {
	Name          string
	Gender        string
	MaritalStatus string
}

func GetPerson(name string, gender string, maritalStatus string) Person {
	return Person{
		Name:          name,
		Gender:        gender,
		MaritalStatus: maritalStatus,
	}
}

type Criteria interface {
	MeetCriteria(persons []Person) []Person
}

type CriteriaMala struct {
}

// 按照性别男过滤
func (s *CriteriaMala) MeetCriteria(persons []Person) []Person {
	var femalePersions []Person
	for _, person := range persons {
		if person.Gender == "Male" {
			femalePersions = append(femalePersions, person)
		}
	}
	return femalePersions
}

type CriteriaFemale struct {
}

// 按照性别女过滤
func (s *CriteriaFemale) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, Person := range persons {
		if Person.Gender == "Female" {
			femalePersons = append(femalePersons, Person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {
}

// 按照未婚过滤
func (s *CriteriaSingle) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, person := range persons {
		if person.MaritalStatus == "Single" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

// 使用需要的过滤组合
func (s *AndCriteria) AndCriteria(criteria Criteria, otherCriteria Criteria) {
	s.criteria = criteria
	s.otherCriteria = otherCriteria
}

// 多重组合过滤 and
func (s *AndCriteria) MeetCriteria(persons []Person) []Person {
	firstCriteriaPersons := s.criteria.MeetCriteria(persons)
	return s.otherCriteria.MeetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

// 多重组合过滤 or
func (s *OrCriteria) OrCriteria(criteria Criteria, otherCriteria Criteria) {
	s.criteria = criteria
	s.otherCriteria = otherCriteria
}

func (s *OrCriteria) MeetCriteria(persons []Person) []Person {
	var p []Person
	p = append(s.criteria.MeetCriteria(persons))
	p = append(s.otherCriteria.MeetCriteria(persons))
	return p
}

func main() {
	var persons []Person
	persons = append(persons, GetPerson("Robert", "Male", "Single"))
	persons = append(persons, GetPerson("John", "Male", "Married"))
	persons = append(persons, GetPerson("Laura", "Female", "Married"))
	persons = append(persons, GetPerson("Diana", "Female", "Single"))
	persons = append(persons, GetPerson("Mike", "Male", "Single"))
	persons = append(persons, GetPerson("Bobby", "Male", "Single"))

	male := new(CriteriaMala)
	fmt.Println(male.MeetCriteria(persons))

	female := new(CriteriaFemale)
	fmt.Println(female.MeetCriteria(persons))

	single := new(CriteriaSingle)
	fmt.Println(single.MeetCriteria(persons))

	singleMale := new(AndCriteria)
	singleMale.AndCriteria(single, male)
	fmt.Println(singleMale.MeetCriteria(persons))

	singleFemale := new(OrCriteria)
	singleFemale.OrCriteria(single, male)
	fmt.Println(singleFemale.MeetCriteria(persons))
}
