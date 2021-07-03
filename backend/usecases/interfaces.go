package usecases

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
