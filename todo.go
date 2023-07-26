package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
type List []item

func (l *List) Add(task string) {

	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)

}

func (l *List) Complete(i int) error {
	// list := *l
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("Item %d is not found in the list", i)
	}
	(*l)[i-1].Done = true
	(*l)[i-1].CompletedAt = time.Now()
	return nil
}

// func (l List) Delete(it item) {

// }
// func (l List) Save(it item) {

// }
func (l *List) Get() {
	fmt.Println(*l)
}
