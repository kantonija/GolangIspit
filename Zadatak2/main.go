package main

import (
	"fmt"
)

type Employee struct {
	Name     string
	Position string
	Tasks    []string
}

func main() {
	Employee1 := Employee{
		Name:     "Marko",
		Position: "Programer",
		Tasks:    []string{"Programiranje", "Testiranje"},
	}

	Employee2 := Employee{
		Name:     "Ivan",
		Position: "Arhitekt",
		Tasks:    []string{"Skiciranje", "Nadogradnja", "Crtanje"},
	}

	Employee1.AddTask("Odrzavanje")
	Employee1.CompleteTasks()
	Employee2.AddTask("Odrzavanje")
	Employee2.CompleteTasks()
}

func (e *Employee) AddTask(zadatak string) {
	e.Tasks = append(e.Tasks, zadatak)
	fmt.Println(e.Tasks)
}

func (e *Employee) CompleteTasks() {
	for i := len(e.Tasks) - 1; i >= 0; i-- {
		fmt.Printf("Employee %s completed task: %s.\n", e.Name, e.Tasks[i])
		e.Tasks = append(e.Tasks[:i], e.Tasks[i+1:]...)
	}
	fmt.Println("Zadaci su:", e.Tasks)
}
