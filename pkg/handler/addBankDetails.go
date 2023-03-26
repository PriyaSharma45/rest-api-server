package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api-server/pkg/model"
	"rest-api-server/pkg/postgres_client"
)

type SuccessMessage struct {
	Message string `json:"message"`
}

func (rh *RouteHandler) AddCardDetails(w http.ResponseWriter, r *http.Request) {

	var request model.CardDetails

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = postgres_client.InsertBankDetails(rh.Engine, request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	message := SuccessMessage{
		Message: "Value inserted successfully ",
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Fatalf("error adding response %f", err)
	}
}

type GetCardDetails struct {
	UserID int `json:"user_id"`
}

func (rh *RouteHandler) GetCardDetails(w http.ResponseWriter, r *http.Request) {
	var request GetCardDetails

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	card_details, err := postgres_client.GetBankDetails(rh.Engine, request.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(card_details)
	if err != nil {
		log.Fatalf("error adding response %f", err)
	}
}
