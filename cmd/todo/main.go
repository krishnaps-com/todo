package main

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedOn   time.Time
	CompletedOn time.Time
}

type List []item

func (l *List) Add(i item) {
	*l = append(*l, i)
}
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("The element at i doesn't exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedOn = time.Now()
	return nil
}
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("The element at %d doesn't exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}
func (l *List) listing() {
	fmt.Println((*l)[0])
}
func main() {
	task1 := item{"analyse", false, time.Now(), time.Time{}}
	task2 := item{"plan", false, time.Now(), time.Time{}}
	task3 := item{"implement", false, time.Now(), time.Time{}}
	var todo_list List
	todo_list.Add(task1)
	todo_list.Add(task2)
	todo_list.Add(task3)
	todo_list.listing()
	//fmt.Println(todo_list)
	//todo_list.Complete(1)
	//fmt.Println(todo_list)
	//todo_list.Delete(1)
	//fmt.Println(todo_list)
}
