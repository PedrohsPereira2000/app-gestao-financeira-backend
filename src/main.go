package main

import (
	"app-gerenciamento-financeiro/src/router"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := router.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Altere isso para a origem do seu frontend
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Criar o manipulador com o middleware CORS
	handler := corsHandler.Handler(router)

	// Iniciando o servidor na porta 8080
	port := 8080
	fmt.Printf("Servidor rodando em http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
