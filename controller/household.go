package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ussii39/go_rest_api/model"
	"github.com/ussii39/go_rest_api/view"
)

// UserController require *sql.Db to initialize
// This controller hove CRUD methods
type HouseholdController struct {
	Db *sql.DB
}


// GetUsers return All Users
func (tc *HouseholdController) GetHouseholds(w http.ResponseWriter, r *http.Request) {
	houseHold, err := model.GetHouseHolds(r.Context(), tc.Db)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get houseHold error: %v", err))
		return
	}
	view.RenderhouseHolds(w, houseHold)
}

// CreateHousehold create new Task, and return that Task
func (tc *HouseholdController) CreateHousehold(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	var houseHold model.HouseHold
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	err = json.Unmarshal(body, &houseHold)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("json parse error: %v", err))
		return
	}

	insertID, err := model.CreateHouseHold(r.Context(), tc.Db, &houseHold)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create houseHold error: %v", err))
		return
	}
	createHousehold, err := model.GetHouseHoldByID(r.Context(), tc.Db, insertID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get houseHold error: %v", err))
		return
	}
	view.RenderhouseHold(w, createHousehold, http.StatusCreated)
}