package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ussii39/go_rest_api/model"
)

type houseHoldResponse struct {
	HouseHold []*model.HouseHold `json:"houseHold"`
}

func RenderhouseHolds(w http.ResponseWriter, houseHold []*model.HouseHold) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	outputJson, err := json.Marshal(houseHold)
	fmt.Print(string(outputJson))
	s, err := json.Marshal(houseHoldResponse{HouseHold: houseHold})
	if err != nil {
		RenderInternalServerError(w, "cant't encode tasks response json")
		return
	}
	fmt.Fprintln(w, string(s))
}

func RenderhouseHold(w http.ResponseWriter, houseHold *model.HouseHold, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	s, err := json.Marshal(houseHold)
	if err != nil {
		RenderInternalServerError(w, "cant't encode houseHold response json")
		return
	}
	fmt.Fprintln(w, string(s))
}