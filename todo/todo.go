package todo

import (
	"fmt"
	"time"
)

type todoItem struct {
	Id      uint      `json:"id"`
	Msg     string    `json:"msg"`
	Detail  string    `json:"detail"`
	Done    bool      `json:"done"`
	Created time.Time `json:"created"`
}

type Todo struct {
	Items []todoItem `json:"items"`
}

func uintMax(x, y uint) uint {
	if x > y {
		return x
	} else {
		return y
	}
}

func (item todoItem) String() string {
	if item.Done {
		return fmt.Sprintf("%d Done\t%s: %s (Created: %s)",
			item.Id,
			item.Msg,
			item.Detail,
			item.Created.Format("2006-01-02 15:04"))
	} else {
		return fmt.Sprintf("%d %s: %s (Created: %s)",
			item.Id,
			item.Msg,
			item.Detail,
			item.Created.Format("2006-01-02 15:04"))
	}
}

func (todo Todo) MaxID() uint {
	var maxid uint
	for _, v := range todo.Items {
		maxid = uintMax(maxid, v.Id)
	}
	return maxid
}

func (todo Todo) FindID(id uint) (int, bool) {
	if len(todo.Items) == 0 {
		return 0, false
	}

	for i, v := range todo.Items {
		if v.Id == id {
			return i, true
		}
	}
	return 0, false
}

func (todo *Todo) AddTodo(msg, detail string) error {
	var maxid = todo.MaxID()
	if len(msg) == 0 {
		return fmt.Errorf("Message is empty.")
	}
	var new_item = todoItem{
		Id:      maxid + 1,
		Msg:     msg,
		Detail:  detail,
		Done:    false,
		Created: time.Now(),
	}
	todo.Items = append(todo.Items, new_item)
	return nil
}

func (todo Todo) PrintItems(allshow bool) {
	for _, v := range todo.Items {
		if allshow || !v.Done {
			fmt.Printf("%s\n", v)
		}
	}
}

func (todo *Todo) AllDone() int {
	if len(todo.Items) == 0 {
		return 0
	}

	var change_num = 0
	for _, v := range todo.Items {
		if !v.Done {
			v.Done = true
			change_num++
		}
	}
	return change_num
}

func (todo *Todo) SetStatus(doneid uint, status bool) error {
	idx, exist := todo.FindID(doneid)
	if !exist {
		return fmt.Errorf("%d is not found.", doneid)
	}

	todo.Items[idx].Done = status
	return nil
}

func (todo *Todo) Clean() {
	var new_items = make([]todoItem, 0, len(todo.Items))
	for _, v := range todo.Items {
		if !v.Done {
			new_items = append(new_items, v)
		}
	}
	todo.Items = new_items
}
