package main

import (
	"GoFitnessApp/internal/api"
	"GoFitnessApp/internal/repository"
	"GoFitnessApp/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar repositório
	repo := repository.NewInMemoryExercicioRepository()
	
	// Inicializar service
	exercicioService := service.NewExercicioService(repo)
	
	// Inicializar handler
	exercicioHandler := api.NewExercicioHandler(exercicioService)
	
	// Configurar router
	router := gin.Default()
	
	// Registrar rotas
	exercicioHandler.RegisterRoutes(router)
	
	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "GoFitnessApp está rodando!",
		})
	})
	
	// Iniciar servidor
	log.Println("Servidor iniciando na porta 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
