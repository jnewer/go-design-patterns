package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strconv"
)

type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates *list.List
}

func NewEmployee(name, dept string, salary int) *Employee {
	sub := list.New()
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: sub,
	}
}

func (e *Employee) Add(emp Employee) {
	e.Subordinates.PushBack(emp)
}

func (e *Employee) Remove(emp Employee) {
	for i := e.Subordinates.Front(); i != nil; i.Next() {
		if reflect.DeepEqual(i.Value, emp) {
			e.Subordinates.Remove(i)
		}
	}
}

func (e *Employee) ToString() string {
	return "[ Name: " + e.Name + ", dept: " + e.Dept + ", salary： " + strconv.Itoa(e.Salary) + "]"
}

func main() {
	ceo := NewEmployee("kong", "ceo", 9999)
	pm := NewEmployee("ceo下属张三", "技术主管", 9000)
	programmer1 := NewEmployee("张三下属李四", "技术", 8000)
	programmer2 := NewEmployee("张三下属王五", "技术", 8000)
	minister := NewEmployee("ceo下属赵六", "财务主管", 5000)
	finance1 := NewEmployee("赵六下属陈七", "财务", 3000)
	finance2 := NewEmployee("赵六下属牛八", "财务", 2900)

	ceo.Add(*pm)
	ceo.Add(*minister)

	pm.Add(*programmer1)
	pm.Add(*programmer2)

	minister.Add(*finance1)
	minister.Add(*finance2)

	fmt.Println(ceo.ToString())

	fmt.Println("####################")
	for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee)
		fmt.Println(em.ToString())
		fmt.Println("************************")
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
		fmt.Println("----------------")
	}
}
