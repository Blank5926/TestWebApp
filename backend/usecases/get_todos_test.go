package usecases_test

import (
	"testing"
	"github.com/Blank5926/TestWebApp/entities"
	"github.com/Blank5926/TestWebApp/usecases"
)

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("something went wrong")
}

func TestGetTodos(t *testing.T) {
	repo := new(MockTodosRepo)
	_, err := usecases.GetTodos(repo)

	if err != usecases.ErrInternal {
		t.Fatalf("expected ErrInternal: Got: %v", err)
	}
}
