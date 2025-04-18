package handler

import (
	"encoding/json"
	"fmt"
	"gogetsms-api/client"
	"gogetsms-api/model"
	"gogetsms-api/utils"
	"net/http"
)

func GetAllRentList(w http.ResponseWriter, r *http.Request) {
	data, err := client.GetRentList()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	response := model.RentListResponse{
		Status: "success",
		Values: data.Values,
		Quantity: data.Quantity,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetFilteredRentList(w http.ResponseWriter, r *http.Request) {
	data, err := client.GetRentList()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	filtered := make(map[string]model.RentDetail)
	for key, detail := range data.Values {
		fmt.Println("Key: ", key, "IsLessThan3Days: ", utils.IsLessThan3Days(detail.EndDate))
		if utils.IsLessThan3Days(detail.EndDate) {
			filtered[key] = detail
		}
	}

	response := model.RentListResponse{
		Status: "success",
		Values: filtered,
		Quantity: len(filtered),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
