package repository

import (
	"GoFitnessApp/internal/model"
	"errors"
	"sync"
)

type ExercicioRepository interface {
	Create(exercicio *model.Exercicio) error
	GetByID(id string) (*model.Exercicio, error)
	GetAll() ([]*model.Exercicio, error)
	Update(exercicio *model.Exercicio) error
	Delete(id string) error
}

// Implementação em memória (você pode substituir por banco de dados depois)
type InMemoryExercicioRepository struct {
	exercicios map[string]*model.Exercicio
	mutex      sync.RWMutex
}

func NewInMemoryExercicioRepository() ExercicioRepository {
	return &InMemoryExercicioRepository{
		exercicios: make(map[string]*model.Exercicio),
	}
}

func (r *InMemoryExercicioRepository) Create(exercicio *model.Exercicio) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.exercicios[exercicio.ID]; exists {
		return errors.New("exercício já existe")
	}
	
	r.exercicios[exercicio.ID] = exercicio
	return nil
}

func (r *InMemoryExercicioRepository) GetByID(id string) (*model.Exercicio, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	exercicio, exists := r.exercicios[id]
	if !exists {
		return nil, errors.New("exercício não encontrado")
	}
	
	return exercicio, nil
}

func (r *InMemoryExercicioRepository) GetAll() ([]*model.Exercicio, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	exercicios := make([]*model.Exercicio, 0, len(r.exercicios))
	for _, exercicio := range r.exercicios {
		exercicios = append(exercicios, exercicio)
	}
	
	return exercicios, nil
}

func (r *InMemoryExercicioRepository) Update(exercicio *model.Exercicio) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.exercicios[exercicio.ID]; !exists {
		return errors.New("exercício não encontrado")
	}
	
	r.exercicios[exercicio.ID] = exercicio
	return nil
}

func (r *InMemoryExercicioRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.exercicios[id]; !exists {
		return errors.New("exercício não encontrado")
	}
	
	delete(r.exercicios, id)
	return nil
}
