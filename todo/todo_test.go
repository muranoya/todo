package todo

import (
	"testing"
	"time"
)

func TestUintMax(t *testing.T) {
	if uintMax(0, 0) != 0 {
		t.Fail()
	}
	if uintMax(1, 0) != 1 {
		t.Fail()
	}
	if uintMax(0, 1) != 1 {
		t.Fail()
	}
	if uintMax(10, 10) != 10 {
		t.Fail()
	}
}

func createTodoItems() Todo {
	var todo Todo
	var items = make([]todoItem, 0)
	items = append(items, todoItem{
		Id:      0,
		Msg:     "test_id_0",
		Detail:  "test_detail_0",
		Done:    false,
		Created: time.Now(),
	})
	items = append(items, todoItem{
		Id:      1,
		Msg:     "test_id_1",
		Detail:  "test_detail_1",
		Done:    false,
		Created: time.Now(),
	})
	items = append(items, todoItem{
		Id:      2,
		Msg:     "test_id_2",
		Detail:  "test_detail_2",
		Done:    false,
		Created: time.Now(),
	})
	items = append(items, todoItem{
		Id:      3,
		Msg:     "test_id_3",
		Detail:  "test_detail_3",
		Done:    true,
		Created: time.Now(),
	})
	items = append(items, todoItem{
		Id:      4,
		Msg:     "test_id_4",
		Detail:  "test_detail_4",
		Done:    false,
		Created: time.Now(),
	})
	todo.Items = items
	return todo
}

func createEmptyTodoItems() Todo {
	var todo Todo
	return todo
}

func TestMaxID(t *testing.T) {
	var todo = createTodoItems()
	if todo.MaxID() != 4 {
		t.Fail()
	}
}

func TestFindID(t *testing.T) {
    // existed ID
	{
		var todo = createTodoItems()
		idx, exist := todo.FindID(2)
		if idx != 2 || !exist {
			t.Fail()
		}
	}
    // Find in empty todo items
    {
        var todo = createEmptyTodoItems()
        _, exist := todo.FindID(1)
        if exist {
            t.Fail()
        }
    }
    // Not found id
    {
        var todo = createTodoItems()
        _, exist := todo.FindID(10)
        if exist {
            t.Fail()
        }
    }
}

func TestAddTodo(t *testing.T) {
    // add new item to empty todo items
    {
        var todo = createEmptyTodoItems()
        if err := todo.AddTodo("new todo", "this is for test of AddTodo"); err != nil {
            t.Fail()
        }
    }
    // add new item
    {
        var todo = createTodoItems()
        var items_len = len(todo.Items)
        if err := todo.AddTodo("new todo", "this is for test of AddTodo"); err != nil {
            t.Fail()
        }
        if items_len + 1 != len(todo.Items) {
            t.Fail()
        }
    }
    // Fail AddTodo because msg is empty
    {
        var todo = createTodoItems()
        if err := todo.AddTodo("", "this is for test of AddTodo"); err == nil {
            t.Fail()
        }
    }
}

func TestAllDone(t *testing.T) {
    // empty todo items
    {
        var todo = createEmptyTodoItems()
        var change_num = todo.AllDone()
        if change_num != 0 {
            t.Fail()
        }
    }
    // success AllDone
    {
        var todo = createTodoItems()
        var change_num = todo.AllDone()
        if change_num != 4 {
            t.Fail()
        }
    }
}

func TestSetStatus(t *testing.T) {
    // empty todo items
    {
        var todo = createEmptyTodoItems()
        if err := todo.SetStatus(0, true); err == nil {
            t.Fail()
        }
    }
    // change status to done from incomplete
    {
        var todo = createTodoItems()
        if err := todo.SetStatus(3, false); err != nil {
            t.Fail()
        }
        var change_num = todo.AllDone()
        if change_num != 5 {
            t.Fail()
        }
    }
    // change status to incomplete from done
    {
        var todo = createTodoItems()
        if err := todo.SetStatus(1, true); err != nil {
            t.Fail()
        }
        var change_num = todo.AllDone()
        if change_num != 3 {
            t.Fail()
        }
    }
}

func TestClean(t *testing.T) {
    // empty todo items
    {
        var todo = createEmptyTodoItems()
        var items_len = len(todo.Items)
        todo.Clean()
        if items_len != 0 || len(todo.Items) != 0 {
            t.Fail()
        }
    }
    // remove item
    {
        var todo = createTodoItems()
        todo.Clean()
        if len(todo.Items) != 4 {
            t.Fail()
        }
        todo.Clean()
        if len(todo.Items) != 4 {
            t.Fail()
        }
    }
}
