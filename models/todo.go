package models

type Todo struct {
	id          int
	description string
	done        bool
}

func CreateTodo(desc string, todosList []Todo) {
	todo := Todo{description: desc, done: false}
	todosList = append(todosList, todo)
}
