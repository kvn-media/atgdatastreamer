// internal/controllers/dataTank_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kvn-media/atgdatastreamer/internal/models"
	"github.com/kvn-media/atgdatastreamer/internal/usecase"
)

type DataTankController interface {
	CreateDataTank(w http.ResponseWriter, r *http.Request)
	GetDataTanks(w http.ResponseWriter, r *http.Request)
	UpdateDataTank(w http.ResponseWriter, r *http.Request)
	DeleteDataTank(w http.ResponseWriter, r *http.Request)
	ReadFromSerial(w http.ResponseWriter, r *http.Request)
}

type dataTankController struct {
	DataTankUsecase usecase.DataTankUsecase
}

// NewDataTankController creates a new instance of DataTankController
func NewDataTankController(usecase usecase.DataTankUsecase) DataTankController {
	return &dataTankController{
		DataTankUsecase: usecase,
	}
}

func (c *dataTankController) CreateDataTank(w http.ResponseWriter, r *http.Request) {
	var dataTank models.DataTank
	err := json.NewDecoder(r.Body).Decode(&dataTank)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = c.DataTankUsecase.CreateDataTank(&dataTank)
	if err != nil {
		http.Error(w, "Failed to create DataTank", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dataTank)
}

func (c *dataTankController) GetDataTanks(w http.ResponseWriter, r *http.Request) {
	dataTanks, err := c.DataTankUsecase.GetDataTanks()
	if err != nil {
		http.Error(w, "Failed to fetch DataTanks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataTanks)
}

func (c *dataTankController) UpdateDataTank(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid DataTank ID", http.StatusBadRequest)
		return
	}

	var dataTank models.DataTank
	err = json.NewDecoder(r.Body).Decode(&dataTank)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	dataTank.ID = id
	err = c.DataTankUsecase.UpdateDataTank(&dataTank)
	if err != nil {
		http.Error(w, "Failed to update DataTank", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataTank)
}

func (c *dataTankController) DeleteDataTank(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid DataTank ID", http.StatusBadRequest)
		return
	}

	err = c.DataTankUsecase.DeleteDataTank(id)
	if err != nil {
		http.Error(w, "Failed to delete DataTank", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *dataTankController) ReadFromSerial(w http.ResponseWriter, r *http.Request) {
	data, err := c.DataTankUsecase.ReadFromSerial()
	if err != nil {
		http.Error(w, "Failed to read from serial", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": string(data)})
}
