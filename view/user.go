package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ussii39/go_rest_api/model"
)

type usersResponse struct {
	Total int           `json:"total"`
	Users []*model.User `json:"users"`
}

func RenderUsers(w http.ResponseWriter, users []*model.User) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	s, err := json.Marshal(usersResponse{Total: len(users), Users: users})
	if err != nil {
		RenderInternalServerError(w, "cant't encode users response json")
		return
	}
	fmt.Fprintln(w, string(s))
}

func RenderUser(w http.ResponseWriter, user *model.User, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	s, err := json.Marshal(user)
	if err != nil {
		RenderInternalServerError(w, "cant't encode user response json")
		return
	}
	fmt.Fprintln(w, string(s))
}
