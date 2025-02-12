package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Importar el paquete docs (para Swagger) y el godotenv
	_ "microservicecreateworkshops/docs"
	"microservicecreateworkshops/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	// Importamos el paquete para manejar CORS
	"github.com/gorilla/handlers"
)

// @title Workshop API
// @version 1.0
// @description API de microservicio para gestión de workshops
// @BasePath /

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env (o no existe).")
	}

	r := mux.NewRouter()

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Usar el handler de Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Configurar CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Permite todas las URLs (ajustar en producción)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Permitir estos métodos
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Permitir estos headers
		handlers.AllowCredentials(),
	)

	// Tomar el puerto desde la variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8097" // Valor por defecto si no existe en .env
	}

	fmt.Printf("API escuchando en el puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler(r))) // Aplica el middleware CORS
}
