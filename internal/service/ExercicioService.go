package service

import (
	"GoFitnessApp/internal/model"
	"GoFitnessApp/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ExercicioService interface {
	CreateExercicio(nome string, series, repeticoes, descanso int) (*model.Exercicio, error)
	GetExercicio(id string) (*model.Exercicio, error)
	GetAllExercicios() ([]*model.Exercicio, error)
	UpdateExercicio(id, nome string, series, repeticoes, descanso int) (*model.Exercicio, error)
	DeleteExercicio(id string) error
}

type exercicioService struct {
	repo repository.ExercicioRepository
}

func NewExercicioService(repo repository.ExercicioRepository) ExercicioService {
	return &exercicioService{
		repo: repo,
	}
}

func (s *exercicioService) CreateExercicio(nome string, series, repeticoes, descanso int) (*model.Exercicio, error) {
	if nome == "" {
		return nil, errors.New("nome do exercício é obrigatório")
	}
	if series <= 0 || repeticoes <= 0 || descanso < 0 {
		return nil, errors.New("valores inválidos para séries, repetições ou descanso")
	}

	exercicio := &model.Exercicio{
		ID:         uuid.New().String(),
		Nome:       nome,
		Series:     series,
		Repeticoes: repeticoes,
		Descanso:   descanso,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := s.repo.Create(exercicio)
	if err != nil {
		return nil, err
	}

	return exercicio, nil
}

func (s *exercicioService) GetExercicio(id string) (*model.Exercicio, error) {
	if id == "" {
		return nil, errors.New("ID é obrigatório")
	}

	return s.repo.GetByID(id)
}

func (s *exercicioService) GetAllExercicios() ([]*model.Exercicio, error) {
	return s.repo.GetAll()
}

func (s *exercicioService) UpdateExercicio(id, nome string, series, repeticoes, descanso int) (*model.Exercicio, error) {
	if id == "" {
		return nil, errors.New("ID é obrigatório")
	}
	if nome == "" {
		return nil, errors.New("nome do exercício é obrigatório")
	}
	if series <= 0 || repeticoes <= 0 || descanso < 0 {
		return nil, errors.New("valores inválidos para séries, repetições ou descanso")
	}

	// Verificar se existe
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Atualizar
	exercicio := &model.Exercicio{
		ID:         id,
		Nome:       nome,
		Series:     series,
		Repeticoes: repeticoes,
		Descanso:   descanso,
		CreatedAt:  existing.CreatedAt,
		UpdatedAt:  time.Now(),
	}

	err = s.repo.Update(exercicio)
	if err != nil {
		return nil, err
	}

	return exercicio, nil
}

func (s *exercicioService) DeleteExercicio(id string) error {
	if id == "" {
		return errors.New("ID é obrigatório")
	}

	return s.repo.Delete(id)
}

