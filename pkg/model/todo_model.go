package model

type TodoModel struct {
	TodoId   uint64 `json:"todoid"`
	TodoText string `json:"todotext"`
	UserID   uint64 `json:"userid"`
}

func GetAllTodo() ([]TodoModel, error) {
	todoList := []TodoModel{}
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return todoList, err
	}
	for rows.Next() {
		var todo TodoModel
		err := rows.Scan(&todo.TodoId, &todo.TodoText, &todo.UserID)
		if err != nil {
			return todoList, err
		} else {
			todoList = append(todoList, todo)
		}
	}
	rows.Close()
	return todoList, err
}

func CreateTodo(todo *TodoModel) error {
	_, err := db.Exec("INSERT INTO todo (todoText,userID) VALUES (?,?)", todo.TodoText, todo.UserID)
	return err
}

func GetUserTodo(userID string) ([]TodoModel, error) {
	todoList := []TodoModel{}
	rows, err := db.Query("SELECT * FROM todo WHERE userID=?", userID)
	if err != nil {
		return todoList, err
	}
	for rows.Next() {
		var todo TodoModel
		err := rows.Scan(&todo.TodoId, &todo.TodoText, &todo.UserID)
		if err != nil {
			return todoList, err
		} else {
			todoList = append(todoList, todo)
		}
	}
	rows.Close()
	return todoList, err
}

func DeleteTodo(todoID uint64) (int64, error) {
	rows, err := db.Exec("DELETE FROM todo WHERE todoId=?", todoID)

	if err != nil {
		return 0, err
	}
	if ch, err := rows.RowsAffected(); err != nil {
		return 0, err
	} else {
		return ch, err
	}
}
