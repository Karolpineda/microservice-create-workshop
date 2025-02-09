package controllers

import (
	"encoding/json"
	"microservicecreateworkshops/config"
	"microservicecreateworkshops/models"
	"net/http"

	"github.com/google/uuid"
)

// CreateWorkshop
// @Summary Crea un nuevo workshop
// @Description Crea un nuevo workshop y lo guarda en la base de datos
// @Tags Workshops
// @Accept  json
// @Produce  json
// @Param   workshop  body  models.Workshop  true  "Datos del Workshop"
// @Success 200 {object} models.Workshop
// @Router /workshops [post]
func CreateWorkshop(w http.ResponseWriter, r *http.Request) {
	var workshop models.Workshop
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&workshop); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	workshop.ID = uuid.New()
	db := config.SetupDatabase()
	createdWorkshop := models.CreateWorkshop(db, &workshop)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkshop)
}

// HealthCheck
// @Summary Verifica el estado del microservicio
// @Description Retorna un mensaje que indica que el microservicio está en funcionamiento
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Microservice is up and running"})
}

// GetAllWorkshops
// @Summary Obtiene todos los Workshops
// @Description Devuelve una lista JSON con todos los Workshops existentes en la base de datos
// @Tags Workshops
// @Produce json
// @Success 200 {array} models.Workshop
// @Router /workshops [get]
func GetAllWorkshops(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDatabase()
	workshops := models.GetAllWorkshops(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workshops)
}
