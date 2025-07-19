package api

import (
	"GoFitnessApp/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExercicioHandler struct {
	service service.ExercicioService
}

func NewExercicioHandler(service service.ExercicioService) *ExercicioHandler {
	return &ExercicioHandler{
		service: service,
	}
}

type CreateExercicioRequest struct {
	Nome       string `json:"nome" binding:"required"`
	Series     int    `json:"series" binding:"required,min=1"`
	Repeticoes int    `json:"repeticoes" binding:"required,min=1"`
	Descanso   int    `json:"descanso" binding:"min=0"`
}

type UpdateExercicioRequest struct {
	Nome       string `json:"nome" binding:"required"`
	Series     int    `json:"series" binding:"required,min=1"`
	Repeticoes int    `json:"repeticoes" binding:"required,min=1"`
	Descanso   int    `json:"descanso" binding:"min=0"`
}

// POST /exercicios
func (h *ExercicioHandler) CreateExercicio(c *gin.Context) {
	var req CreateExercicioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercicio, err := h.service.CreateExercicio(req.Nome, req.Series, req.Repeticoes, req.Descanso)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exercicio)
}

// GET /exercicios/:id
func (h *ExercicioHandler) GetExercicio(c *gin.Context) {
	id := c.Param("id")
	
	exercicio, err := h.service.GetExercicio(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercicio)
}

// GET /exercicios
func (h *ExercicioHandler) GetAllExercicios(c *gin.Context) {
	exercicios, err := h.service.GetAllExercicios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercicios)
}

// PUT /exercicios/:id
func (h *ExercicioHandler) UpdateExercicio(c *gin.Context) {
	id := c.Param("id")
	
	var req UpdateExercicioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exercicio, err := h.service.UpdateExercicio(id, req.Nome, req.Series, req.Repeticoes, req.Descanso)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exercicio)
}

// DELETE /exercicios/:id
func (h *ExercicioHandler) DeleteExercicio(c *gin.Context) {
	id := c.Param("id")
	
	err := h.service.DeleteExercicio(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Registrar rotas
func (h *ExercicioHandler) RegisterRoutes(router *gin.Engine) {
	exercicios := router.Group("/exercicios")
	{
		exercicios.POST("", h.CreateExercicio)
		exercicios.GET("", h.GetAllExercicios)
		exercicios.GET("/:id", h.GetExercicio)
		exercicios.PUT("/:id", h.UpdateExercicio)
		exercicios.DELETE("/:id", h.DeleteExercicio)
	}
}