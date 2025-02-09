package routes

import (
	"microservicecreateworkshops/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Ruta Health
	r.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	// Ruta para crear un nuevo workshop
	r.HandleFunc("/workshops", controllers.CreateWorkshop).Methods("POST")
}
